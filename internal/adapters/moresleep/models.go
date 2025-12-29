package moresleep

import (
	"strings"
	"time"
)

// DataValue represents the nested data structure used by moresleep API
// for both conference and session data fields
type DataValue struct {
	Value       interface{} `json:"value"`
	PrivateData bool        `json:"privateData"`
}

// FlexibleTime is a custom time type that handles various time formats from the API
type FlexibleTime struct {
	time.Time
}

// UnmarshalJSON implements custom JSON unmarshaling for flexible time formats
func (ft *FlexibleTime) UnmarshalJSON(data []byte) error {
	// Remove quotes
	s := strings.Trim(string(data), "\"")
	if s == "" || s == "null" {
		return nil
	}

	// Try various time formats
	formats := []string{
		time.RFC3339,
		time.RFC3339Nano,
		"2006-01-02T15:04:05.999999",
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
	}

	var parseErr error
	for _, format := range formats {
		t, err := time.Parse(format, s)
		if err == nil {
			ft.Time = t
			return nil
		}
		parseErr = err
	}

	return parseErr
}

// ConferenceResponse represents the API response for a conference
type ConferenceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// SessionResponse represents the API response for a talk/session
type SessionResponse struct {
	ID           string               `json:"id"`
	ConferenceID string               `json:"conferenceId"`
	Status       string               `json:"status"`
	PostedBy     string               `json:"postedBy"`
	Data         map[string]DataValue `json:"data"`
	Speakers     []SpeakerResponse    `json:"speakers"`
	Created      FlexibleTime         `json:"created"`
	LastUpdated  FlexibleTime         `json:"lastUpdated"`
}

// SpeakerResponse represents the API response for a speaker
type SpeakerResponse struct {
	ID    string               `json:"id"`
	Name  string               `json:"name"`
	Email string               `json:"email"`
	Data  map[string]DataValue `json:"data"`
}

// ConferencesAPIResponse wraps the list of conferences
type ConferencesAPIResponse struct {
	Conferences []ConferenceResponse `json:"conferences"`
}

// SessionsAPIResponse wraps the list of sessions
type SessionsAPIResponse struct {
	Sessions []SessionResponse `json:"sessions"`
}

// extractStringValue safely extracts a string value from DataValue
func extractStringValue(dv DataValue) string {
	if dv.Value == nil {
		return ""
	}
	if str, ok := dv.Value.(string); ok {
		return str
	}
	return ""
}

// extractStringSliceValue safely extracts a string slice from DataValue
func extractStringSliceValue(dv DataValue) []string {
	if dv.Value == nil {
		return []string{}
	}

	// Handle interface{} slice
	if slice, ok := dv.Value.([]interface{}); ok {
		result := make([]string, 0, len(slice))
		for _, item := range slice {
			if str, ok := item.(string); ok {
				result = append(result, str)
			}
		}
		return result
	}

	// Handle string slice directly
	if slice, ok := dv.Value.([]string); ok {
		return slice
	}

	return []string{}
}

// extractTimeValue safely extracts a time.Time value from DataValue
func extractTimeValue(dv DataValue) *time.Time {
	if dv.Value == nil {
		return nil
	}

	// Try to parse as string
	if str, ok := dv.Value.(string); ok && str != "" {
		// Try RFC3339 format first
		if t, err := time.Parse(time.RFC3339, str); err == nil {
			return &t
		}
		// Try other common formats
		formats := []string{
			time.RFC3339Nano,
			"2006-01-02T15:04:05",
			"2006-01-02 15:04:05",
		}
		for _, format := range formats {
			if t, err := time.Parse(format, str); err == nil {
				return &t
			}
		}
	}

	return nil
}
