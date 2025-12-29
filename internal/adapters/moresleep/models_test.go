package moresleep

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExtractStringValue(t *testing.T) {
	t.Run("valid string value", func(t *testing.T) {
		dv := DataValue{Value: "test string", PrivateData: false}
		result := extractStringValue(dv)
		assert.Equal(t, "test string", result)
	})

	t.Run("nil value", func(t *testing.T) {
		dv := DataValue{Value: nil, PrivateData: false}
		result := extractStringValue(dv)
		assert.Equal(t, "", result)
	})

	t.Run("non-string value", func(t *testing.T) {
		dv := DataValue{Value: 123, PrivateData: false}
		result := extractStringValue(dv)
		assert.Equal(t, "", result)
	})

	t.Run("empty string", func(t *testing.T) {
		dv := DataValue{Value: "", PrivateData: false}
		result := extractStringValue(dv)
		assert.Equal(t, "", result)
	})
}

func TestExtractStringSliceValue(t *testing.T) {
	t.Run("valid interface slice", func(t *testing.T) {
		dv := DataValue{
			Value:       []interface{}{"tag1", "tag2", "tag3"},
			PrivateData: false,
		}
		result := extractStringSliceValue(dv)
		assert.Equal(t, []string{"tag1", "tag2", "tag3"}, result)
	})

	t.Run("valid string slice", func(t *testing.T) {
		dv := DataValue{
			Value:       []string{"keyword1", "keyword2"},
			PrivateData: false,
		}
		result := extractStringSliceValue(dv)
		assert.Equal(t, []string{"keyword1", "keyword2"}, result)
	})

	t.Run("nil value", func(t *testing.T) {
		dv := DataValue{Value: nil, PrivateData: false}
		result := extractStringSliceValue(dv)
		assert.NotNil(t, result)
		assert.Len(t, result, 0)
	})

	t.Run("empty slice", func(t *testing.T) {
		dv := DataValue{Value: []interface{}{}, PrivateData: false}
		result := extractStringSliceValue(dv)
		assert.NotNil(t, result)
		assert.Len(t, result, 0)
	})

	t.Run("mixed type slice - only strings extracted", func(t *testing.T) {
		dv := DataValue{
			Value:       []interface{}{"valid", 123, "another", nil, "third"},
			PrivateData: false,
		}
		result := extractStringSliceValue(dv)
		assert.Equal(t, []string{"valid", "another", "third"}, result)
	})

	t.Run("non-slice value", func(t *testing.T) {
		dv := DataValue{Value: "not a slice", PrivateData: false}
		result := extractStringSliceValue(dv)
		assert.NotNil(t, result)
		assert.Len(t, result, 0)
	})
}

func TestExtractTimeValue(t *testing.T) {
	t.Run("valid RFC3339 time string", func(t *testing.T) {
		timeStr := "2024-12-29T15:04:05Z"
		dv := DataValue{Value: timeStr, PrivateData: false}
		result := extractTimeValue(dv)

		assert.NotNil(t, result)
		expected, _ := time.Parse(time.RFC3339, timeStr)
		assert.True(t, expected.Equal(*result))
	})

	t.Run("valid RFC3339Nano time string", func(t *testing.T) {
		timeStr := "2024-12-29T15:04:05.123456789Z"
		dv := DataValue{Value: timeStr, PrivateData: false}
		result := extractTimeValue(dv)

		assert.NotNil(t, result)
		expected, _ := time.Parse(time.RFC3339Nano, timeStr)
		assert.True(t, expected.Equal(*result))
	})

	t.Run("valid time without timezone", func(t *testing.T) {
		timeStr := "2024-12-29T15:04:05"
		dv := DataValue{Value: timeStr, PrivateData: false}
		result := extractTimeValue(dv)

		assert.NotNil(t, result)
	})

	t.Run("valid time with space separator", func(t *testing.T) {
		timeStr := "2024-12-29 15:04:05"
		dv := DataValue{Value: timeStr, PrivateData: false}
		result := extractTimeValue(dv)

		assert.NotNil(t, result)
	})

	t.Run("nil value", func(t *testing.T) {
		dv := DataValue{Value: nil, PrivateData: false}
		result := extractTimeValue(dv)
		assert.Nil(t, result)
	})

	t.Run("empty string", func(t *testing.T) {
		dv := DataValue{Value: "", PrivateData: false}
		result := extractTimeValue(dv)
		assert.Nil(t, result)
	})

	t.Run("invalid time string", func(t *testing.T) {
		dv := DataValue{Value: "not a valid time", PrivateData: false}
		result := extractTimeValue(dv)
		assert.Nil(t, result)
	})

	t.Run("non-string value", func(t *testing.T) {
		dv := DataValue{Value: 123456789, PrivateData: false}
		result := extractTimeValue(dv)
		assert.Nil(t, result)
	})
}

func TestDataValue_PrivateData(t *testing.T) {
	t.Run("private data flag is preserved", func(t *testing.T) {
		publicDV := DataValue{Value: "public", PrivateData: false}
		privateDV := DataValue{Value: "private", PrivateData: true}

		assert.False(t, publicDV.PrivateData)
		assert.True(t, privateDV.PrivateData)
	})
}
