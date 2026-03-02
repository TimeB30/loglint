package rules

import "testing"

func TestLowerCaseRule(t *testing.T) {
	lcr := LowerCaseRule{}

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "начинается с маленькой латинской буквы",
			input:    "hello world",
			expected: true,
		},
		{
			name:     "одна маленькая латинская буква",
			input:    "a",
			expected: true,
		},
		{
			name:     "начинается с маленькой русской буквы",
			input:    "привет мир",
			expected: true,
		},
		{
			name:     "начинается с маленькой немецкой буквы",
			input:    "münchen ist schön",
			expected: true,
		},
		{
			name:     "начинается с маленькой французской буквы",
			input:    "élève français",
			expected: true,
		},
		{
			name:     "начинается с маленькой греческой буквы",
			input:    "αλφα",
			expected: true,
		},
		{
			name:     "маленькая буква с последующими цифрами",
			input:    "a12345",
			expected: true,
		},
		{
			name:     "маленькая буква с последующими спецсимволами",
			input:    "a!@#$%",
			expected: true,
		},

		{
			name:     "начинается с большой латинской буквы",
			input:    "Hello world",
			expected: false,
		},
		{
			name:     "начинается с большой русской буквы",
			input:    "Привет мир",
			expected: false,
		},
		{
			name:     "начинается с большой немецкой буквы",
			input:    "München",
			expected: false,
		},
		{
			name:     "одна большая буква",
			input:    "A",
			expected: false,
		},
		{
			name:     "начинается с цифры",
			input:    "123 hello",
			expected: false,
		},
		{
			name:     "начинается с пробела",
			input:    " hello",
			expected: false,
		},
		{
			name:     "начинается с табуляции",
			input:    "\thello",
			expected: false,
		},
		{
			name:     "начинается с новой строки",
			input:    "\nhello",
			expected: false,
		},
		{
			name:     "начинается со спецсимвола",
			input:    "@hello",
			expected: false,
		},
		{
			name:     "начинается с эмодзи",
			input:    "😊 hello",
			expected: false,
		},
		{
			name:     "начинается с китайского иероглифа",
			input:    "你好世界",
			expected: false,
		},
		{
			name:     "пустая строка",
			input:    "",
			expected: false,
		},
		{
			name:     "только пробелы",
			input:    "   ",
			expected: false,
		},
		{
			name:     "буква верхнего регистра в середине",
			input:    "hEllo",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lcr.Check(&tt.input)
			if result != tt.expected {
				t.Errorf("Check(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
