package utils

import "regexp"

// Fungsi validasi email regex
func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`\S+@\S+\.\S+`)
	return re.MatchString(email)
}
