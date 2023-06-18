package validate

import "testing"

func TestLocale(t *testing.T) {
	locales := []struct {
		locale  string
		isValid bool
	}{
		{"en_ae", true},
		{"ar_ae", true},
		{"en_US", false},
		{"en-us", false},
		{"en", false},
		{"ae", false},
		{"", false},
	}

	for _, l := range locales {
		if Locale(l.locale) != l.isValid {
			t.Errorf("Locale %s should be %t", l.locale, l.isValid)
		}
	}
}
