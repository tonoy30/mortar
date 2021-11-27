package repl

import (
	"bufio"
	"fmt"
	"io"
	"mortar/lexer"
	"mortar/parser"
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
			p := parser.New(l)
			program := p.ParseProgram()
			if len(p.Errors()) != 0 {
				printParserErrors(out, p.Errors())
				continue
			}
			io.WriteString(out, program.String())
			io.WriteString(out, "\n")
		}
	}
}
func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
