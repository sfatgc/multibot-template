package main

import (
	"log"
	"os"

	// Blank-import the function package so the init() runs
	_ "github.com/sfatgc/multibot"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	//funcframework.RegisterHTTPFunctionContext(context.Background(), "dispatchMessages", dispatchMessages)

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
	log.Panicln("started")
}
