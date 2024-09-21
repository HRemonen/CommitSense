/*
Package config provides functionality for reading, creating and modifying configuration files for CommitSense.

This file includes utility functions for interacting with configuration files.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package config

import (
	"encoding/json"
	"fmt"
	"os"

	colorprinter "commitsense/pkg/printer"

	"github.com/spf13/viper"
)

var (
	configFile         *Config
	configFileName     = "commitsense.config.json"
	defaultVersion     = 1
	defaultCommitTypes = []string{"feat", "fix", "docs", "style", "refactor", "perf", "test", "build", "ci", "chore", "revert"}
	defaultSkipCITypes = []string{"docs"}
)

// Config represents the configuration settings for the application.
type Config struct {
	Version     int      `json:"version"`
	CommitTypes []string `json:"commit_types"`
	SkipCITypes []string `json:"skip_ci_types"`
}

func NewDefault() *Config {
	return &Config{
		Version:     defaultVersion,
		CommitTypes: defaultCommitTypes,
		SkipCITypes: defaultSkipCITypes,
	}
}


func init() {
	if !Exists() {
		cfg := NewDefault()

		err := Write(cfg)
		if err != nil {
			colorprinter.ColorPrint("error", "Error creating default config: %v", err)
			os.Exit(1)
		}

		colorprinter.ColorPrint("info", "\nCould not find an existing configuration file")
		colorprinter.ColorPrint("success", "Created default configuration file at %v", configFileName)
	}

	config, err := Read()
	if err != nil {
		os.Exit(1)
	}

	configFile = config
}

// Exists checks if the configuration file exists in the project's root directory.
func Exists() bool {
	if fi, err := os.Stat(configFileName); err != nil || fi.IsDir() {
		return false
	}
	return true
}

// ReadConfigFile reads the configuration file from the project's root directory.
func Read() (*Config, error) {
	viper.SetConfigFile(configFileName)

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
		Version:     viper.GetInt("version"),
		CommitTypes: viper.GetStringSlice("commit_types"),
		SkipCITypes: viper.GetStringSlice("skip_ci_types"),
	}, nil
}

// WriteConfigFile writes the configuration file to the project's root directory.
func Write(config *Config) error {
	viper.SetConfigFile(configFileName)

	viper.Set("version", config.Version)
	viper.Set("commit_types", config.CommitTypes)
	viper.Set("skip_ci_types", config.SkipCITypes)

	return viper.WriteConfig()
}

// ShowConfigSettings prints out the current configuration settings.
func Show() error {
	colorprinter.ColorPrint("success", "\nShowing current configuration settings")
	config := configFile

	colorprinter.ColorPrint("info", "Using configuration file: %v", configFileName)

	colorprinter.ColorPrint("bold", "\nAllowed commit types:")
	printConfig(config.CommitTypes)

	colorprinter.ColorPrint("bold", "Skipping CI on types:")
	printConfig(config.SkipCITypes)

	return nil
}

func printConfig(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		colorprinter.ColorPrint("error", "Error printing YAML: %v", err)

		return
	}

	fmt.Println(string(jsonData))
}
