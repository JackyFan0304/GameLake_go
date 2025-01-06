package utils

import (
	"regexp"
)

// ValidateEmail checks if the provided email is valid.
func ValidateEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// HashPassword hashes the provided password using a suitable hashing algorithm.
func HashPassword(password string) (string, error) {
	// Implementation for hashing the password
	return "", nil
}

// ComparePasswords compares a hashed password with a plain password.
func ComparePasswords(hashedPassword, plainPassword string) bool {
	// Implementation for comparing passwords
	return false
}

// FormatCurrency formats a float64 value to a currency string.
func FormatCurrency(amount float64) string {
	// Implementation for formatting currency
	return ""
}