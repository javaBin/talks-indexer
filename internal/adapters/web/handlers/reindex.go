package handlers

import (
	"log/slog"
	"net/http"

	"github.com/javaBin/talks-indexer/internal/adapters/web/templates"
)

// HandleReindexAll triggers a full reindex of all conferences
func (h *Handler) HandleReindexAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slog.InfoContext(ctx, "web: starting full reindex")

	err := h.indexer.ReindexAll(ctx)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err != nil {
		slog.ErrorContext(ctx, "web: failed to reindex all", "error", err)
		templates.ResultError("Failed to reindex: "+err.Error()).Render(ctx, w)
		return
	}

	slog.InfoContext(ctx, "web: full reindex completed")
	templates.ResultSuccess("Successfully reindexed all conferences").Render(ctx, w)
}

// HandleReindexConference triggers a reindex for a single conference
func (h *Handler) HandleReindexConference(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	slug := r.FormValue("slug")
	if slug == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		templates.ResultError("Please select a conference").Render(ctx, w)
		return
	}

	slog.InfoContext(ctx, "web: starting conference reindex", "slug", slug)

	err := h.indexer.ReindexConference(ctx, slug)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err != nil {
		slog.ErrorContext(ctx, "web: failed to reindex conference", "slug", slug, "error", err)
		templates.ResultError("Failed to reindex conference: "+err.Error()).Render(ctx, w)
		return
	}

	slog.InfoContext(ctx, "web: conference reindex completed", "slug", slug)
	templates.ResultSuccess("Successfully reindexed conference: "+slug).Render(ctx, w)
}

// HandleReindexTalk triggers a reindex for a single talk
func (h *Handler) HandleReindexTalk(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	talkID := r.FormValue("talkId")
	if talkID == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		templates.ResultError("Please enter a talk ID").Render(ctx, w)
		return
	}

	slog.InfoContext(ctx, "web: starting talk reindex", "talkID", talkID)

	err := h.indexer.ReindexTalk(ctx, talkID)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err != nil {
		slog.ErrorContext(ctx, "web: failed to reindex talk", "talkID", talkID, "error", err)
		templates.ResultError("Failed to reindex talk: "+err.Error()).Render(ctx, w)
		return
	}

	slog.InfoContext(ctx, "web: talk reindex completed", "talkID", talkID)
	templates.ResultSuccess("Successfully reindexed talk: "+talkID).Render(ctx, w)
}
