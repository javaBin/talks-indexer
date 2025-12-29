package http

import (
	"context"
)

// Indexer defines the interface for indexing operations
type Indexer interface {
	// ReindexAll triggers a full reindex of all conferences
	ReindexAll(ctx context.Context) error

	// ReindexConference reindexes a specific conference by its slug
	ReindexConference(ctx context.Context, slug string) error

	// ReindexTalk reindexes a specific talk by its ID
	ReindexTalk(ctx context.Context, talkID string) error
}

// Handler holds the HTTP handler dependencies
type Handler struct {
	indexer Indexer
}

// NewHandler creates a new HTTP handler with the provided indexer service
func NewHandler(indexer Indexer) *Handler {
	return &Handler{
		indexer: indexer,
	}
}
