package http

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

// WebhookResponse represents the webhook response
type WebhookResponse struct {
	Status string `json:"status"`
}

// HandleWebhook handles incoming webhook requests
func (h *Handler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("failed to read webhook request body", "error", err)
		h.writeErrorResponse(w, "failed to read request body", err)
		return
	}
	defer r.Body.Close()

	// Log the incoming webhook
	slog.Info("webhook received",
		"content-type", r.Header.Get("Content-Type"),
		"content-length", len(body),
		"body", string(body),
	)

	// Return success response
	response := WebhookResponse{
		Status: "received",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("failed to encode webhook response", "error", err)
	}
}
