package rules

import (
	"strings"
	"unicode"
)

type SensitiveRule struct{}

func (sr SensitiveRule) Check(s *string) bool {
	if len(*s) == 0 {
		return true
	}
	sensitiveWords := []string{
		"password", "pass", "pwd", "token", "api_key",
		"secret", "key", "credential", "auth",
	}
	sensitiveMap := make(map[string]bool)
	for _, word := range sensitiveWords {
		sensitiveMap[word] = true
	}
	words := strings.FieldsFunc(*s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_'
	})

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		if sensitiveMap[lowerWord] {
			return false
		}
	}

	return true

}

func (sr SensitiveRule) Name() string {
	return "Sensitive data isn't allowed"
}
