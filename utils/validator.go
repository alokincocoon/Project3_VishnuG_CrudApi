package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/vishnu-g-k/student-management/models"
)

func ValidateEmail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func ValidatePhone(value string) bool {
	// Function for validating one phone number

	phRegex := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	return phRegex.MatchString(value)
}

func ValidatePhones(phones []map[string]interface{}) error {
	if len(phones) < 1 {
		return errors.New("at least 1 phone number is required")
	}
	invalidPhones := []string{}
	for _, phone := range phones {
		p := phone["number"].(string)
		if !ValidatePhone(p) {
			invalidPhones = append(invalidPhones, p)
		}
	}
	if len(invalidPhones) > 0 {
		return fmt.Errorf("invalid phone numbers '%s'", strings.Join(invalidPhones, ", "))
	}
	return nil
}

func ValidateStudent(s *models.Student) error {
	if s.Email == "" {
		return errors.New("email id is required")
	}
	if !ValidateEmail(s.Email) {
		return fmt.Errorf("invalid email id '%s'", s.Email)
	}
	if err := ValidatePhones(s.Phones); err != nil {
		return err
	}

	return nil
}
