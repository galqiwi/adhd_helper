package adhd_helper

import (
	"fmt"
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
	return nil
}
