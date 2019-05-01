package main

import (
	"fmt"
	"os"

	"github.com/cyruzin/trailer/trailer/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
