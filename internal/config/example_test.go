package config_test

import (
	"context"
	"fmt"

	"github.com/javaBin/talks-indexer/internal/config"
)

// ExampleLoad demonstrates how to load configuration from environment variables
func ExampleLoad() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	fmt.Printf("Port: %d\n", cfg.Port)
	fmt.Printf("Moresleep URL: %s\n", cfg.MoresleepURL)
	// Output will vary based on environment variables
}

// ExampleMustLoad demonstrates how to load configuration with panic on error
func ExampleMustLoad() {
	cfg := config.MustLoad()
	fmt.Printf("Config loaded successfully with port: %d\n", cfg.Port)
	// Output will vary based on environment variables
}

// ExampleWithConfig demonstrates how to attach config to context
func ExampleWithConfig() {
	cfg := &config.Config{
		Port:             8080,
		MoresleepURL:     "http://localhost:8082",
		ElasticsearchURL: "http://localhost:9200",
		PrivateIndex:     "javazone_private",
		PublicIndex:      "javazone_public",
	}

	ctx := context.Background()
	ctx = config.WithConfig(ctx, cfg)

	// Now the context contains the config and can be passed through the application
	retrievedCfg := config.GetConfig(ctx)
	fmt.Printf("Retrieved port from context: %d\n", retrievedCfg.Port)
	// Output: Retrieved port from context: 8080
}
