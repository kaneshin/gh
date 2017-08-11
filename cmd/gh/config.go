package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type (
	// Config object.
	Config struct {
		data map[string]internal
		env  string
	}

	internal struct {
		Owner       string `toml:"owner" yaml:"owner"`
		Repo        string `toml:"repo" yaml:"repo"`
		AccessToken string `toml:"access_token" yaml:"access_token"`
	}
)

var (
	ErrUnsupportedFile = errors.New("unsupported file")
)

// MustNewConfig returns a new config. name cannot be empty.
func MustNewConfig(name string) *Config {
	c, err := NewConfig(name)
	if err != nil {
		panic(err)
	}
	return c
}

// NewConfig returns a new config. name cannot be empty.
func NewConfig(name string) (*Config, error) {
	fpath := filepath.Clean(name)
	conf := Config{}

	switch filepath.Ext(fpath) {
	case ".yml", ".yaml":
		b, err := ioutil.ReadFile(fpath)
		if err != nil {
			return nil, err
		}
		if err := yaml.Unmarshal(b, &conf.data); err != nil {
			return nil, err
		}

	case ".tml", ".toml":
		if _, err := toml.DecodeFile(fpath, &conf.data); err != nil {
			return nil, err
		}

	default:
		return nil, ErrUnsupportedFile
	}

	return &conf, nil
}

// MergeConfig returns merged config.
func MergeConfig(c ...*Config) *Config {
	conf := Config{data: map[string]internal{}}
	for _, v := range c {
		if conf.env == "" && v.env != "" {
			conf.env = v.env
		}

		for key, vv := range v.data {
			if _, ok := conf.data[key]; !ok {
				conf.data[key] = vv
			}
		}
	}
	return &conf
}

// WithEnv sets an environment of config.
func (c *Config) WithEnv(env string) *Config {
	c.env = env
	return c
}

// Owner returns a raw owner string.
func (c Config) Owner() string {
	if d, ok := c.data[c.env]; ok {
		return os.ExpandEnv(d.Owner)
	}
	return ""
}

// Repo returns a raw repo string.
func (c Config) Repo() string {
	if d, ok := c.data[c.env]; ok {
		return os.ExpandEnv(d.Repo)
	}
	return ""
}

// AccessToken returns a raw access token string.
func (c Config) AccessToken() string {
	if d, ok := c.data[c.env]; ok {
		return os.ExpandEnv(d.AccessToken)
	}
	return ""
}
