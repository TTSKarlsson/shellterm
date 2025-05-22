package shellterm

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/ttskarlsson/shellprint/ansi"
	"github.com/ttskarlsson/shellprint/constants"
)

func (t *Term) addLine(line string) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.lines = append(t.lines, line)
}

type Term struct {
	stdin   *os.File
	stdout  *os.File
	stderr  *os.File
	lines   []string
	buf     *bytes.Buffer
	mutex   *sync.Mutex
	width   int
	height  int
	padding rune
}

func NewTerm() *Term {
	term := &Term{
		stdin:  os.Stdin,
		stdout: os.Stdout,
		stderr: os.Stderr,
		lines:  make([]string, 0),
		buf:    &bytes.Buffer{},
		mutex:  &sync.Mutex{},
		width:  0,
		height: 0,
		padding: ' ',
	}
	term.UpdateSize()
	term.stdout.WriteString(constants.AlternateScreen)
	term.stdout.WriteString(constants.ClearScreen)
	term.stdout.WriteString(constants.Home)
	term.stdout.WriteString(constants.HideCursor)
	term.stdout.WriteString(constants.SaveCursor)
	return term
}

func (t *Term) SetStdin(stdin *os.File) *Term {
	t.stdin = stdin
	return t
}

func (t *Term) SetStdout(stdout *os.File) *Term {
	t.stdout = stdout
	return t
}

func (t *Term) SetStderr(stderr *os.File) *Term {
	t.stderr = stderr
	return t
}

func (t *Term) SetPadding(padding rune) *Term {
	t.padding = padding
	return t
}

func (t *Term) Close() {
	fmt.Print(constants.ShowCursor)
	fmt.Print(constants.ResetAlternateScreen)
}

func (t *Term) Fill(r rune) *Term {
	t.stdout.WriteString(ansi.Home())
	for range t.height {
		t.stdout.WriteString(RepeatRune(r, t.width))
	}
	return t
}

func (t *Term) UpdateSize() {
	width, height := GetSize(t.stdout.Fd())
	t.width = width
	t.height = height
}

func (t *Term) String() string {
	sb := strings.Builder{}
	sb.WriteString("Term {\n")
	sb.WriteString(fmt.Sprintf("  Width: %d\n", t.width))
	sb.WriteString(fmt.Sprintf("  Height: %d\n", t.height))
	sb.WriteString("}\n")
	return sb.String()
}

func (t *Term) Writeln(ansiline string) *Term {
	cropped := ansi.PadEndOrCropAnsiString(ansiline, t.padding, t.width)
	t.addLine(cropped)
	return t
}

func (t *Term) Flush() {
	fmt.Print(constants.RestoreCursor)
	for _, line := range t.lines {
		fmt.Printf("%s%s", constants.LineStart, line)
	}
	t.lines = t.lines[:0]
}
