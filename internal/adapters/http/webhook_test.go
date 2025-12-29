package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleWebhook_Success(t *testing.T) {
	indexer := &mockIndexer{}
	handler := NewHandler(indexer)

	// Create request with body
	body := `{"event": "talk.updated", "conferenceId": "javazone-2024"}`
	req := httptest.NewRequest(http.MethodPost, "/api/webhook", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Call handler
	handler.HandleWebhook(w, req)

	// Assert response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	// Parse response body
	var response WebhookResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, "received", response.Status)
}

func TestHandleWebhook_EmptyBody(t *testing.T) {
	indexer := &mockIndexer{}
	handler := NewHandler(indexer)

	// Create request with empty body
	req := httptest.NewRequest(http.MethodPost, "/api/webhook", strings.NewReader(""))
	w := httptest.NewRecorder()

	// Call handler
	handler.HandleWebhook(w, req)

	// Assert response - should still succeed
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response WebhookResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, "received", response.Status)
}

func TestHandleWebhook_LargeBody(t *testing.T) {
	indexer := &mockIndexer{}
	handler := NewHandler(indexer)

	// Create request with large body
	largeBody := strings.Repeat("x", 10000)
	req := httptest.NewRequest(http.MethodPost, "/api/webhook", strings.NewReader(largeBody))
	w := httptest.NewRecorder()

	// Call handler
	handler.HandleWebhook(w, req)

	// Assert response
	assert.Equal(t, http.StatusOK, w.Code)

	var response WebhookResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, "received", response.Status)
}

func TestHandleWebhook_DifferentContentTypes(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		body        string
	}{
		{
			name:        "JSON content type",
			contentType: "application/json",
			body:        `{"key": "value"}`,
		},
		{
			name:        "Form content type",
			contentType: "application/x-www-form-urlencoded",
			body:        "key=value",
		},
		{
			name:        "Text content type",
			contentType: "text/plain",
			body:        "plain text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexer := &mockIndexer{}
			handler := NewHandler(indexer)

			req := httptest.NewRequest(http.MethodPost, "/api/webhook", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", tt.contentType)
			w := httptest.NewRecorder()

			handler.HandleWebhook(w, req)

			assert.Equal(t, http.StatusOK, w.Code)

			var response WebhookResponse
			err := json.NewDecoder(w.Body).Decode(&response)
			require.NoError(t, err)

			assert.Equal(t, "received", response.Status)
		})
	}
}
