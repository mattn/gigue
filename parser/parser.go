package parser

import (
	"errors"
	"github.com/suzuken/gs/lexer"
	"github.com/suzuken/gs/types"
	"strconv"
)

type Parser struct {
	lex *lexer.Lex
}

// New returns new parser
func New(lex *lexer.Lex) *Parser {
	return &Parser{lex}
}

// Parse scans given tokens and return it into expression.
func (p *Parser) Parse() (exps types.Expression, err error) {
	// start s-expression
	if p.lex.TokenText() == "(" {
		var list []types.Expression
		// recursive scan until ')'
		for {
			p.lex.Scan()
			if p.lex.TokenText() == ")" {
				break
			}
			ex, err := p.Parse()
			if err != nil {
				return nil, err
			}
			list = append(list, ex)
		}
		return list, nil
	} else if p.lex.TokenText() == ")" {
		return exps, errors.New("unexpected ')'")
	} else {
		token := p.lex.TokenText()
		if token == "#" {
			p.lex.Scan()
			if t := p.lex.TokenText(); t == "t" {
				return types.Boolean(true), nil
			} else if t == "f" {
				return types.Boolean(false), nil
			} else {
				return types.Symbol(token + t), nil
			}
		}
		// try conversion to float. if failed, deal with symbol.
		if n, err := strconv.ParseFloat(token, 64); err == nil {
			return types.Number(n), nil
		}
		return types.Symbol(token), nil
	}
}