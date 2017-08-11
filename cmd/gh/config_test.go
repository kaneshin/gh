package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Clearenv()
}

func testMustNewConfig(t *testing.T) *Config {
	conf, err := NewConfig("testdata/")
	if assert.NoError(t, err) {
		if assert.NotNil(t, conf) {
			conf.WithEnv("development")
		}
	}
	return conf
}

func Test_NewConfig(t *testing.T) {
	t.Run("No Exists", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		tests := []struct {
			name string
		}{
			{""},
			{"testdata/noexists.tml"},
		}

		for _, tt := range tests {
			tt := tt
			conf, err := NewConfig(tt.name)
			assert.Nil(conf)
			assert.Error(err)
		}
	})

	t.Run("Unsupported Extension", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		tests := []struct {
			name string
		}{
			{""},
			{"testdata/configtml"},
			{"testdata/config.txt"},
		}

		for _, tt := range tests {
			tt := tt
			conf, err := NewConfig(tt.name)
			assert.Nil(conf)
			assert.Error(err)
			assert.Equal(ErrUnsupportedFile, err)
		}
	})

	t.Run("Normally", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		tests := []struct {
			name        string
			env         string
			owner       string
			repo        string
			accessToken string
		}{
			{
				name:        "testdata/config.tml",
				env:         "gh",
				owner:       "kaneshin",
				repo:        "gh",
				accessToken: "GITHUB_ACCESS_TOKEN",
			},
			{
				name:        "testdata/config.yml",
				env:         "test",
				owner:       "kaneshin",
				repo:        "test",
				accessToken: "GITHUB_ACCESS_TOKEN",
			},
		}

		for _, tt := range tests {
			tt := tt
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				conf, err := NewConfig(tt.name)
				assert.NotNil(conf)
				assert.NoError(err)

				conf.WithEnv(tt.env)
				assert.Equal(tt.owner, conf.Owner())
				assert.Equal(tt.repo, conf.Repo())
				assert.Equal(tt.accessToken, conf.AccessToken())
			})
		}
	})
}
