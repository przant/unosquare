package caesar

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func CaesarCipher(msg string, key int) (cipher []byte, err error) {
	if key < 0 {
		cipher, err = nil, fmt.Errorf("cipher key cannot be lower than 0: %d", key)
		return cipher, err
	}

	for _, char := range msg {
		if unicode.IsLetter(char) {
			switch {
			case unicode.IsLower(char):
				cChar := (char + rune(key)%97)
				cipher = utf8.AppendRune(cipher, cChar)
			case unicode.IsUpper(char):
				cChar := (char + rune(key)%65)
				cipher = utf8.AppendRune(cipher, cChar)
			}
		} else {
			cipher = append(cipher, byte(char))
		}
	}

	return cipher, err
}
