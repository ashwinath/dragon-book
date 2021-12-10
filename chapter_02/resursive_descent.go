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
