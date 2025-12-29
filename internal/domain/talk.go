package domain

import "time"

// Talk represents a conference talk submission with all fields needed for indexing.
// Data fields are stored dynamically to accommodate varying fields across conferences.
type Talk struct {
	ID             string     `json:"id"`
	ConferenceID   string     `json:"conferenceId"`
	ConferenceSlug string     `json:"conferenceSlug"`
	ConferenceName string     `json:"conferenceName"`
	Status         string     `json:"status"`
	Speakers       Speakers   `json:"speakers"`
	Created        *time.Time `json:"created,omitempty"`
	LastUpdated    *time.Time `json:"lastUpdated,omitempty"`

	// Data contains all public data fields from the talk submission
	Data map[string]interface{} `json:"data,omitempty"`

	// PrivateData contains fields marked as private (only indexed to private index)
	PrivateData map[string]interface{} `json:"privateData,omitempty"`
}

// ToPublic returns a copy of the Talk without private data and email fields for public indexing
func (t Talk) ToPublic() Talk {
	return Talk{
		ID:             t.ID,
		ConferenceID:   t.ConferenceID,
		ConferenceSlug: t.ConferenceSlug,
		ConferenceName: t.ConferenceName,
		Status:         t.Status,
		Speakers:       t.Speakers.ToPublic(),
		Created:        t.Created,
		LastUpdated:    t.LastUpdated,
		Data:           filterEmailFields(t.Data),
		// PrivateData intentionally omitted
	}
}

// ToPrivate returns a copy of the Talk with privateData merged into data for private indexing
func (t Talk) ToPrivate() Talk {
	// Merge data and privateData into a single map
	mergedData := make(map[string]interface{})
	for k, v := range t.Data {
		mergedData[k] = v
	}
	for k, v := range t.PrivateData {
		mergedData[k] = v
	}

	return Talk{
		ID:             t.ID,
		ConferenceID:   t.ConferenceID,
		ConferenceSlug: t.ConferenceSlug,
		ConferenceName: t.ConferenceName,
		Status:         t.Status,
		Speakers:       t.Speakers.ToPrivate(),
		Created:        t.Created,
		LastUpdated:    t.LastUpdated,
		Data:           mergedData,
		// PrivateData intentionally omitted - merged into Data
	}
}
