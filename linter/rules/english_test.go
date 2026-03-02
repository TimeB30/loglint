package rules

import "testing"

func TestEnglishRule(t *testing.T) {
	er := EnglishRule{}

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "только английские буквы (нижний регистр)",
			input:    "hello",
			expected: true,
		},
		{
			name:     "только английские буквы (верхний регистр)",
			input:    "HELLO",
			expected: true,
		},
		{
			name:     "смешанный регистр",
			input:    "HelloWorld",
			expected: true,
		},
		{
			name:     "с пробелами",
			input:    "Hello World",
			expected: true,
		},
		{
			name:     "с цифрами",
			input:    "Hello123World456",
			expected: true,
		},
		{
			name:     "с пробелами и цифрами",
			input:    "Hello 123 World 456",
			expected: true,
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
			input:    "Привет",
			expected: false,
		},
		{
			name:     "смесь английских и русских букв",
			input:    "HelloПривет",
			expected: false,
		},
		{
			name:     "китайские иероглифы",
			input:    "你好",
			expected: false,
		},
		{
			name:     "немецкие буквы с умлаутом",
			input:    "München",
			expected: false,
		},
		{
			name:     "французские буквы с акцентами",
			input:    "élève",
			expected: false,
		},
		{
			name:     "испанская ñ",
			input:    "España",
			expected: false,
		},
		{
			name:     "датская буква ø",
			input:    "Cøpenhagen",
			expected: false,
		},
		{
			name:     "специальные символы",
			input:    "Hello@World",
			expected: true,
		},
		{
			name:     "эмодзи",
			input:    "Hello 😊",
			expected: true,
		},
		{
			name:     "знаки пунктуации",
			input:    "Hello, World!",
			expected: true,
		},
		{
			name:     "одиночная английская буква",
			input:    "a",
			expected: true,
		},
		{
			name:     "одиночная русская буква",
			input:    "а",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := er.Check(&tt.input)
			if result != tt.expected {
				t.Errorf("Check(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
