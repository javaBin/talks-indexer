package web

import (
	"net/http"

	"github.com/javaBin/talks-indexer/internal/adapters/web/handlers"
)

// RegisterRoutes registers all web routes with the provided mux
func RegisterRoutes(mux *http.ServeMux, h *handlers.Handler) {
	// Admin dashboard
	mux.HandleFunc("GET /admin", h.HandleDashboard)

	// htmx endpoints for reindex operations
	mux.HandleFunc("POST /admin/reindex/all", h.HandleReindexAll)
	mux.HandleFunc("POST /admin/reindex/conference", h.HandleReindexConference)
	mux.HandleFunc("POST /admin/reindex/talk", h.HandleReindexTalk)
}
