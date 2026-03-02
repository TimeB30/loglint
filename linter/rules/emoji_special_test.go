package rules

import "testing"

func TestEmojiSpecialRule(t *testing.T) {
	esr := EmojiSpecialRule{}

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "только буквы",
			input:    "HelloWorld",
			expected: true,
		},
		{
			name:     "буквы и цифры",
			input:    "Hello123World456",
			expected: true,
		},
		{
			name:     "буквы, цифры и пробелы",
			input:    "Hello 123 World 456",
			expected: true,
		},
		{
			name:     "специальный символ @",
			input:    "Hello@World",
			expected: false,
		},
		{
			name:     "специальный символ #",
			input:    "Hello#World",
			expected: false,
		},
		{
			name:     "специальный символ $",
			input:    "Hello$World",
			expected: false,
		},
		{
			name:     "специальный символ %",
			input:    "Hello%World",
			expected: false,
		},
		{
			name:     "специальный символ ^",
			input:    "Hello^World",
			expected: false,
		},
		{
			name:     "специальный символ &",
			input:    "Hello&World",
			expected: false,
		},
		{
			name:     "специальный символ *",
			input:    "Hello*World",
			expected: false,
		},
		{
			name:     "специальный символ +",
			input:    "Hello+World",
			expected: false,
		},
		{
			name:     "эмодзи",
			input:    "Hello 😊 World",
			expected: false,
		},
		{
			name:     "несколько эмодзи",
			input:    "Hello 😊 🎉 World",
			expected: false,
		},
		{
			name:     "пустая строка",
			input:    "",
			expected: true,
		},
		{
			name:     "только пробелы",
			input:    "   ",
			expected: true,
		},
		{
			name:     "русские буквы",
			input:    "ПриветМир",
			expected: true,
		},
		{
			name:     "китайские иероглифы",
			input:    "你好世界",
			expected: true,
		},
		{
			name:     "смешанные алфавиты",
			input:    "Hello Привет 123",
			expected: true,
		},
		{
			name:     "знаки пунктуации",
			input:    "Hello, World!",
			expected: false,
		},
		{
			name:     "дефис",
			input:    "Hello-World",
			expected: false,
		},
		{
			name:     "подчеркивание",
			input:    "Hello_World",
			expected: false,
		},
		{
			name:     "точка",
			input:    "Hello.World",
			expected: false,
		},
		{
			name:     "цифры в начале",
			input:    "123Hello",
			expected: true,
		},
		{
			name:     "только цифры",
			input:    "123456",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := esr.Check(&tt.input)
			if result != tt.expected {
				t.Errorf("Check(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
