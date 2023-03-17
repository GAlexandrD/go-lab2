package lab2

import (
	"strconv"
	"strings"
)

// TODO: document this function.
// PostfixToPrefix converts
func PostfixToPrefix(input string) (string, error) {
	inputElems := strings.Split(input, " ")
	stack := make([]string, 0)
	result := ""
	for _, el := range inputElems {
		if isOperator(el) {
			if len(stack) < 2 {
				return "", InvalidExpressionError{}
			}
			operand1 := stack[len(stack)-1]
			operand2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			newOp := el + " " + operand2 + " " + operand1
			stack = append(stack, newOp)
		} else {
			_, err := strconv.ParseFloat(el, 64);
			if err != nil {
				return "", InvalidExpressionError{}
			}
			stack = append(stack, el)
		}
	}

	if len(stack) != 1 {
		return "", InvalidExpressionError{}
	}
	result += stack[len(stack)-1]

	return result, nil
}

func isOperator(e string) bool {
	if e == "+" || e == "-" || e == "*" || e == "/" || e == "^" {
		return true
	}
	return false
}


type InvalidExpressionError struct {}

func (e InvalidExpressionError) Error() string {
	return "Invalid Expression"
}