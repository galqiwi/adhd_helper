package text_to_mp3

import "os/exec"

func SaveMp3(text string, path string) error {
	return exec.Command("google_speech", text, "-o", path).Run()
}
