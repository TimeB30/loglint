package rules

import "unicode"

type EnglishRule struct{}

func (er EnglishRule) Check(s *string) bool {
	for _, r := range *s {
		if unicode.IsLetter(r) {
			if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') {
				return false
			}

		}
	}
	return true
}
func (er EnglishRule) Name() string {
	return "Message should be only in english"
}
