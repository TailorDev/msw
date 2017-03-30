// Package config contains the code to read user's configuration.
// Inspired by https://benaiah.me/posts/configuring-go-apps-with-toml/
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
)

type Config struct {
	Buffer BufferOptions
}

type BufferOptions struct {
	AccessToken string
	ProfileIDs  []string
	BufferSize  int
}

var DefaultConfig = Config{
	Buffer: BufferOptions{
		BufferSize: 10,
	},
}

var configDirName = "msw"

func GetDefaultConfigDir() (string, error) {
	var configDirLocation string

	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	switch runtime.GOOS {
	case "linux":
		// Use the XDG_CONFIG_HOME variable if it is set, otherwise
		// $HOME/.config/example
		xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfigHome != "" {
			configDirLocation = xdgConfigHome
		} else {
			configDirLocation = filepath.Join(homeDir, ".config", configDirName)
		}

	default:
		// On other platforms we just use $HOME/.example
		hiddenConfigDirName := "." + configDirName
		configDirLocation = filepath.Join(homeDir, hiddenConfigDirName)
	}

	return configDirLocation, nil
}

func LoadConfig(configFile string) (*Config, error) {
	conf := DefaultConfig

	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return &conf, fmt.Errorf("configuration file is missing (%s).", configFile)
	} else if err != nil {
		return &conf, err
	}

	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		return &conf, err
	}

	return &conf, nil
}

func LoadDefaultConfig() (*Config, error) {
	dir, err := GetDefaultConfigDir()
	if err != nil {
		return nil, err
	}

	return LoadConfig(filepath.Join(dir, "msw.toml"))
}
