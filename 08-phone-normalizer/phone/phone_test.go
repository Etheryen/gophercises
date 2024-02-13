package phone_test

import (
	"08-phone-normalizer/phone"
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"123-456-7890", "1234567890"},
		{"1234567892", "1234567892"},
		{"(123)456-7892", "1234567892"},
		{"((  -123)--( -456-7892", "1234567892"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got := phone.Normalize(test.input)
			if got != test.want {
				t.Errorf(
					"Normalize(\"%v\") = %v; want %v",
					test.input,
					got,
					test.want,
				)
			}
		})
	}
}
