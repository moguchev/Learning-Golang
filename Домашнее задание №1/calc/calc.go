package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	unarMinus string = "!" // унарный минус
	mult      string = "*"
	div       string = "/"
	plus      string = "+"
	minus     string = "-"
	leftBr    string = "("
	rightBr   string = ")"
)

// Priority -- приоритет операций
type Priority int

var baseOperators = map[string]Priority{
	"!": 4,
	"*": 3,
	"/": 3,
	"+": 2,
	"-": 2,
	"(": 1,
	")": 1,
}

func main() {
	var expression string
	fmt.Fscan(os.Stdin, &expression)
	result, err := calculate(expression)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func calculate(expression string) (float64, error) {
	tokens, err := splitTokens(expression)
	if err != nil {
		return 0, err
	}

	postfix, err := convertInfixToPostfix(tokens)
	if err != nil {
		return 0, err
	}

	result, err := parsePostfix(postfix)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// функция парсит введеное выражение, определяет его правильность и разделяет на токены
func splitTokens(str string) (res []string, err error) {
	var number string
	res = make([]string, 0)
	for _, char := range str {
		if char == ' ' {
			continue
		}
		if _, isOp := baseOperators[string(char)]; isOp && char != '!' {
			if number != "" {
				res = append(res, number, string(char))
			} else {
				res = append(res, string(char))
			}
			number = ""
		} else {
			if unicode.IsLetter(char) {
				return nil, errors.New("Wrong expression! Letters are not expected")
			} else if unicode.IsNumber(char) {
				number += string(char)
			} else if char == '.' && number != "" {
				if !strings.Contains(number, ".") {
					number += string(char)
				} else {
					return nil, errors.New("Wrong expression! Invalid float number")
				}
			}
		}
	}
	if number != "" {
		res = append(res, number)
	}
	return res, nil
}

func transferFromStack(stack *[]string, result *[]string) error {
	var operator string
	for {
		// Достаём "(" или оператор из стэка
		if len(*stack) < 1 {
			return errors.New("Неправильное выражение, проверьте скобки")
		}
		operator, *stack = (*stack)[len(*stack)-1], (*stack)[:len(*stack)-1]
		if operator == leftBr {
			break
		}
		*result = append(*result, operator) // добавляем оператор в результат
	}
	return nil
}

func processShuntingYard(token string, nextMinusIsUnar *bool, stack *[]string, result *[]string) {
	if o1, isOp := baseOperators[token]; isOp { // Текущий токен это оператор
		if token == minus && *nextMinusIsUnar {
			*stack = append(*stack, unarMinus)
		} else {
			for len(*stack) > 0 {
				// берём верхний оператор из стэка
				op := (*stack)[len(*stack)-1]
				// порядок важности операторов ( или если скобка )
				if o2, isOp := baseOperators[op]; !isOp || o1 > o2 {
					break
				}
				// Верхний элемент - оператор который нужно доставать, делаем pop
				*stack = (*stack)[:len(*stack)-1] // pop
				*result = append(*result, op)
			}
			// пушим токен в стэк(это новый оператор)
			*stack = append(*stack, token)
		}
	} else { // текущй токен - операнд
		*result = append(*result, token) // добавляем токен в результат
		*nextMinusIsUnar = false
	}
}

// переводим из инфиксной записи(обычного выражения) в постфиксную
func convertInfixToPostfix(infix []string) (result []string, err error) {
	var nextMinusIsUnar bool = true
	var stack []string
	result = make([]string, 0)
	for _, token := range infix {
		switch token {
		case leftBr:
			nextMinusIsUnar = true
			stack = append(stack, token) // пушим "(" в стэк
		case rightBr:
			nextMinusIsUnar = false
			err := transferFromStack(&stack, &result)
			if err != nil {
				return nil, err
			}
		default:
			processShuntingYard(token, &nextMinusIsUnar, &stack, &result)
		}
	}
	// оставшиеся операторы берём из стэка и добавляем в результат
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return result, nil
}

func parsePostfix(postfix []string) (float64, error) {
	var stack []float64

	for _, token := range postfix {
		switch token {
		case plus:
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case minus:
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case mult:
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case div:
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case unarMinus: // Это наш унарный минус который мы записали в convertInfixToPostfix
			stack[len(stack)-1] = -stack[len(stack)-1]
		default:
			f, err := strconv.ParseFloat(token, 64)
			if err != nil {
				log.Fatal(err)
			}
			stack = append(stack, f)
		}
	}
	if len(stack) == 1 {
		return stack[0], nil
	}
	return 0, errors.New("wrong postfix input")
}
