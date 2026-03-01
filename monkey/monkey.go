package monkey

import (
	"fmt"
	"os"

	"github.com/AlanValdevenito/monkey-interpreter/repl"
)

// Monkey represents the Monkey interpreter, managing REPL and script execution.
type Monkey struct {
	repl *repl.REPL
}

// New creates a new Monkey interpreter instance.
func New(input, output *os.File) *Monkey {
	return &Monkey{
		repl: repl.New(input, output),
	}
}

// StartREPL launches the interactive REPL session.
func (m *Monkey) StartREPL() {
	m.repl.Start()
}

// RunScriptFile executes a Monkey script from the given filename.
func (m *Monkey) RunScriptFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	// Print file content with prompt
	fInfo, err := f.Stat()
	if err == nil && fInfo.Size() > 0 {
		content := make([]byte, fInfo.Size())
		if _, err := f.ReadAt(content, 0); err == nil {
			if _, errF := fmt.Fprintf(m.repl.Output(), ">> %s\n", content); errF != nil {
				return errF
			}
		}
		// Reset file pointer for tokenizing
		_, _ = f.Seek(0, 0)
	}

	m.repl.RunScript(f)
	return nil
}
