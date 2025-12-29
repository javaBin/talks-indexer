package moresleep

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMapConference(t *testing.T) {
	cr := ConferenceResponse{
		ID:   "conf-123",
		Name: "JavaZone 2024",
		Slug: "javazone2024",
	}

	conf := MapConference(cr)

	assert.Equal(t, "conf-123", conf.ID)
	assert.Equal(t, "JavaZone 2024", conf.Name)
	assert.Equal(t, "javazone2024", conf.Slug)
}

func TestMapConferences(t *testing.T) {
	crs := []ConferenceResponse{
		{ID: "conf-1", Name: "Conference 1", Slug: "conf1"},
		{ID: "conf-2", Name: "Conference 2", Slug: "conf2"},
	}

	confs := MapConferences(crs)

	assert.Len(t, confs, 2)
	assert.Equal(t, "conf-1", confs[0].ID)
	assert.Equal(t, "conf-2", confs[1].ID)
}

func TestMapConferences_Empty(t *testing.T) {
	crs := []ConferenceResponse{}
	confs := MapConferences(crs)

	assert.NotNil(t, confs)
	assert.Len(t, confs, 0)
}

func TestMapSpeaker(t *testing.T) {
	t.Run("with all data fields", func(t *testing.T) {
		sr := SpeakerResponse{
			ID:    "speaker-1",
			Name:  "Jane Doe",
			Email: "jane@example.com",
			Data: map[string]DataValue{
				"bio":        {Value: "Experienced developer", PrivateData: false},
				"twitter":    {Value: "@janedoe", PrivateData: false},
				"pictureUrl": {Value: "https://example.com/jane.jpg", PrivateData: false},
			},
		}

		speaker := MapSpeaker(sr)

		assert.Equal(t, "speaker-1", speaker.ID)
		assert.Equal(t, "Jane Doe", speaker.Name)
		assert.Equal(t, "Experienced developer", speaker.Data["bio"])
		assert.Equal(t, "@janedoe", speaker.Data["twitter"])
		assert.Equal(t, "https://example.com/jane.jpg", speaker.Data["pictureUrl"])
	})

	t.Run("with missing data fields", func(t *testing.T) {
		sr := SpeakerResponse{
			ID:    "speaker-2",
			Name:  "John Smith",
			Email: "john@example.com",
			Data:  map[string]DataValue{},
		}

		speaker := MapSpeaker(sr)

		assert.Equal(t, "speaker-2", speaker.ID)
		assert.Equal(t, "John Smith", speaker.Name)
		assert.Nil(t, speaker.Data["bio"])
		assert.Nil(t, speaker.Data["twitter"])
		assert.Nil(t, speaker.Data["pictureUrl"])
	})

	t.Run("with nil data values", func(t *testing.T) {
		sr := SpeakerResponse{
			ID:   "speaker-3",
			Name: "Alice Brown",
			Data: map[string]DataValue{
				"bio":        {Value: nil, PrivateData: false},
				"twitter":    {Value: nil, PrivateData: false},
				"pictureUrl": {Value: nil, PrivateData: false},
			},
		}

		speaker := MapSpeaker(sr)

		assert.Equal(t, "speaker-3", speaker.ID)
		assert.Nil(t, speaker.Data["bio"])
		assert.Nil(t, speaker.Data["twitter"])
		assert.Nil(t, speaker.Data["pictureUrl"])
	})
}

func TestMapSpeakers(t *testing.T) {
	srs := []SpeakerResponse{
		{
			ID:   "speaker-1",
			Name: "Speaker One",
			Data: map[string]DataValue{
				"bio": {Value: "Bio 1", PrivateData: false},
			},
		},
		{
			ID:   "speaker-2",
			Name: "Speaker Two",
			Data: map[string]DataValue{
				"bio": {Value: "Bio 2", PrivateData: false},
			},
		},
	}

	speakers := MapSpeakers(srs)

	assert.Len(t, speakers, 2)
	assert.Equal(t, "speaker-1", speakers[0].ID)
	assert.Equal(t, "speaker-2", speakers[1].ID)
}

