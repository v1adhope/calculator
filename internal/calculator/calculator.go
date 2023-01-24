package calculator

import (
	"fmt"
)

const WrongSyntax = "wrong syntax, see examples"

// Consistency is important, only for expansion
var Operations = [...]string{"add", "sub", "div", "mul"}

type calculator struct {
	Action   string
	Rez, Val float64
}

func New() *calculator {
	return &calculator{}
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
		return fmt.Errorf("unknown action")
	case Operations[0]:
		c.Addition()
	case Operations[1]:
		c.Subtraction()
	case Operations[2]:
		c.Division()
	case Operations[3]:
		c.Multiplication()
	}

	return nil
}

func (c *calculator) AllClear() {
	c.Rez = 0
}
