/*
Package colorprinter provides functionality for printing colored stdout messages for CommitSense.

This file includes utility functions for interacting with colored stdout messages.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
package colorprinter

import (
	"github.com/fatih/color"
)

var (
	successColor = color.New(color.FgGreen).Add(color.Bold)
	infoColor    = color.New(color.FgCyan).Add(color.Bold)
	errorColor   = color.New(color.FgRed).Add(color.Bold)
	boldColor    = color.New(color.Bold)
)

// ColorPrint prints out a colored messaged according to the variant given as parameter.
func ColorPrint(variant string, text string) {
	var printer *color.Color

	switch variant {
	case "success":
		printer = successColor
	case "info":
		printer = infoColor
	case "error":
		printer = errorColor
	case "bold":
		printer = boldColor
	default:
		printer = color.New()
	}

	printer.Println(text)
}
