package handlers

import (
	"context"
	"sync"

	"github.com/javaBin/talks-indexer/internal/domain"
)

// Indexer defines the interface for reindexing operations
type Indexer interface {
	ReindexAll(ctx context.Context) error
	ReindexConference(ctx context.Context, slug string) error
	ReindexTalk(ctx context.Context, talkID string) error
}

// ConferenceProvider defines the interface for fetching conferences
type ConferenceProvider interface {
	GetConferences(ctx context.Context) ([]domain.Conference, error)
}

// Handler handles web UI requests for the admin dashboard
type Handler struct {
	indexer     Indexer
	provider    ConferenceProvider
	conferences []domain.Conference
	confMu      sync.RWMutex
}

// NewHandler creates a new web Handler with the provided dependencies
func NewHandler(indexer Indexer, provider ConferenceProvider) *Handler {
	return &Handler{
		indexer:  indexer,
		provider: provider,
	}
}

// getConferences returns cached conferences, fetching them if not yet cached
func (h *Handler) getConferences(ctx context.Context) ([]domain.Conference, error) {
	h.confMu.RLock()
	if h.conferences != nil {
		defer h.confMu.RUnlock()
		return h.conferences, nil
	}
	h.confMu.RUnlock()

	h.confMu.Lock()
	defer h.confMu.Unlock()

	// Double-check after acquiring write lock
	if h.conferences != nil {
		return h.conferences, nil
	}

	conferences, err := h.provider.GetConferences(ctx)
	if err != nil {
		return nil, err
	}
	h.conferences = conferences
	return conferences, nil
}
