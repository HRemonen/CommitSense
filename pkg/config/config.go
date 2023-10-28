package config

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

var UserHomeDir string

var (
	configFile         = "commitsense.toml"
	defaultCommitTypes = []string{"feat", "fix", "chore", "docs", "style", "refactor", "perf", "test", "build", "ci"}
	defaultSkipCITypes = []string{"docs"}
)

// Config represents the configuration settings for the application.
type Config struct {
	CommitTypes []string `toml:"commit_types"`
	SkipCITypes []string `toml:"skip_ci_types"`
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting the user's home directory: ", err)
		os.Exit(1)
	}

	UserHomeDir = homeDir
}

// ConfigFileExists checks if the configuration file exists in the user's home directory.
func ConfigFileExists() bool {
	if fi, err := os.Stat(path.Join(UserHomeDir, configFile)); err != nil || fi.IsDir() {
		return false
	}
	return true
}

// ReadConfigFile reads the configuration file from the user's home directory.
func ReadConfigFile() (*Config, error) {
	viper.SetConfigFile(path.Join(UserHomeDir, configFile))
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		CommitTypes: viper.GetStringSlice("commit_types"),
		SkipCITypes: viper.GetStringSlice("skip_ci_types"),
	}, nil
}

// WriteConfigFile writes the configuration file to the user's home directory.
func WriteConfigFile(config *Config) error {
	viper.SetConfigFile(path.Join(UserHomeDir, configFile))
	viper.SetConfigType("toml")
	viper.Set("commit_types", config.CommitTypes)
	viper.Set("skip_ci_types", config.SkipCITypes)

	return viper.WriteConfig()
}

// CreateDefaultConfig writes a default configuration file to the user's home directory.
func CreateDefaultConfig() error {
	return WriteConfigFile(&Config{
		CommitTypes: defaultCommitTypes,
		SkipCITypes: defaultSkipCITypes,
	})
} 