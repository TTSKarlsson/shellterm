package shellterm

import (
	"fmt"

	"golang.org/x/term"
)

func GetSize(fd uintptr) (int, int) {
	if !term.IsTerminal(int(fd)) {
		return 0, 0
	}
	w, h, err := term.GetSize(int(fd))
	if err != nil {
		fmt.Println(err)
		return 0, 0
	}
	return w, h
}
