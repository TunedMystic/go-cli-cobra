package main

import (
	"fmt"
	"os"

	"github.com/tunedmystic/go-cli-cobra/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