func TestMapSpeakers_Empty(t *testing.T) {
	srs := []SpeakerResponse{}
	speakers := MapSpeakers(srs)

	assert.NotNil(t, speakers)
	assert.Len(t, speakers, 0)
}

func TestMapTalk(t *testing.T) {
	now := time.Now()
	startTime := now.Add(1 * time.Hour).Format(time.RFC3339)
	endTime := now.Add(2 * time.Hour).Format(time.RFC3339)

	t.Run("with all data fields", func(t *testing.T) {
		sr := SessionResponse{
			ID:           "talk-1",
			ConferenceID: "conf-1",
			Status:       "APPROVED",
			PostedBy:     "speaker@example.com",
			Data: map[string]DataValue{
				"title":            {Value: "Advanced Go Patterns", PrivateData: false},
				"abstract":         {Value: "Learn advanced patterns in Go", PrivateData: false},
				"intendedAudience": {Value: "Advanced developers", PrivateData: false},
				"language":         {Value: "English", PrivateData: false},
				"format":           {Value: "Workshop", PrivateData: false},
				"level":            {Value: "Advanced", PrivateData: false},
				"keywords":         {Value: []interface{}{"go", "patterns", "advanced"}, PrivateData: false},
				"room":             {Value: "Room B", PrivateData: false},
				"startTime":        {Value: startTime, PrivateData: false},
				"endTime":          {Value: endTime, PrivateData: false},
			},
			Speakers: []SpeakerResponse{
				{
					ID:   "speaker-1",
					Name: "Expert Speaker",
					Data: map[string]DataValue{
						"bio": {Value: "Go expert", PrivateData: false},
					},
				},
			},
			Created:     FlexibleTime{Time: now},
			LastUpdated: FlexibleTime{Time: now},
		}

		talk := MapTalk(sr, "javazone2024", "JavaZone 2024")

		assert.Equal(t, "talk-1", talk.ID)
		assert.Equal(t, "conf-1", talk.ConferenceID)
		assert.Equal(t, "javazone2024", talk.ConferenceSlug)
		assert.Equal(t, "JavaZone 2024", talk.ConferenceName)
		assert.Equal(t, "Advanced Go Patterns", talk.Data["title"])
		assert.Equal(t, "Learn advanced patterns in Go", talk.Data["abstract"])
		assert.Equal(t, "Advanced developers", talk.Data["intendedAudience"])
		assert.Equal(t, "English", talk.Data["language"])
		assert.Equal(t, "Workshop", talk.Data["format"])
		assert.Equal(t, "Advanced", talk.Data["level"])
		assert.Equal(t, []interface{}{"go", "patterns", "advanced"}, talk.Data["keywords"])
		assert.Equal(t, "APPROVED", talk.Status)
		assert.Equal(t, "Room B", talk.Data["room"])
		assert.Equal(t, "speaker@example.com", talk.PrivateData["postedBy"])
		assert.Equal(t, startTime, talk.Data["startTime"])
		assert.Equal(t, endTime, talk.Data["endTime"])
		assert.Len(t, talk.Speakers, 1)
		assert.Equal(t, "Expert Speaker", talk.Speakers[0].Name)
	})

	t.Run("with minimal data fields", func(t *testing.T) {
		sr := SessionResponse{
			ID:           "talk-2",
			ConferenceID: "conf-2",
			Status:       "SUBMITTED",
			PostedBy:     "newbie@example.com",
			Data:         map[string]DataValue{},
			Speakers:     []SpeakerResponse{},
			Created:      FlexibleTime{Time: now},
			LastUpdated:  FlexibleTime{Time: now},
		}

		talk := MapTalk(sr, "test-conf", "Test Conference")

		assert.Equal(t, "talk-2", talk.ID)
		assert.Equal(t, "conf-2", talk.ConferenceID)
		assert.Equal(t, "test-conf", talk.ConferenceSlug)
		assert.Equal(t, "Test Conference", talk.ConferenceName)
		assert.Nil(t, talk.Data["title"])
		assert.Nil(t, talk.Data["abstract"])
		assert.Nil(t, talk.Data["intendedAudience"])
		assert.Nil(t, talk.Data["language"])
		assert.Nil(t, talk.Data["format"])
		assert.Nil(t, talk.Data["level"])
		assert.Nil(t, talk.Data["keywords"])
		assert.Equal(t, "SUBMITTED", talk.Status)
		assert.Nil(t, talk.Data["room"])
		assert.Equal(t, "newbie@example.com", talk.PrivateData["postedBy"])
		assert.Nil(t, talk.Data["startTime"])
		assert.Nil(t, talk.Data["endTime"])
		assert.Len(t, talk.Speakers, 0)
	})

	t.Run("with nil data values", func(t *testing.T) {
		sr := SessionResponse{
			ID:           "talk-3",
			ConferenceID: "conf-3",
			Status:       "DRAFT",
			PostedBy:     "test@example.com",
			Data: map[string]DataValue{
				"title":            {Value: nil, PrivateData: false},
				"abstract":         {Value: nil, PrivateData: false},
				"intendedAudience": {Value: nil, PrivateData: false},
				"language":         {Value: nil, PrivateData: false},
				"format":           {Value: nil, PrivateData: false},
				"level":            {Value: nil, PrivateData: false},
				"keywords":         {Value: nil, PrivateData: false},
				"room":             {Value: nil, PrivateData: false},
				"startTime":        {Value: nil, PrivateData: false},
				"endTime":          {Value: nil, PrivateData: false},
			},
			Speakers:    []SpeakerResponse{},
			Created:     FlexibleTime{Time: now},
			LastUpdated: FlexibleTime{Time: now},
		}

		talk := MapTalk(sr, "", "")

		assert.Nil(t, talk.Data["title"])
		assert.Nil(t, talk.Data["abstract"])
		assert.Nil(t, talk.Data["keywords"])
		assert.Nil(t, talk.Data["startTime"])
		assert.Nil(t, talk.Data["endTime"])
	})
}

