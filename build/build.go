package build

import (
	"fmt"
	"io/ioutil"
	"mortar/evaluator"
	"mortar/lexer"
	"mortar/object"
	"mortar/parser"
)

func Run(filepath string) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	data := string(file)
	l := lexer.New(data)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Printf("%s\n", evaluated.Inspect())
	}
}
