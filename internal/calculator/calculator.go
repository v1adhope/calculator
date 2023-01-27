package calculator

import (
	"fmt"
)

const (
	// Errors
	WrongSyntax   = "wrong syntax, see examples"
	UnknownAction = "unknown action"

	// Operations
	Addition       = "add"
	Subtraction    = "sub"
	Division       = "div"
	Multiplication = "mul"
	AllClear       = "ac"
)

type calculator struct {
	Action        string
	Rez, Val      float64
	IsFirstResult bool
}

func New() *calculator {
	//If the first calculation, suggest entering both numbers
	return &calculator{IsFirstResult: true}
}

func (c *calculator) Addition() {
	c.Rez += c.Val
}

func (c *calculator) Subtraction() {
	c.Rez -= c.Val
}

func (c *calculator) Division() {
	c.Rez /= c.Val
}

func (c *calculator) Multiplication() {
	c.Rez *= c.Val
}

func (c *calculator) Calculate() error {
	switch c.Action {
	default:
		return fmt.Errorf(UnknownAction)
	case Addition:
		c.Addition()
	case Subtraction:
		c.Subtraction()
	case Division:
		c.Division()
	case Multiplication:
		c.Multiplication()
	}

	c.IsFirstResult = false
	return nil
}

func (c *calculator) AllClear() {
	c.Rez = 0
	c.IsFirstResult = true
}
