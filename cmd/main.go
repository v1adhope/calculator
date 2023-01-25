package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/v1adhope/calculator/internal/calculator"
)

func main() {
	calc := calculator.New()

	greenFMT := color.New(color.FgGreen, color.Bold)
	magentaFMT := color.New(color.FgHiMagenta, color.Bold)
	redFMT := color.New(color.FgRed, color.Bold)
	yellowFMT := color.New(color.FgYellow, color.Bold)

	keywordFMT := greenFMT.SprintFunc()

	fmt.Printf("Prompt: Actions (%s - Addition,  %s - Subtraction, %s - Division, %s - Multiplication)\n",
		keywordFMT(calculator.Operations[0]),
		keywordFMT(calculator.Operations[1]),
		keywordFMT(calculator.Operations[2]),
		keywordFMT(calculator.Operations[3]))
	fmt.Println("Prompt: First argument (number a)")
	fmt.Println("Prompt: Second argument (number b)")
	fmt.Println("Prompt: Then print an action and a next number to continue")
	fmt.Printf("Print %s for example, %s for clear the result and %s for exit\n",
		keywordFMT("eg"),
		keywordFMT("ac"),
		keywordFMT("q"))

	isFirstRez := true //If the first calculation, suggest entering both numbers

Loop:
	for {
		magentaFMT.Print("\nPrint action: ")

		var err error
		calc.Action, err = calculator.ReadAction()
		if err != nil {
			redFMT.Println(err)
			continue Loop
		}

		switch calc.Action {
		case "eg":
			eg(yellowFMT)
			continue Loop
		case "ac":
			calc.AllClear()
			isFirstRez = true
			yellowFMT.Println("Result cleared")
			continue Loop
		case "q":
			break Loop
		default:
			if isFirstRez {
				magentaFMT.Print("Print the first number: ")

				calc.Rez, err = calculator.ReadNumber()
				if err != nil {
					redFMT.Println(err)
					continue Loop
				}
			}

			magentaFMT.Print("Print the next number: ")

			calc.Val, err = calculator.ReadNumber()
			if err != nil {
				redFMT.Println(err)
				continue Loop
			}

			err = calc.Calculate()
			if err != nil {
				redFMT.Println(err)
				continue Loop
			}

			greenFMT.Printf("Result: %v\n", calc.Rez)
			if isFirstRez {
				isFirstRez = false
			}
		}
	}
}

func eg(colorFMT *color.Color) {
	colorFMT.Println("\nExample:")
	colorFMT.Println("\tPrint action:")
	colorFMT.Println("\tadd")
	colorFMT.Println("\tPrint the first number")
	colorFMT.Println("\t12")
	colorFMT.Println("\tPrint the next number")
	colorFMT.Println("\t10")
	colorFMT.Println("\tResult: 22")

	colorFMT.Println("\n\tPrint action:")
	colorFMT.Println("\tdiv")
	colorFMT.Println("\tPrint the next number")
	colorFMT.Println("\t2")
	colorFMT.Println("\tResult: 11")
}