func TestMapTalks(t *testing.T) {
	now := time.Now()

	srs := []SessionResponse{
		{
			ID:           "talk-1",
			ConferenceID: "conf-1",
			Status:       "APPROVED",
			PostedBy:     "speaker1@example.com",
			Data: map[string]DataValue{
				"title": {Value: "Talk 1", PrivateData: false},
			},
			Speakers:    []SpeakerResponse{},
			Created:     FlexibleTime{Time: now},
			LastUpdated: FlexibleTime{Time: now},
		},
		{
			ID:           "talk-2",
			ConferenceID: "conf-1",
			Status:       "SUBMITTED",
			PostedBy:     "speaker2@example.com",
			Data: map[string]DataValue{
				"title": {Value: "Talk 2", PrivateData: false},
			},
			Speakers:    []SpeakerResponse{},
			Created:     FlexibleTime{Time: now},
			LastUpdated: FlexibleTime{Time: now},
		},
	}

	talks := MapTalks(srs, "javazone2024", "JavaZone 2024")

	assert.Len(t, talks, 2)
	assert.Equal(t, "talk-1", talks[0].ID)
	assert.Equal(t, "Talk 1", talks[0].Data["title"])
	assert.Equal(t, "javazone2024", talks[0].ConferenceSlug)
	assert.Equal(t, "JavaZone 2024", talks[0].ConferenceName)
	assert.Equal(t, "talk-2", talks[1].ID)
	assert.Equal(t, "Talk 2", talks[1].Data["title"])
	assert.Equal(t, "javazone2024", talks[1].ConferenceSlug)
	assert.Equal(t, "JavaZone 2024", talks[1].ConferenceName)
}

func TestMapTalks_Empty(t *testing.T) {
	srs := []SessionResponse{}
	talks := MapTalks(srs, "test", "Test")

	assert.NotNil(t, talks)
	assert.Len(t, talks, 0)
}
