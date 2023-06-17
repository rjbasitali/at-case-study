package validate

func Locale(s string) bool {
	if s == "en_ae" || s == "ar_ae" {
		return true
	}
	return false
}
