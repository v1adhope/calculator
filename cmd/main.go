package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/v1adhope/calculator/internal/calculator"
	"github.com/v1adhope/calculator/internal/colorize"
)

const (
	_opExample = "eg"
	_opExit    = "q"
)

func main() {
	calc := calculator.New()
	colorFMT := colorize.New()

	printManual(colorFMT)

	for {
		colorFMT.Magenta.Print("\nPrint action: ")

		var err error
		calc.Action, err = calculator.ReadAction()
		if err != nil {
			colorFMT.Red.Println(err)
			continue
		}

		switch calc.Action {
		case _opExit:
			os.Exit(0)
		case _opExample:
			printExample(colorFMT)
			continue
		case calculator.OpAllClear:
			calc.AllClear()
			colorFMT.Yellow.Println("Result cleared")
			continue
		case calculator.OpSquareRoot:
			var choice string

			if calc.Res != 0 {
				colorFMT.Yellow.Print("From the result?(Y/n): ")

				choice, err = calculator.ReadAction()
				if err != nil {
					colorFMT.Red.Println(err)
					continue
				}
			}

			if strings.EqualFold(choice, "n") || calc.IsFirstRes {
				colorFMT.Magenta.Print("Print the X number: ")

				calc.Res, err = calculator.ReadNumber()
				if err != nil {
					colorFMT.Red.Println(err)
					continue
				}
			}

			err = calc.Calculate()
			if err != nil {
				colorFMT.Red.Println(err)
				continue
			}

			colorFMT.Green.Printf("Result: %.4v\n", calc.Res)
		default:
			if calc.IsFirstRes {
				colorFMT.Magenta.Print("Print the first number: ")

				calc.Res, err = calculator.ReadNumber()
				if err != nil {
					colorFMT.Red.Println(err)
					continue
				}
			}

			colorFMT.Magenta.Print("Print the next number: ")

			calc.Val, err = calculator.ReadNumber()
			if err != nil {
				colorFMT.Red.Println(err)
				continue
			}

			err = calc.Calculate()
			if err != nil {
				colorFMT.Red.Println(err)
				continue
			}

			colorFMT.Green.Printf("Result: %v\n", calc.Res)
		}
	}
}

func printManual(colorFMT *colorize.ColorFMT) {
	GreenKeyword := colorFMT.Green.SprintFunc()

	fmt.Printf("Prompt: Base actions (%s - Addition,  %s - Subtraction, %s - Division, %s - Multiplication)\n",
		GreenKeyword(calculator.OpAddition),
		GreenKeyword(calculator.OpSubtraction),
		GreenKeyword(calculator.OpDivision),
		GreenKeyword(calculator.OpMultiplication))
	fmt.Printf("Prompt: Extra actions (%s - Square root of result or X)\n",
		GreenKeyword(calculator.OpSquareRoot))
	fmt.Println("Prompt: First argument (number a)")
	fmt.Println("Prompt: Second argument (number b)")
	fmt.Println("Prompt: Then print an action and a next number to continue")
	fmt.Printf("Print %s for example, %s for clear the result and %s for exit\n",
		GreenKeyword(_opExample),
		GreenKeyword(calculator.OpAllClear),
		GreenKeyword(_opExit))
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
