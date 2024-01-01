package play

import "os/exec"

func Play(path string) error {
	return exec.Command("vlc", path).Run()
}
