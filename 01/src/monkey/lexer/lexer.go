package lexer

type Lexer struct {
    input           string
    position        int
    readPostition   int
    ch              byte
}

func New(input string) *Lexer {
    l := &Lexer{input: input}
    return l
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPostition += 1
}


