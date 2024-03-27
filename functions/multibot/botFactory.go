package multibot

import "gopkg.in/telebot.v3"

func CreateBot(bot_name string, bot_telegram_token string, bot_webhook_secret string, bot_webhook_url string) (TgBotInterface, error) {

	tbs := telebot.Settings{Token: bot_telegram_token, Verbose: true}

	tele_bot, err := telebot.NewBot(tbs)
	if err != nil {
		return nil, err
	}

	var bot TgBotInterface

	switch bot_name {
	case "BULGAKTEER":
		bot = &TgBotBulgakteer{
			TgBot{bot_name,
				err,
				tbs.Verbose,
				telebot.Webhook{MaxConnections: 5, SecretToken: bot_webhook_secret, Endpoint: &telebot.WebhookEndpoint{PublicURL: bot_webhook_url}},
				tele_bot},
		}
	case "SFATGC":
		bot = &TgBotSFATGC{
			TgBot{bot_name,
				err,
				tbs.Verbose,
				telebot.Webhook{MaxConnections: 5, SecretToken: bot_webhook_secret, Endpoint: &telebot.WebhookEndpoint{PublicURL: bot_webhook_url}},
				tele_bot},
		}
	default:
		bot = &TgBot{
			bot_name,
			err,
			tbs.Verbose,
			telebot.Webhook{MaxConnections: 5, SecretToken: bot_webhook_secret, Endpoint: &telebot.WebhookEndpoint{PublicURL: bot_webhook_url}},
			tele_bot,
		}
	}

	bot.GetTgBot().Handle(telebot.OnText, bot.Handle)

	if !bot.CheckWebhook() {
		err := bot.GetTgBot().SetWebhook(bot.GetTgWebhook())

		if err != nil {
			return bot, err
		}
	}

	return bot, nil
}
