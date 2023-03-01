package calculator

import (
	"errors"
	"math"
)

var (
	ErrWrongSyntax   = errors.New("wrong syntax, see examples")
	ErrUnknownAction = errors.New("unknown action")
)

const (
	OpAddition       = "add"
	OpSubtraction    = "sub"
	OpDivision       = "div"
	OpMultiplication = "mul"
	OpSquareRoot     = "sqrt"
	OpAllClear       = "ac"
)

type calculator struct {
	Action     string
	Res, Val   float64
	IsFirstRes bool
}

func New() *calculator {
	return &calculator{IsFirstRes: true}
}

func (c *calculator) Addition() {
	c.Res += c.Val
}

func (c *calculator) Subtraction() {
	c.Res -= c.Val
}

func (c *calculator) Division() {
	c.Res /= c.Val
}

func (c *calculator) Multiplication() {
	c.Res *= c.Val
}

func (c *calculator) SquareRoot() {
	c.Res = math.Sqrt(c.Res)
}

func (c *calculator) Calculate() error {
	switch c.Action {
	default:
		return ErrUnknownAction
	case OpAddition:
		c.Addition()
	case OpSubtraction:
		c.Subtraction()
	case OpDivision:
		c.Division()
	case OpMultiplication:
		c.Multiplication()
	case OpSquareRoot:
		c.SquareRoot()
	}

	c.IsFirstRes = false
	return nil
}

func (c *calculator) AllClear() {
	c.Res = 0
	c.IsFirstRes = true
}
