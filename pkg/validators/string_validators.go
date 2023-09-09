/*
Package validators provides a collection of utility functions for validating various data types.

This package includes functions for validating different types of data, such as strings, numbers, and more. It serves as a library of validation functions that can be used across various applications.

Usage:
  - Explore the available validation functions to find the one that suits your validation needs.
  - Call the relevant validation function to ensure the integrity of your data.

Copyright © 2023 HENRI REMONEN <henri@remonen.fi>
*/
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
