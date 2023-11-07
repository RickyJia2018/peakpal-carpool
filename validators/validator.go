package validators

import (
	"fmt"
	"net/mail"
	"reflect"
	"regexp"
	"time"
)

var (
	// isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidName   = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isValidNumber = regexp.MustCompile(`^[0-9]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateNumber(value int64, min int64, max int64) error {
	if value < min || value > max {
		return fmt.Errorf("must be between %d and %d", min, max)
	}
	return nil
}
func ValidateTime(value time.Time, min time.Time, max time.Time) error {
	if value.Before(min) || value.After(max) {
		return fmt.Errorf("must be between %s and %s", min, max)
	}
	return nil
}
func ValidateBool(value interface{}) error {
	valueType := reflect.TypeOf(value)
	if valueType.Kind() != reflect.Bool {
		return fmt.Errorf("must be boolean")
	}
	return nil
}

func ValidatePhone(value string) error {
	if err := ValidateString(value, 9, 10); err != nil {
		return err
	}
	if !isValidNumber(value) {
		return fmt.Errorf("must contain only numbers")
	}
	return nil
}

func ValidateName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidName(value) {
		return fmt.Errorf("must contain only letters or spaces")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}

func ValidateID(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be a positive integer")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}
