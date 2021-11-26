package repl

import (
	"bufio"
	"fmt"
	"io"
	"mortar/lexer"
	"mortar/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, "%v", PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		switch line {
		case "exit":
			return
		case "help":
			fmt.Fprintln(out, "Commands:")
		default:
			l := lexer.New(line)
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Fprintf(out, "%+v\n", tok)
			}
		}

	}
}
