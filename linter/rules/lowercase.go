package rules

import "unicode"

type LowerCaseRule struct{}

func (lcr LowerCaseRule) Check(s *string) bool {
	if len(*s) == 0 {
		return false
	}
	runes := []rune(*s)
	if unicode.IsLetter(runes[0]) && unicode.IsLower(runes[0]) {
		return true
	}
	return false
}

func (lcr LowerCaseRule) Name() string {
	return "Message should start with the lower case letter"
}
