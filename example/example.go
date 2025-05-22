package main

import (
	"time"

	"github.com/ttskarlsson/shellterm"
)

func main() {
	term := shellterm.NewTerm()
	defer term.Close()
	for {
		term.Writeln("Something")
		term.Flush()
		time.Sleep(500 * time.Millisecond)
		term.Writeln("Something more")
		term.Flush()
		time.Sleep(500 * time.Millisecond)
	}
}

