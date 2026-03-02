package rules

import "testing"

func TestSensitiveRule(t *testing.T) {
	sr := SensitiveRule{}

	tests := []struct {
		name     string
		input    string
		expected bool
	}{

		{
			name:     "отдельное слово 'password'",
			input:    "password",
			expected: false,
		},
		{
			name:     "отдельное слово 'pass'",
			input:    "pass",
			expected: false,
		},
		{
			name:     "отдельное слово 'pwd'",
			input:    "pwd",
			expected: false,
		},
		{
			name:     "отдельное слово 'token'",
			input:    "token",
			expected: false,
		},
		{
			name:     "отдельное слово 'api_key'",
			input:    "api_key",
			expected: false,
		},
		{
			name:     "отдельное слово 'secret'",
			input:    "secret",
			expected: false,
		},
		{
			name:     "отдельное слово 'key'",
			input:    "key",
			expected: false,
		},
		{
			name:     "отдельное слово 'credential'",
			input:    "credential",
			expected: false,
		},
		{
			name:     "отдельное слово 'auth'",
			input:    "auth",
			expected: false,
		},
		{
			name:     "слово в предложении",
			input:    "my password is 12345",
			expected: false,
		},
		{
			name:     "слово с пунктуацией в конце",
			input:    "password!",
			expected: false,
		},
		{
			name:     "слово с пунктуацией в начале",
			input:    "!password",
			expected: false,
		},
		{
			name:     "слово в верхнем регистре",
			input:    "PASSWORD",
			expected: false,
		},
		{
			name:     "слово в смешанном регистре",
			input:    "PaSsWoRd",
			expected: false,
		},
		{
			name:     "несколько чувствительных слов",
			input:    "username: admin, password: 123, token: abc",
			expected: false,
		},
		{
			name:     "слово с подчеркиванием",
			input:    "api_key=12345",
			expected: false,
		},
		{
			name:     "слово в JSON",
			input:    `{"password": "12345", "user": "admin"}`,
			expected: false,
		},
		{
			name:     "слово в URL параметре",
			input:    "https://example.com?api_key=12345&token=abc",
			expected: false,
		},

		{
			name:     "пустая строка",
			input:    "",
			expected: true,
		},
		{
			name:     "обычный текст без чувствительных слов",
			input:    "Hello world, how are you?",
			expected: true,
		},
		{
			name:     "слова, содержащие 'pass' как часть",
			input:    "passage passport passwordless",
			expected: true,
		},
		{
			name:     "слова, содержащие 'key' как часть",
			input:    "monkey donkey turkey keyboard",
			expected: true,
		},
		{
			name:     "слова, содержащие 'auth' как часть",
			input:    "author authentic authority",
			expected: true,
		},
		{
			name:     "слова, содержащие 'secret' как часть",
			input:    "secretary secrete",
			expected: true,
		},
		{
			name:     "только пробелы",
			input:    "   ",
			expected: true,
		},
		{
			name:     "только цифры",
			input:    "12345",
			expected: true,
		},
		{
			name:     "только спецсимволы",
			input:    "!@#$%^&*()",
			expected: true,
		},
		{
			name:     "длинный текст без чувствительных слов",
			input:    "The quick brown fox jumps over the lazy dog. The weather is nice today. Let's go for a walk in the park.",
			expected: true,
		},
		{
			name:     "текст с цифрами и спецсимволами без чувствительных слов",
			input:    "Total price: $123.45, discount: 10%",
			expected: true,
		},
		{
			name:     "текст с email без чувствительных слов",
			input:    "Contact us at support@example.com for help",
			expected: true,
		},
		{
			name:     "текст с URL без чувствительных слов",
			input:    "Visit our website: https://example.com",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sr.Check(&tt.input)
			if result != tt.expected {
				t.Errorf("Check(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
