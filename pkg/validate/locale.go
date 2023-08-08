package validate

// Locale validates the locale.
// It accepts a string as a parameter.
// It returns true if the locale is valid.
// Valid locales are "en_ae" and "ar_ae".
func Locale(s string) bool {
	if s == "en_ae" || s == "ar_ae" {
		return true
	}
	return false
}
