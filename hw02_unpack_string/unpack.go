package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var (
		arrayRune  []rune
		runeLast   rune
		strUnzip   string
		runeBefore bool
	)
	arrayRune = []rune(str)
	for _, value := range arrayRune {
		// проверяем отрицательный кейс, что используется цифра, но не число
		if runeBefore && unicode.IsDigit(value) {
			return "", ErrInvalidString
		}

		// проверяем что строка не начинается с цифры
		if unicode.IsDigit(arrayRune[0]) {
			return "", ErrInvalidString
		}

		// записываем для того, чтобы проверить, что будет цифра, а не число
		runeBefore = unicode.IsDigit(value)

		// если не цифра и не символ
		if !unicode.IsDigit(value) || !unicode.IsLetter(runeLast) {
			// если число или цифра
			if unicode.IsLetter(runeLast) && unicode.IsLetter(value) {
				strUnzip += string(runeLast)
			}
			runeLast = value
			continue
		}
		// умножаем строку
		strUnzip += strings.Repeat(string(runeLast), int(value-'0'))
		runeLast = value
	}
	// если символ пишем просто так
	if unicode.IsLetter(runeLast) {
		strUnzip += string(runeLast)
	}
	return strUnzip, nil
}
