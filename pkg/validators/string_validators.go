package validators

import "fmt"

// ValidateStringNotEmpty checks if a string is not empty.
func ValidateStringNotEmpty(s string) error {
	if len(s) > 0 {
		return nil
	}
	return fmt.Errorf("please enter a valid string")
}

// ValidateStringYesNo checks if a string is "yes" or "no" (case-insensitive).
func ValidateStringYesNo(s string) error {
	if s == "Y" || s == "N" || s == "y" || s == "n" {
		return nil
	}
	return fmt.Errorf("please enter Y or N")
}
