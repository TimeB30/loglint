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
			expected: true,
		},
		{
			name:     "отдельное слово 'pass'",
			input:    "pass",
			expected: true,
		},
		{
			name:     "отдельное слово 'pwd'",
			input:    "pwd",
			expected: true,
		},
		{
			name:     "отдельное слово 'token'",
			input:    "token",
			expected: true,
		},
		{
			name:     "отдельное слово 'api_key'",
			input:    "api_key",
			expected: true,
		},
		{
			name:     "отдельное слово 'secret'",
			input:    "secret",
			expected: true,
		},
		{
			name:     "отдельное слово 'key'",
			input:    "key",
			expected: true,
		},
		{
			name:     "отдельное слово 'credential'",
			input:    "credential",
			expected: true,
		},
		{
			name:     "отдельное слово 'auth'",
			input:    "auth",
			expected: true,
		},
		{
			name:     "слово в предложении",
			input:    "my password is 12345",
			expected: true,
		},
		{
			name:     "слово с пунктуацией в конце",
			input:    "password!",
			expected: true,
		},
		{
			name:     "слово с пунктуацией в начале",
			input:    "!password",
			expected: true,
		},
		{
			name:     "слово в верхнем регистре",
			input:    "PASSWORD",
			expected: true,
		},
		{
			name:     "слово в смешанном регистре",
			input:    "PaSsWoRd",
			expected: true,
		},
		{
			name:     "несколько чувствительных слов",
			input:    "username: admin, password: 123, token: abc",
			expected: true,
		},
		{
			name:     "слово с подчеркиванием",
			input:    "api_key=12345",
			expected: true,
		},
		{
			name:     "слово в JSON",
			input:    `{"password": "12345", "user": "admin"}`,
			expected: true,
		},
		{
			name:     "слово в URL параметре",
			input:    "https://example.com?api_key=12345&token=abc",
			expected: true,
		},

		{
			name:     "пустая строка",
			input:    "",
			expected: false,
		},
		{
			name:     "обычный текст без чувствительных слов",
			input:    "Hello world, how are you?",
			expected: false,
		},
		{
			name:     "слова, содержащие 'pass' как часть",
			input:    "passage passport passwordless",
			expected: false,
		},
		{
			name:     "слова, содержащие 'key' как часть",
			input:    "monkey donkey turkey keyboard",
			expected: false,
		},
		{
			name:     "слова, содержащие 'auth' как часть",
			input:    "author authentic authority",
			expected: false,
		},
		{
			name:     "слова, содержащие 'secret' как часть",
			input:    "secretary secrete",
			expected: false,
		},
		{
			name:     "только пробелы",
			input:    "   ",
			expected: false,
		},
		{
			name:     "только цифры",
			input:    "12345",
			expected: false,
		},
		{
			name:     "только спецсимволы",
			input:    "!@#$%^&*()",
			expected: false,
		},
		{
			name:     "длинный текст без чувствительных слов",
			input:    "The quick brown fox jumps over the lazy dog. The weather is nice today. Let's go for a walk in the park.",
			expected: false,
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
