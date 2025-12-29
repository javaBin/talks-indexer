package handlers

import (
	"log/slog"
	"net/http"

	"github.com/javaBin/talks-indexer/internal/adapters/web/templates"
)

// HandleDashboard renders the admin dashboard page
func (h *Handler) HandleDashboard(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	conferences, err := h.getConferences(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to fetch conferences", "error", err)
		http.Error(w, "Failed to load conferences", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := templates.Dashboard(conferences).Render(ctx, w); err != nil {
		slog.ErrorContext(ctx, "failed to render dashboard", "error", err)
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
		return
	}
}
