package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

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
		hash := os.Args[3]
		fmt.Println(string(CatFile(hash)))

	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}
