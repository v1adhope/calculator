package calculator

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadAction() (string, error) {
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", ErrWrongSyntax
	}

	str = strings.TrimSuffix(str, "\n")

	return str, nil
}

func ReadNumber() (float64, error) {
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return 0, ErrWrongSyntax
	}

	str = strings.TrimSuffix(str, "\n")
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, ErrWrongSyntax
	}

	return value, nil
}
