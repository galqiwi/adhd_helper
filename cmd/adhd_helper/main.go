package main

import (
	"fmt"
	"github.com/galqiwi/adhd_helper/internal/adhd_helper"
	"os"
)

func main() {
	err := Main()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func Main() error {
	return adhd_helper.Main(getConf())
}
