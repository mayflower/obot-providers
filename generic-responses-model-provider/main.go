package main

import (
	"fmt"
	"os"

	"github.com/obot-platform/providers/openai-model-provider/proxy"
)

func main() {
	cfg := &proxy.Config{
		APIKey:          os.Getenv("OBOT_GENERIC_RESPONSES_MODEL_PROVIDER_API_KEY"), // optional
		ListenPort:      os.Getenv("PORT"),
		BaseURL:         os.Getenv("OBOT_GENERIC_RESPONSES_MODEL_PROVIDER_BASE_URL"),
		RewriteModelsFn: proxy.RewriteAllModelsWithUsage("llm"),
		Name:            "Generic Responses",
	}

	if err := proxy.Run(cfg); err != nil {
		fmt.Printf("failed to run generic-responses-model-provider proxy: %v\n", err)
		os.Exit(1)
	}
}
