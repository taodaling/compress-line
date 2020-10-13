package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func decompress() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		s = strings.ReplaceAll(s, `\n`, "\n")
		fmt.Println(s)
	}
}

func compress() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Print(s)
		fmt.Print(`\n`)
	}
}

func main() {
	inv := false
	for _, arg := range os.Args[1:] {
		if arg == "-i" {
			inv = true
		}
	}
	if inv {
		decompress()
	} else {
		compress()
	}
}
