package ports

import (
	"context"

	"github.com/javaBin/talks-indexer/internal/domain"
)

// TalkSource defines the interface for fetching talks from moresleep
type TalkSource interface {
	// GetConferences retrieves all available conferences
	GetConferences(ctx context.Context) ([]domain.Conference, error)

	// GetTalks retrieves all talks for a specific conference
	GetTalks(ctx context.Context, conferenceID string) ([]domain.Talk, error)

	// GetTalk retrieves a single talk by its ID
	GetTalk(ctx context.Context, talkID string) (*domain.Talk, error)
}
