package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func IsMorse(text string) string {
	var isMorse bool
	if text == "" {
		return "пустая строка"
	}
	for _, r := range text {
		if r == ' ' || r == '-' || r == '.' {
			isMorse = true
		} else {
			isMorse = false
		}

	}

	if isMorse {
		if morse.ToText(text) == "" {
			return "Невозможно расшифровать код Морзе"
		} else {
			return morse.ToText(text)
		}
	}
	return morse.ToMorse(text)

}
