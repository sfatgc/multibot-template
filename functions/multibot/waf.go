package multibot

import (
	"log"
	"os"

	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/types"
)

func logError(error types.MatchedRule) {
	msg := error.ErrorLog()
	log.Printf("[logError][%s] %s\n", error.Rule().Severity(), msg)
}

func createWAF() coraza.WAF {
	directivesFile := "./serverless_function_source_code/waf.conf"
	if s := os.Getenv("DIRECTIVES_FILE"); s != "" {
		directivesFile = s
	}

	waf, err := coraza.NewWAF(
		coraza.NewWAFConfig().
			WithErrorCallback(logError).
			WithDirectivesFromFile(directivesFile),
	)
	if err != nil {
		log.Fatal(err)
	}
	return waf
}
