package rules

import "unicode"

type EmojiSpecialRule struct{}

func (esr EmojiSpecialRule) Check(s *string) bool {
	for _, r := range *s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func (esr EmojiSpecialRule) Name() string {
	return "Emoji's and special characters doesn't allowed"
}
