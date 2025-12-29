package ports

import (
	"context"

	"github.com/javaBin/talks-indexer/internal/domain"
)

// SearchIndex defines the interface for Elasticsearch operations
type SearchIndex interface {
	// BulkIndex indexes multiple talks into the specified index
	BulkIndex(ctx context.Context, indexName string, talks []domain.Talk) error

	// DeleteIndex removes an index from Elasticsearch
	DeleteIndex(ctx context.Context, indexName string) error

	// CreateIndex creates a new index with the specified mapping
	CreateIndex(ctx context.Context, indexName string, mapping string) error

	// IndexExists checks if an index exists in Elasticsearch
	IndexExists(ctx context.Context, indexName string) (bool, error)
}
