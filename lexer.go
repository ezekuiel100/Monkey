package lexer

type Token struct{
    type string
    literal string
}

type Lexer struct {
	input string
	position int       
	readPosition int  // point to the next character in the input
	ch byte
}

func New(input string) *Lexer{
	l := &Lexer{input: input}
    l.readChar()
	return l
}

func (l *Lexer) readChar(){
	if l.readPosition >= len(l.input){
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}


