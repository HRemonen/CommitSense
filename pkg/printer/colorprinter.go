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