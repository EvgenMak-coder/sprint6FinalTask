package service

import (
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Encode(text string) string {

	var str morse.ErrNoEncoding

	if text == "" {
		str.Text = text
		return str.Error()
	}

	containsProhibited := strings.ContainsFunc(text, func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsNumber(r)) && (r == ' ' || r == '-' || r == '.')
	})

	if containsProhibited {
		if morse.ToText(text) == "" {
			str.Text = text
			return str.Error()
		} else {
			return morse.ToText(text)
		}
	}
	return morse.ToMorse(text)

}
