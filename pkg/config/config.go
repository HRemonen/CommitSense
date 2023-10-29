/*
Package config provides functionality for reading, creating and modifying configuration files for CommitSense.

This file includes utility functions for interacting with configuration files.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package config

import (
	"fmt"
	"os"
	"strings"

	colorprinter "commitsense/pkg/printer"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// UserHomeDir represents the path to the user's home directory.
// ApplicationConfig represents the configuration settings for the application.
var (
	UserHomeDir       string
	ApplicationConfig *Config
)

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
		colorprinter.ColorPrint("error", "Error getting the user's home directory: %v", err)
		os.Exit(1)
	}

	config, err := ReadConfigFile()
	if err != nil {
		os.Exit(1)
	}

	UserHomeDir = homeDir
	ApplicationConfig = config
}

// Exists checks if the configuration file exists in the user's home directory.
func Exists() bool {
	if fi, err := os.Stat(configFile); err != nil || fi.IsDir() {
		return false
	}
	return true
}

// ReadConfigFile reads the configuration file from the user's home directory.
func ReadConfigFile() (*Config, error) {
	viper.SetConfigFile(configFile)
	viper.AddConfigPath(UserHomeDir)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			colorprinter.ColorPrint("error", "Error config file not found: %v", err)
			return nil, err
		}
		// Config file was found but another error was produced
		colorprinter.ColorPrint("error", "Error reading config file: %v", err)
		return nil, err
	}

	return &Config{
		CommitTypes: viper.GetStringSlice("commit_types"),
		SkipCITypes: viper.GetStringSlice("skip_ci_types"),
	}, nil
}

// WriteConfigFile writes the configuration file to the user's home directory.
func WriteConfigFile(config *Config) error {
	viper.SetConfigFile(configFile)
	viper.AddConfigPath(UserHomeDir)

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

// ShowConfigSettings prints out the current configuration settings.
func ShowConfigSettings() error {
	colorprinter.ColorPrint("success", "\nShowing current configuration settings")
	config := ApplicationConfig

	colorprinter.ColorPrint("info", "Using configuration file: %v", configFile)

	colorprinter.ColorPrint("bold", "\nAllowed commit types:")
	printYAML(config.CommitTypes)

	colorprinter.ColorPrint("bold", "Skipping CI on types:")
	printYAML(config.SkipCITypes)

	return nil
}

func printYAML(data interface{}) {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		colorprinter.ColorPrint("error", "Error printing YAML: %v", err)

		return
	}

	// Use strings.Replace to add proper indentation
	indentedYAML := strings.ReplaceAll(string(yamlData), "\n", "\n  ")
	fmt.Println("  " + indentedYAML)
}
