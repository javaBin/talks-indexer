package domain

import "strings"

// Speaker represents a person presenting a talk at a conference.
type Speaker struct {
	ID   string `json:"id"`
	Name string `json:"name"`

	// Data contains all public data fields from the speaker
	Data map[string]interface{} `json:"data,omitempty"`

	// PrivateData contains fields marked as private (only indexed to private index)
	PrivateData map[string]interface{} `json:"privateData,omitempty"`
}

// ToPublic returns a copy of the Speaker without private data and email fields
func (s Speaker) ToPublic() Speaker {
	return Speaker{
		ID:   s.ID,
		Name: s.Name,
		Data: filterEmailFields(s.Data),
		// PrivateData intentionally omitted
	}
}

// ToPrivate returns a copy of the Speaker with privateData merged into data
func (s Speaker) ToPrivate() Speaker {
	mergedData := make(map[string]interface{})
	for k, v := range s.Data {
		mergedData[k] = v
	}
	for k, v := range s.PrivateData {
		mergedData[k] = v
	}

	return Speaker{
		ID:   s.ID,
		Name: s.Name,
		Data: mergedData,
		// PrivateData intentionally omitted - merged into Data
	}
}

// Speakers is a slice of Speaker with helper methods
type Speakers []Speaker

// ToPublic returns a copy of all speakers without private data
func (ss Speakers) ToPublic() Speakers {
	result := make(Speakers, len(ss))
	for i, s := range ss {
		result[i] = s.ToPublic()
	}
	return result
}

// ToPrivate returns a copy of all speakers with private data merged into data
func (ss Speakers) ToPrivate() Speakers {
	result := make(Speakers, len(ss))
	for i, s := range ss {
		result[i] = s.ToPrivate()
	}
	return result
}

// filterEmailFields removes any fields that contain email addresses
func filterEmailFields(data map[string]interface{}) map[string]interface{} {
	if data == nil {
		return nil
	}

	result := make(map[string]interface{})
	for key, value := range data {
		// Skip fields with "email" in the key name
		if strings.Contains(strings.ToLower(key), "email") {
			continue
		}

		// Check if the value looks like an email address
		if str, ok := value.(string); ok && looksLikeEmail(str) {
			continue
		}

		result[key] = value
	}
	return result
}

// looksLikeEmail checks if a string appears to be an email address
func looksLikeEmail(s string) bool {
	// Simple check: contains @ and has text before and after it
	atIndex := strings.Index(s, "@")
	return atIndex > 0 && atIndex < len(s)-1 && strings.Contains(s[atIndex:], ".")
}
