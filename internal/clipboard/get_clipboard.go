package clipboard

import (
	"os/exec"
	"strings"
)

func GetClipboard(useSelection bool) (string, error) {
	selection := "clipboard"
	if useSelection {
		selection = "primary"
	}
	cmd := exec.Command("xclip", "-selection", selection, "-o", "/dev/stdout")
	stdoutBuf := strings.Builder{}
	cmd.Stdout = &stdoutBuf
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return stdoutBuf.String(), nil
}
