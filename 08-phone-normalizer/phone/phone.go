package phone

import "bytes"

func Normalize(phoneNumber string) string {
	var buf bytes.Buffer

	for _, char := range phoneNumber {
		if isDigit(char) {
			buf.WriteRune(char)
		}
	}

	return buf.String()
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
