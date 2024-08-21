package repl

import (
	"Interpreter/lexer"
	"Interpreter/token"
	"bufio"
	"fmt"

	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		
		line := scanner.Text()
		
		l := lexer.NewLexer(line)
		
		for tok := l.GetNextToken(); tok.Type != token.EOF; tok = l.GetNextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
