package helper

import (
	"os"
	"strings"
)

func IsFileExists(filename string) bool {

	info, err := os.Stat(filename)

	if os.IsNotExist(err) {

	   return false
	}

	return !info.IsDir()
}

func CapitalizeFirstLetter(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + word[1:]
}