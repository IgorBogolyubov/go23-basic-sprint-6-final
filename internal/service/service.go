package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Service(in string) string {
	ch := "ЙЦУКЕНГШЩЗХЪФЫВАПРОЛДЖЭЯЧСМИТЬБЮЁйцукенгшщзхъфывапролджэячсмитьбю"
	out := ""
	if strings.ContainsAny(in, ch) {
		for _, val := range in {
			out += morse.RuneToMorse(rune(val)) + " "
		}
	} else {
		for _, val := range strings.Split(in, " ") {
			out += string(morse.MorseToRune(val))

		}
	}

	return out
}
