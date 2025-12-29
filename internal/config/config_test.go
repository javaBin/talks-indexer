package config

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		expected *Config
		wantErr  bool
	}{
		{
			name:    "load with defaults",
			envVars: map[string]string{},
			expected: &Config{
				Port:              8080,
				MoresleepURL:      "http://localhost:8082",
				MoresleepUser:     "",
				MoresleepPassword: "",
				ElasticsearchURL:  "http://localhost:9200",
				PrivateIndex:      "javazone_private",
				PublicIndex:       "javazone_public",
			},
			wantErr: false,
		},
		{
			name: "load with custom values",
			envVars: map[string]string{
				"PORT":               "9090",
				"MORESLEEP_URL":      "https://api.example.com",
				"MORESLEEP_USER":     "testuser",
				"MORESLEEP_PASSWORD": "testpass",
				"ELASTICSEARCH_URL":  "https://es.example.com:9200",
				"PRIVATE_INDEX":      "custom_private",
				"PUBLIC_INDEX":       "custom_public",
			},
			expected: &Config{
				Port:              9090,
				MoresleepURL:      "https://api.example.com",
				MoresleepUser:     "testuser",
				MoresleepPassword: "testpass",
				ElasticsearchURL:  "https://es.example.com:9200",
				PrivateIndex:      "custom_private",
				PublicIndex:       "custom_public",
			},
			wantErr: false,
		},
		{
			name: "load with partial custom values",
			envVars: map[string]string{
				"PORT":           "3000",
				"MORESLEEP_USER": "admin",
			},
			expected: &Config{
				Port:              3000,
				MoresleepURL:      "http://localhost:8082",
				MoresleepUser:     "admin",
				MoresleepPassword: "",
				ElasticsearchURL:  "http://localhost:9200",
				PrivateIndex:      "javazone_private",
				PublicIndex:       "javazone_public",
			},
			wantErr: false,
		},
		{
			name: "invalid port value",
			envVars: map[string]string{
				"PORT": "invalid",
			},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear environment before each test
			clearConfigEnv()

			// Set test environment variables
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}
			defer clearConfigEnv()

			cfg, err := Load()

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, cfg)
			} else {
				require.NoError(t, err)
				require.NotNil(t, cfg)
				assert.Equal(t, tt.expected.Port, cfg.Port)
				assert.Equal(t, tt.expected.MoresleepURL, cfg.MoresleepURL)
				assert.Equal(t, tt.expected.MoresleepUser, cfg.MoresleepUser)
				assert.Equal(t, tt.expected.MoresleepPassword, cfg.MoresleepPassword)
				assert.Equal(t, tt.expected.ElasticsearchURL, cfg.ElasticsearchURL)
				assert.Equal(t, tt.expected.PrivateIndex, cfg.PrivateIndex)
				assert.Equal(t, tt.expected.PublicIndex, cfg.PublicIndex)
			}
		})
	}
}

func TestWithConfig(t *testing.T) {
	cfg := &Config{
		Port:              8080,
		MoresleepURL:      "http://localhost:8082",
		MoresleepUser:     "testuser",
		MoresleepPassword: "testpass",
		ElasticsearchURL:  "http://localhost:9200",
		PrivateIndex:      "javazone_private",
		PublicIndex:       "javazone_public",
	}

	ctx := context.Background()
	ctxWithConfig := WithConfig(ctx, cfg)

	assert.NotNil(t, ctxWithConfig)
	assert.NotEqual(t, ctx, ctxWithConfig)
}

func TestGetConfig(t *testing.T) {
	t.Run("get config from context", func(t *testing.T) {
		cfg := &Config{
			Port:              9090,
			MoresleepURL:      "https://api.example.com",
			MoresleepUser:     "user",
			MoresleepPassword: "pass",
			ElasticsearchURL:  "https://es.example.com:9200",
			PrivateIndex:      "private",
			PublicIndex:       "public",
		}

		ctx := WithConfig(context.Background(), cfg)
		retrievedCfg := GetConfig(ctx)

		require.NotNil(t, retrievedCfg)
		assert.Equal(t, cfg.Port, retrievedCfg.Port)
		assert.Equal(t, cfg.MoresleepURL, retrievedCfg.MoresleepURL)
		assert.Equal(t, cfg.MoresleepUser, retrievedCfg.MoresleepUser)
		assert.Equal(t, cfg.MoresleepPassword, retrievedCfg.MoresleepPassword)
		assert.Equal(t, cfg.ElasticsearchURL, retrievedCfg.ElasticsearchURL)
		assert.Equal(t, cfg.PrivateIndex, retrievedCfg.PrivateIndex)
		assert.Equal(t, cfg.PublicIndex, retrievedCfg.PublicIndex)
	})

	t.Run("panic when config not in context", func(t *testing.T) {
		ctx := context.Background()
		assert.Panics(t, func() {
			GetConfig(ctx)
		})
	})
}

func TestMustLoad(t *testing.T) {
	t.Run("successful load", func(t *testing.T) {
		clearConfigEnv()
		defer clearConfigEnv()

		os.Setenv("PORT", "8080")

		cfg := MustLoad()
		require.NotNil(t, cfg)
		assert.Equal(t, 8080, cfg.Port)
	})
}

// clearConfigEnv removes all config-related environment variables
func clearConfigEnv() {
	os.Unsetenv("PORT")
	os.Unsetenv("MORESLEEP_URL")
	os.Unsetenv("MORESLEEP_USER")
	os.Unsetenv("MORESLEEP_PASSWORD")
	os.Unsetenv("ELASTICSEARCH_URL")
	os.Unsetenv("PRIVATE_INDEX")
	os.Unsetenv("PUBLIC_INDEX")
}
