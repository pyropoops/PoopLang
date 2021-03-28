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

		var tokens []token.Token
		for {
			tok := l.NextToken()
			if tok.Type == token.EOF {
				break
			}
			tokens = append(tokens, tok)
		}

		if len(l.GetErrors()) > 0 {
			for _, err := range l.GetErrors() {
				fmt.Println(err)
			}
		} else {
			for _, tok := range tokens {
				fmt.Printf("[%s, %s]\n", tok.Type, tok.Literal)
			}
		}
	}
}
