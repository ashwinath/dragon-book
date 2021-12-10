package two

import (
	"fmt"
)

type parser struct {
	value    []rune
	token    rune
	position int
	err      error
}

func (p *parser) match(symbol rune) {
	if p.token != symbol {
		p.err = fmt.Errorf("expected %c, got %c", symbol, p.token)
	}

	// Only advance if there is no error or else we get an index out of range panic
	if p.err == nil {
		p.getToken()
	}
}

func (p *parser) getToken() {
	p.token = p.value[p.position]
	p.position++
}

func (p *parser) s1() {
	switch p.token {
	case 'a':
		p.getToken()
	case '+', '-':
		p.getToken()
		p.s1()
		p.s1()
	default:
		p.err = fmt.Errorf("did not understand token, %c", p.token)
		return
	}
}

// S -> + S S | - S S | a
func (p *parser) PositiveNegativeA() error {
	p.getToken()
	p.s1()
	return p.err
}

func (p *parser) s2() {
	if p.token == '(' {
		p.match('(')
		p.s2()
		p.match(')')
		p.s2()
	} else if p.token != '\n' && p.token != '(' && p.token != ')' {
		p.err = fmt.Errorf("unexpected symbol %c", p.token)
	}
}

// S -> S ( S ) S | ε
func (p *parser) Brackets() error {
	p.getToken()
	p.s2()
	return p.err
}
