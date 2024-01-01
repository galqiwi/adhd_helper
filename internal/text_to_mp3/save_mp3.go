package text_to_mp3

import (
	"os/exec"
	"unicode"
)

var engAlphabet = "abcdefghijklmnopqrstuvwxyz"
var rusAlphabet = "абвгдеёжзийклмнопрстуфхчшщъыьэюя"

// Inefficient, but good enough
func isRuneFromAlphabet(r rune, alphabet string) bool {
	r = unicode.ToLower(r)
	for _, template := range alphabet {
		if template == r {
			return true
		}
	}
	return false
}

func isRussian(text string) bool {
	nRussian := 0
	nEnglish := 0
	for _, r := range text {
		if isRuneFromAlphabet(r, engAlphabet) {
			nEnglish += 1
		}
		if isRuneFromAlphabet(r, rusAlphabet) {
			nRussian += 1
		}
	}
	return nRussian > nEnglish
}

func SaveMp3(text string, path string) error {
	args := []string{text, "-o", path}
	if isRussian(text) {
		args = append(args, "-l", "ru")
	}
	return exec.Command("google_speech", args...).Run()
}
