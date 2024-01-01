package clipboard

import (
	"os/exec"
	"strings"
)

func GetClipboard() (string, error) {
	cmd := exec.Command("xclip", "-selection", "clipboard", "-o", "/dev/stdout")
	stdoutBuf := strings.Builder{}
	cmd.Stdout = &stdoutBuf
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return stdoutBuf.String(), nil
}
