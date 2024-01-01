package adhd_helper

import (
	"github.com/galqiwi/adhd_helper/internal/clipboard"
	"github.com/galqiwi/adhd_helper/internal/play"
	"github.com/galqiwi/adhd_helper/internal/text_to_mp3"
	"os"
	"path/filepath"
)

func Main(cfg *Config) error {
	dirname, err := os.MkdirTemp("", "")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(dirname)
	}()

	outFile := filepath.Join(dirname, "output.mp3")

	text, err := clipboard.GetClipboard(cfg.UseSelection)
	if err != nil {
		return err
	}
	err = text_to_mp3.SaveMp3(text, outFile)
	if err != nil {
		return err
	}

	if cfg.Play {
		err = play.Play(outFile)
		if err != nil {
			return err
		}
	}

	return nil
}
