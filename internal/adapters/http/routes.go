package http

import (
	"net/http"
)

// RegisterRoutes registers all HTTP routes with the provided mux
func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	// Health check endpoint
	mux.HandleFunc("GET /health", h.HandleHealth)

	// Reindex endpoints
	mux.HandleFunc("POST /api/reindex", h.HandleReindexAll)
	mux.HandleFunc("POST /api/reindex/conference/{slug}", h.HandleReindexConference)
	mux.HandleFunc("POST /api/reindex/talk/{talkId}", h.HandleReindexTalk)

	// Webhook endpoint
	mux.HandleFunc("POST /api/webhook", h.HandleWebhook)
}
