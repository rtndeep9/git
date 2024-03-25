package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
		os.Exit(1)
	}

	switch command := os.Args[1]; command {
	case "init":
		GitInit()
	case "cat-file":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "usage: mygit cat-file <sha>\n")
			os.Exit(1)
		}
		hash := CatFile(os.Args[3])
		fmt.Printf("%s", hash)

	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}
