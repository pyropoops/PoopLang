package repl

import (
	"PoopLang/lexer"
	"PoopLang/token"
	"bufio"
	"fmt"
	"os"
)

const PROMPT = ">> "

func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		text := scanner.Text()
		l := lexer.NewLexer(text)
		tok := l.NextToken()
		for tok.Type != token.EOF {
			fmt.Printf("[%s, %s]\n", tok.Type, tok.Literal)
			tok = l.NextToken()
		}
	}
}