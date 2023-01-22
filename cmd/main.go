package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/v1adhope/calculator/internal/calculator"
)

func main() {
	calc := calculator.New()
	greenFMT := color.New(color.FgGreen, color.Bold)
	keyword := greenFMT.SprintFunc()
	magentaFMT := color.New(color.FgHiMagenta, color.Bold)
	redFMT := color.New(color.FgRed, color.Bold)

	fmt.Printf("Prompt: Actions (%s - Addition,  %s - Subtraction, %s - Division, %s - Multiplication)\n",
		keyword("add"), keyword("sub"), keyword("div"), keyword("mul"))
	fmt.Println("Prompt: First argument (number a)")
	fmt.Println("Prompt: Second argument (number b)")
	fmt.Println("Prompt: Then print an action and a next number to continue")
	fmt.Printf("Print %s for example, %s for clear the result and %s for exit\n", keyword("eg"), keyword("ac"), keyword("q"))

Loop:
	for {
		magentaFMT.Print("\nPrint action: ")

		str, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			redFMT.Println(calculator.WrongSyntax)
			continue Loop
		}

		calc.Action = strings.TrimSuffix(str, "\n")

		switch calc.Action {
		default:
			if calc.Rez == 0 {
				magentaFMT.Print("Print the first number: ")

				str, err := bufio.NewReader(os.Stdin).ReadString('\n')
				if err != nil {
					redFMT.Println(calculator.WrongSyntax)
					continue Loop
				}

				str = strings.TrimSuffix(str, "\n")
				calc.Rez, err = strconv.ParseFloat(str, 64)
				if err != nil {
					redFMT.Println(calculator.WrongSyntax)
					continue Loop
				}
			}

			magentaFMT.Print("Print the next number: ")

			str, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				redFMT.Println(calculator.WrongSyntax)
				continue Loop
			}

			str = strings.TrimSuffix(str, "\n")
			calc.Val, err = strconv.ParseFloat(str, 64)
			if err != nil {
				redFMT.Println(calculator.WrongSyntax)
				continue Loop
			}

			err = calc.Calculate()
			if err != nil {
				redFMT.Println(err)
				continue Loop
			}

			greenFMT.Printf("Result: %v\n", calc.Rez)
		case "eg":
			calculator.Eg()
			continue Loop
		case "ac":
			calc.AllClear()
			redFMT.Println("Result cleared")
			continue Loop
		case "q":
			break Loop
		}
	}
}
