// NOTE: 0 = nil
package calculator

import "fmt"

const WrongSyntax = "wrong syntax, see examples"

var Operations = [...]string{"add", "sub", "div", "mul"}

type calculator struct {
	Action   string
	Rez, Val float64
}

func New() *calculator {
	return &calculator{}
}

func Eg() {
	fmt.Println("\nExample:")
	fmt.Println("\tPrint action:")
	fmt.Println("\tadd")
	fmt.Println("\tPrint the first number")
	fmt.Println("\t12")
	fmt.Println("\tPrint the next number")
	fmt.Println("\t10")
	fmt.Println("\tResult: 22")

	fmt.Println("\n\tPrint action:")
	fmt.Println("\tdiv")
	fmt.Println("\tPrint the next number")
	fmt.Println("\t2")
	fmt.Println("\tResult: 11")
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
		return fmt.Errorf("there is no such action")
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
