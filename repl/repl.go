package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/AlanValdevenito/monkey-interpreter/token"
	"github.com/AlanValdevenito/monkey-interpreter/scanner"
)

const PROMPT = ">> "

// Start initiates the REPL (Read-Eval-Print Loop) for the Monkey interpreter.
func Start(in io.Reader, out io.Writer) {
	inputScanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := inputScanner.Scan()
		if !scanned {
			return
		}

		line := inputScanner.Text()
		s := scanner.New(line)

		for t := s.NextToken(); t.Type != token.EOF; t = s.NextToken() {
			fmt.Fprintf(out, "%+v\n", t)
		}
	}
}