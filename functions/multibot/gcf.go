package multibot

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	txhttp "github.com/corazawaf/coraza/v3/http"
)

var TG_BOTS map[string]*TgBot
var FIRESTORE_CLIENT *firestore.Client
var FIRESTORE_ERR error
var PP_STRIPE_TOKEN string

func init() {

	/*
	 * TELEGRAM_BOTS_LIST env var expected to contain comma-separated bot names.
	 * TELEGRAM_BOTS_LIST=bot1,bot2,bot3
	 */
	bots_list_s, env_success := os.LookupEnv("TELEGRAM_BOTS_LIST")
	if !env_success {
		log.Panic("Error getting TELEGRAM_BOTS_LIST environment variable")
	}

	bots_list := strings.Split(bots_list_s, ",")
	if len(bots_list) < 1 {
		log.Panic("No bots defined in TELEGRAM_BOTS_LIST environment variable")
	}

	/*
	 * TG_BOTS global var maps Telegram secrets to TgBot instances
	 *   "telegram secrets" here are those specified as "secret_token"
	 *   in "setWebhook" API call
	 *   (https://core.telegram.org/bots/api#setwebhook)
	 *
	 * So that way we can quickly identify for which bot particular
	 *   request came, by secret specified as
	 *   "X-Telegram-Bot-Api-Secret-Token" HTTP header value
	 */
	TG_BOTS = make(map[string]*TgBot, len(bots_list))

	/*
	 * The set of env vars suffixed with bot name from TELEGRAM_BOTS_LIST
	 *   is expected to provide Telegram token (one that identified bot
	 *   against Telegram) and Telegram secret (one that identifies
	 *   Telegram against bot)
	 */
	for _, bot_name := range bots_list {

		bot_name = strings.ToUpper(bot_name)

		bot_token_env_var_name := fmt.Sprintf("TELEGRAM_BOT_TOKEN_%s", bot_name)
		bot_token, env_success := os.LookupEnv(bot_token_env_var_name)
		if !env_success {

			log.Fatalf("Error getting %s environment variable", bot_token_env_var_name)

		} else {

			bot_secret_env_var_name := fmt.Sprintf("TELEGRAM_BOT_SECRET_%s", bot_name)
			bot_webhook_secret, env_success := os.LookupEnv(bot_secret_env_var_name)
			if !env_success {

				log.Fatalf("Error getting %s environment variable", bot_secret_env_var_name)

			} else {

				bot_webhook_url, env_success := os.LookupEnv("TELEGRAM_BOT_URL")
				if !env_success {

					log.Fatalf("Error getting BOT_WEBHOOK_URL environment variable")

				} else {

					var err error

					TG_BOTS[bot_webhook_secret], err = NewBot(bot_name, bot_token, bot_webhook_secret, bot_webhook_url)

					if err != nil {
						log.Fatalf("Unable to create bot \"%s\": %s", bot_name, err)
					}

					log.Printf("Bot \"%s\" successfully initialized.", bot_name)

				}

			}

		}

	}

	google_project_id, env_success := os.LookupEnv("GOOGLE_PROJECT_ID")
	if !env_success {
		log.Panic("Error getting GOOGLE_PROJECT_ID environment variable")
	}

	google_firestore_db_id, env_success := os.LookupEnv("GOOGLE_FIRESTORE_DB_ID")
	if !env_success {
		log.Panic("Error getting GOOGLE_FIRESTORE_DB_ID environment variable")
	}

	PP_STRIPE_TOKEN, env_success = os.LookupEnv("PP_STRIPE_TOKEN")
	if !env_success {
		log.Panic("Error getting PP_STRIPE_TOKEN environment variable")
	}

	if FIRESTORE_CLIENT == nil || FIRESTORE_ERR != nil {

		FIRESTORE_CLIENT, FIRESTORE_ERR = firestore.NewClientWithDatabase(context.TODO(), google_project_id, google_firestore_db_id)

		if FIRESTORE_ERR != nil {
			log.Panicf("Error initialising firestore client: \"%s\"", FIRESTORE_ERR)
		}
	}

	waf := createWAF()
	waf_http_handler := txhttp.WrapHandler(waf, http.HandlerFunc(entrypoint))

	functions.HTTP("entrypoint", waf_http_handler.ServeHTTP)

}

func entrypoint(w http.ResponseWriter, r *http.Request) {
	bot_secret := r.Header.Get("X-Telegram-Bot-Api-Secret-Token")
	if bot_secret == "" {
		log.Panic("Header X-Telegram-Bot-Api-Secret-Token is not provided in request. Quiting.")
	}

	bot, ok := TG_BOTS[bot_secret]

	if ok {
		bot.TgWebhook.ServeHTTP(w, r)
	} else {
		log.Panicf("No bot defined for secret %sXXXXXXXXXX%s. Quitting.", bot_secret[0:2], bot_secret[len(bot_secret)-2:])
	}

}
