package cmd

import "fmt"

func printArgs(args []string) {
	if len(args) == 0 {
		return
	}

	for _, arg := range args {
		fmt.Printf("%v\n", arg)
	}
}
