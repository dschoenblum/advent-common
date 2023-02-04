package common

import "fmt"

type Parser struct {
	source string
	offset int
}

func NewParser(source string) Parser {
	return Parser{
		source: source,
		offset: 0,
	}
}

func (p *Parser) AtEnd() bool {
	return p.offset == len(p.source)
}

func (p *Parser) Peek() byte {
	return p.source[p.offset]
}

func (p *Parser) Read() byte {
	r := p.Peek()
	p.Consume()
	return r
}

func (p *Parser) Consume() {
	p.offset++
}

func (p *Parser) Expect(expected byte) {
	r := p.Peek()
	if r != expected {
		panic(fmt.Sprintf("Expected %c, got %c", expected, r))
	}
	p.Consume()
}
