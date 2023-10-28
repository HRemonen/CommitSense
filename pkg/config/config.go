/*
Package config provides functionality for reading, creating and modifying configuration files for CommitSense.

This file includes utility functions for interacting with configuration files.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package config

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

// UserHomeDir represents the path to the user's home directory.
var UserHomeDir string

var (
	configFile         = ".commitsense.yaml"
	defaultCommitTypes = []string{"feat", "fix", "chore", "docs", "style", "refactor", "perf", "test", "build", "ci"}
	defaultSkipCITypes = []string{"docs"}
)

// Config represents the configuration settings for the application.
type Config struct {
	CommitTypes []string `yaml:"commit_types"`
	SkipCITypes []string `yaml:"skip_ci_types"`
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting the user's home directory: ", err)
		os.Exit(1)
	}

	UserHomeDir = homeDir
}

// Exists checks if the configuration file exists in the user's home directory.
func Exists() bool {
	if fi, err := os.Stat(path.Join(UserHomeDir, configFile)); err != nil || fi.IsDir() {
		return false
	}
	return true
}

// ReadConfigFile reads the configuration file from the user's home directory.
func ReadConfigFile() (*Config, error) {
	viper.AddConfigPath(UserHomeDir)
	viper.SetConfigType("yaml")
	viper.SetConfigName(configFile)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Error config file not found: ", err)
			os.Exit(1)
		}
		// Config file was found but another error was produced
		fmt.Println("Error reading config file: ", err)
		os.Exit(1)
	}

	fmt.Println("Using config file: ", viper.ConfigFileUsed())

	return &Config{
		CommitTypes: viper.GetStringSlice("commit_types"),
		SkipCITypes: viper.GetStringSlice("skip_ci_types"),
	}, nil
}

// WriteConfigFile writes the configuration file to the user's home directory.
func WriteConfigFile(config *Config) error {
	viper.SetConfigFile(path.Join(UserHomeDir, configFile))
	viper.SetConfigType("yaml")
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
