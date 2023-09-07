package validators

import "fmt"

func ValidateStringNotEmpty(s string) error {
	if len(s) > 0 {
		return nil
	}
	return fmt.Errorf("please enter a valid string")
}

func ValidateStringYesNo(s string) error {
	if s == "Y" || s == "N" || s == "y" || s == "n" {
		return nil
	}
	return fmt.Errorf("please enter Y or N")
}
