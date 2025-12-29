package moresleep

import (
	"github.com/javaBin/talks-indexer/internal/domain"
)

// isEmptyValue checks if a value is nil or an empty string.
// Empty values cause issues with Elasticsearch dynamic type mapping.
func isEmptyValue(v interface{}) bool {
	if v == nil {
		return true
	}
	if s, ok := v.(string); ok && s == "" {
		return true
	}
	return false
}

// MapConference converts a ConferenceResponse to a domain.Conference
func MapConference(cr ConferenceResponse) domain.Conference {
	return domain.Conference{
		ID:   cr.ID,
		Name: cr.Name,
		Slug: cr.Slug,
	}
}

// MapConferences converts a slice of ConferenceResponse to domain.Conference
func MapConferences(crs []ConferenceResponse) []domain.Conference {
	conferences := make([]domain.Conference, 0, len(crs))
	for _, cr := range crs {
		conferences = append(conferences, MapConference(cr))
	}
	return conferences
}

// MapSpeaker converts a SpeakerResponse to a domain.Speaker
// Separates public and private data fields
func MapSpeaker(sr SpeakerResponse) domain.Speaker {
	speaker := domain.Speaker{
		ID:          sr.ID,
		Name:        sr.Name,
		Data:        make(map[string]interface{}),
		PrivateData: make(map[string]interface{}),
	}

	// Extract all data fields, separating public and private
	// Skip empty values to avoid Elasticsearch dynamic mapping issues
	for key, dv := range sr.Data {
		if isEmptyValue(dv.Value) {
			continue
		}
		if dv.PrivateData {
			speaker.PrivateData[key] = dv.Value
		} else {
			speaker.Data[key] = dv.Value
		}
	}

	return speaker
}

// MapSpeakers converts a slice of SpeakerResponse to domain.Speakers
func MapSpeakers(srs []SpeakerResponse) domain.Speakers {
	speakers := make(domain.Speakers, 0, len(srs))
	for _, sr := range srs {
		speakers = append(speakers, MapSpeaker(sr))
	}
	return speakers
}

// MapTalk converts a SessionResponse to a domain.Talk
// conferenceSlug and conferenceName are needed as they're not part of the session response
// Separates public and private data fields
func MapTalk(sr SessionResponse, conferenceSlug, conferenceName string) domain.Talk {
	talk := domain.Talk{
		ID:             sr.ID,
		ConferenceID:   sr.ConferenceID,
		ConferenceSlug: conferenceSlug,
		ConferenceName: conferenceName,
		Status:         sr.Status,
		Speakers:       MapSpeakers(sr.Speakers),
		Data:           make(map[string]interface{}),
		PrivateData:    make(map[string]interface{}),
	}

	// Only set timestamps if they have valid values
	if !sr.Created.IsZero() {
		talk.Created = &sr.Created.Time
	}
	if !sr.LastUpdated.IsZero() {
		talk.LastUpdated = &sr.LastUpdated.Time
	}

	// Extract all data fields, separating public and private
	// Skip empty values to avoid Elasticsearch dynamic mapping issues
	for key, dv := range sr.Data {
		if isEmptyValue(dv.Value) {
			continue
		}
		if dv.PrivateData {
			talk.PrivateData[key] = dv.Value
		} else {
			talk.Data[key] = dv.Value
		}
	}

	// Add postedBy (submitter email) to private data
	if sr.PostedBy != "" {
		talk.PrivateData["postedBy"] = sr.PostedBy
	}

	return talk
}

// MapTalks converts a slice of SessionResponse to domain.Talk
func MapTalks(srs []SessionResponse, conferenceSlug, conferenceName string) []domain.Talk {
	talks := make([]domain.Talk, 0, len(srs))
	for _, sr := range srs {
		talks = append(talks, MapTalk(sr, conferenceSlug, conferenceName))
	}
	return talks
}
