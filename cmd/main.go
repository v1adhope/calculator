package main

import (
	"fmt"

	"github.com/v1adhope/calculator/internal/calculator"
	"github.com/v1adhope/calculator/internal/colorize"
)

const (
	Example = "eg"
	Exit    = "q"
)

func main() {
	calc := calculator.New()
	colorFMT := colorize.New()

	GreenKeyword := colorFMT.Green.SprintFunc()

	fmt.Printf("Prompt: Actions (%s - Addition,  %s - Subtraction, %s - Division, %s - Multiplication)\n",
		GreenKeyword(calculator.Addition),
		GreenKeyword(calculator.Subtraction),
		GreenKeyword(calculator.Division),
		GreenKeyword(calculator.Multiplication))
	fmt.Println("Prompt: First argument (number a)")
	fmt.Println("Prompt: Second argument (number b)")
	fmt.Println("Prompt: Then print an action and a next number to continue")
	fmt.Printf("Print %s for example, %s for clear the result and %s for exit\n",
		GreenKeyword(Example),
		GreenKeyword(calculator.AllClear),
		GreenKeyword(Exit))

Loop:
	for {
		colorFMT.Magenta.Print("\nPrint action: ")

		var err error
		calc.Action, err = calculator.ReadAction()
		if err != nil {
			colorFMT.Red.Println(err)
			continue Loop
		}

		switch calc.Action {
		case Example:
			printExample(colorFMT)
			continue Loop
		case calculator.AllClear:
			calc.AllClear()
			colorFMT.Yellow.Println("Result cleared")
			continue Loop
		case Exit:
			break Loop
		default:
			if calc.IsFirstResult {
				colorFMT.Magenta.Print("Print the first number: ")

				calc.Rez, err = calculator.ReadNumber()
				if err != nil {
					colorFMT.Red.Println(err)
					continue Loop
				}
			}

			colorFMT.Magenta.Print("Print the next number: ")

			calc.Val, err = calculator.ReadNumber()
			if err != nil {
				colorFMT.Red.Println(err)
				continue Loop
			}

			err = calc.Calculate()
			if err != nil {
				colorFMT.Red.Println(err)
				continue Loop
			}

			colorFMT.Green.Printf("Result: %v\n", calc.Rez)
		}
	}
}

func printExample(colorFMT *colorize.ColorFMT) {
	colorFMT.Yellow.Println("\nExample:")
	colorFMT.Yellow.Println("\tPrint action:")
	colorFMT.Yellow.Println("\tadd")
	colorFMT.Yellow.Println("\tPrint the first number")
	colorFMT.Yellow.Println("\t12")
	colorFMT.Yellow.Println("\tPrint the next number")
	colorFMT.Yellow.Println("\t10")
	colorFMT.Yellow.Println("\tResult: 22")

	colorFMT.Yellow.Println("\n\tPrint action:")
	colorFMT.Yellow.Println("\tdiv")
	colorFMT.Yellow.Println("\tPrint the next number")
	colorFMT.Yellow.Println("\t2")
	colorFMT.Yellow.Println("\tResult: 11")
}
