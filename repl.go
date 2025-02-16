import (
	"bufio"
	"fmt"
	"io"
	"lexer"
	"token"
	)

	const PROMPT = ">> "

	
	func Start(in io.Reader, out io.Writer){
		scanner := bufio.NewScanner(in)

		for{
			fmt.Printf(PROMPT)
			scanner := scanner.Scan()

			if !scanned {
				return
			}

			line := scanner.Text()
			l := lexer.New(line)

			for tok := l.nextToken(); tok.Type != token.EOF; tok = l.nextToken(){
				fmt.Printf("%+v\n" , tok)
			}
		}


	}