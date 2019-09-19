package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var baseOperators = map[string]struct {
	priority int
}{
	"!": {4}, // унарный минус
	"*": {3},
	"/": {3},
	"+": {2},
	"-": {2},
	"(": {1},
	")": {1},
}

func main() {
	var expression string
	fmt.Fscan(os.Stdin, &expression)

	trueExpression, err := transform(expression)
	handleError(err)

	postfix, err := convertInfixToPostfix(trueExpression)
	handleError(err)

	result, err := parsePostfix(postfix)
	handleError(err)

	fmt.Println(result)
}
func handleError(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

// функция парсит введеное выражение, определяет его правильность и рахделяет пробелами на токены
func transform(str string) (res string, err error) {
	var number string
	for _, char := range str {
		if char == ' ' {
			continue
		}
		if _, isOp := baseOperators[string(char)]; isOp && char != '!' {
			if number != "" {
				res += number + " " + string(char) + " "
			} else {
				res += string(char) + " "
			}
			number = ""
		} else {
			if unicode.IsLetter(char) {
				return "", errors.New("Wrong expression! Letters are not expected")
			} else if unicode.IsNumber(char) {
				number += string(char)
			} else if char == '.' && number != "" {
				if !strings.Contains(number, ".") {
					number += string(char)
				} else {
					return "", errors.New("Wrong expression! Invalid float number")
				}
			}
		}
	}
	return res, nil
}

// переводим из инфиксной записи(обычного выражения) в постфиксную
func convertInfixToPostfix(infixStr string) (result string, err error) {
	var nextMinusIsUnar bool = true
	var stack []string
	for _, token := range strings.Fields(infixStr) {
		switch token {
		case "(":
			nextMinusIsUnar = true
			stack = append(stack, token) // пушим "(" в стэк
		case ")":
			nextMinusIsUnar = false
			var operator string
			for {
				// Достаём "(" или оператор из стэка
				if len(stack) < 1 {
					return "", errors.New("Неправильное выражение, проверьте скобки")
				}
				operator, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if operator == "(" {
					break
				}
				result += " " + operator // добавляем оператор в результат
			}
		default:
			if o1, isOp := baseOperators[token]; isOp {
				// Текущий токен это оператор
				if token == "-" && nextMinusIsUnar {
					stack = append(stack, "!")
				} else {
					for len(stack) > 0 {
						// берём верхний оператор из стэка
						op := stack[len(stack)-1]

						// порядок важности операторов ( или если скобка )
						if o2, isOp := baseOperators[op]; !isOp || o1.priority > o2.priority {
							break
						}
						// Верхний элемент - оператор который нужно доставать, делаем pop
						stack = stack[:len(stack)-1] // pop
						result += " " + op
					}
					// пушим токен в стэк(это новый оператор)
					stack = append(stack, token)
				}
			} else {
				if result > "" {
					result += " "
				}
				result += token // добавляем токен в результат
				nextMinusIsUnar = false
			}
		}
	}
	// оставшиеся операторы берём из стэка и добавляем в результат
	for len(stack) > 0 {
		result += " " + stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return result, nil
}

func parsePostfix(postfixStr string) (float64, error) {
	var stack []float64

	for _, token := range strings.Fields(postfixStr) {
		switch token {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "!": // Это наш унарный минус который мы записали в convertInfixToPostfix
			stack[len(stack)-1] = -stack[len(stack)-1]
		default:
			f, err := strconv.ParseFloat(token, 64)
			if err != nil {
				fmt.Println("Not a number")
				os.Exit(1)
			}
			stack = append(stack, f)
		}
	}
	if len(stack) == 1 {
		return stack[0], nil
	}
	return 0, errors.New("wrong postfix input")
}
