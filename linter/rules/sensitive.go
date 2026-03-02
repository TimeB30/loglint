package rules

import "strings"

type SensitiveRule struct{}

func (sr SensitiveRule) Check(s *string) bool {
	sensitiveWords := []string{
		"password", "pass", "pwd", "token", "api_key",
		"secret", "key", "credential", "auth",
	}
	for _, word := range sensitiveWords {
		if strings.Contains(*s, word) {
			return false
		}
	}
	return true
}

func (sr SensitiveRule) Name() string {
	return "Sensitive data doesn't allowed"
}
