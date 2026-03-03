package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/AlanValdevenito/monkey-interpreter/scanner"
	"github.com/AlanValdevenito/monkey-interpreter/token"
)

const PROMPT = ">> "

// REPL encapsulates the Read-Eval-Print Loop logic for the Monkey interpreter.
type REPL struct {
	input  io.Reader
	output io.Writer
}

// New creates a new REPL instance with the given input and output.
func New(input io.Reader, output io.Writer) *REPL {
	return &REPL{
		input:  input,
		output: output,
	}
}

// Start begins the interactive REPL session.
func (r *REPL) Start() {
	scannerInput := bufio.NewScanner(r.input)
	for {
		if _, err := fmt.Fprint(r.output, PROMPT); err != nil {
			return
		}
		scanned := scannerInput.Scan()
		if !scanned {
			return
		}

		line := scannerInput.Text()
		s := scanner.New(line)

		for t := s.NextToken(); t.Type != token.EOF; t = s.NextToken() {
			if _, err := fmt.Fprintf(r.output, "%+v\n", t); err != nil {
				return
			}
		}
	}
}

// RunScript executes a Monkey script from a file, printing tokens to output.
func (r *REPL) RunScript(script io.Reader) {
	buf := bufio.NewReader(script)
	for {
		line, err := buf.ReadString('\n')
		if err != nil && len(line) == 0 {
			break
		}
		s := scanner.New(line)
		for t := s.NextToken(); t.Type != token.EOF; t = s.NextToken() {
			if _, errF := fmt.Fprintf(r.output, "%+v\n", t); errF != nil {
				return
			}
		}
		if err != nil {
			break
		}
	}
}

// Output returns the output writer used by the REPL.
func (r *REPL) Output() io.Writer {
	return r.output
}
