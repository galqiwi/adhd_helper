package main

import (
	"flag"
	"github.com/galqiwi/adhd_helper/internal/adhd_helper"
)

func getConf() *adhd_helper.Config {
	useSelectionFlag := flag.Bool("selection", false, "Use selection instead of clipboard")
	outputPathFlag := flag.String("output", "", "Output file path")
	playFlag := flag.Bool("play", false, "Play audio")

	flag.Parse()

	config := &adhd_helper.Config{
		UseSelection: *useSelectionFlag,
		OutputPath:   *outputPathFlag,
		Play:         *playFlag,
	}

	return config
}
