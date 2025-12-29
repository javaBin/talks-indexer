package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/javaBin/talks-indexer/internal/adapters/elasticsearch"
	httpAdapter "github.com/javaBin/talks-indexer/internal/adapters/http"
	"github.com/javaBin/talks-indexer/internal/adapters/moresleep"
	webAdapter "github.com/javaBin/talks-indexer/internal/adapters/web"
	"github.com/javaBin/talks-indexer/internal/adapters/web/handlers"
	"github.com/javaBin/talks-indexer/internal/app"
	"github.com/javaBin/talks-indexer/internal/config"
)

func main() {
	// Load configuration first to determine logging mode
	cfg := config.MustLoad()

	// Configure logging based on mode
	var logger *slog.Logger
	if cfg.Mode.IsDevelopment() {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}
	slog.SetDefault(logger)

	logger.Info("configuration loaded",
		"mode", cfg.Mode,
		"port", cfg.Port,
		"moresleepURL", cfg.MoresleepURL,
		"elasticsearchURL", cfg.ElasticsearchURL,
		"privateIndex", cfg.PrivateIndex,
		"publicIndex", cfg.PublicIndex,
	)

	// Initialize moresleep client
	moresleepClient := moresleep.New(
		cfg.MoresleepURL,
		cfg.MoresleepUser,
		cfg.MoresleepPassword,
	)
	logger.Info("moresleep client initialized")

	// Initialize elasticsearch client
	esClient, err := elasticsearch.New(cfg.ElasticsearchURL)
	if err != nil {
		logger.Error("failed to create elasticsearch client", "error", err)
		os.Exit(1)
	}
	logger.Info("elasticsearch client initialized")

	// Create indexer service
	indexerService := app.NewIndexerService(
		moresleepClient,
		esClient,
		cfg.PrivateIndex,
		cfg.PublicIndex,
	)
	logger.Info("indexer service initialized")

	// Create HTTP handler
	handler := httpAdapter.NewHandler(indexerService)

	// Create web handler
	webHandler := handlers.NewHandler(indexerService, moresleepClient)

	// Create HTTP server
	mux := http.NewServeMux()
	httpAdapter.RegisterRoutes(mux, handler)
	webAdapter.RegisterRoutes(mux, webHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 60 * time.Second, // Longer for reindex operations
		IdleTimeout:  60 * time.Second,
	}

	// Start server in goroutine
	go func() {
		logger.Info("starting HTTP server", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server error", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("server shutdown error", "error", err)
		os.Exit(1)
	}

	logger.Info("server stopped")
}
