package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
	Operation sign priorities:
		* , / 	- 	1
		+, - 	- 	0
		) 		-	2
		( 		- 	3

	Rules for parsing:
		- An opened bracket is always added to the stack
		- A closed bracket that closes the opened one moves all the operations in them to the output, destroying themselves afterwards

*/

// Secondary methods for easier development

// Methods for stack

// Rang functions

func getRang(operator string) int {
	if operator == "(" {
		return 1
	} else if operator == "+" || operator == "-" {
		return 2
	} else if operator == "*" || operator == "/" {
		return 3
	} else if operator == ")" {
		return 4
	} else {
		return -1
	}
}

// func isInt(s string) bool {
// 	for _, c := range s {
// 		if !unicode.IsDigit(c) {
// 			return false
// 		}
// 	}
// 	return true
// }

// Stack methods and structure

type operation struct {
	operation string
	priority  int
}

type stack struct {
	items []string
}

type f_stack struct {
	items []float32
}

// String stack functions

func (stack *stack) Push(value string) {
	stack.items = append([]string{value}, stack.items...)
}

func (stack *stack) Pop() {
	if stack.IsEmpty() {
		return
	}
	stack.items = stack.items[1:len(stack.items)]
}

func (stack *stack) IsEmpty() bool {
	if len(stack.items) == 0 {
		return true
	} else {
		return false
	}
}

func (stack *stack) Peek() string {
	if stack.IsEmpty() {
		return ""
	} else {
		return stack.items[0]
	}
}

// Float stack functions

func (stack *f_stack) Push(value float32) {
	stack.items = append([]float32{value}, stack.items...)
}

func (stack *f_stack) Pop() {
	if stack.IsEmpty() {
		return
	}
	stack.items = stack.items[1:len(stack.items)]
}

func (stack *f_stack) IsEmpty() bool {
	if len(stack.items) == 0 {
		return true
	} else {
		return false
	}
}

func (stack *f_stack) Peek() float32 {
	if stack.IsEmpty() {
		return 0
	} else {
		return stack.items[0]
	}
}

func ContainsOperation(operation_array []operation, value string) bool {
	for n := 0; n < len(operation_array); n++ {
		if operation_array[n].operation == value {
			return true
		}
	}
	return false
}

func ContainsInt(int_array []int, value int) bool {
	for n := 0; n < len(int_array); n++ {
		if int_array[n] == value {
			return true
		}
	}
	return false
}

func ContainsString(string_array []string, value string) bool {
	for n := 0; n < len(string_array); n++ {
		if string_array[n] == value {
			return true
		}
	}
	return false
}

func FindElementInStringSlice(string_array []string, value string) int {
	for n := 0; n < len(string_array); n++ {
		if string_array[n] == value {
			return n
		}
	}
	return 0
}

func PopExpression(expression []string) []string {
	if len(expression) == 0 {
		return expression
	} else {
		return expression[1:]
	}
}

var NUMBERS = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
var NUMBERS_STRING = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var OPERATION_SIGNS = []operation{
	{"/", 1},
	{")", 2},
	{"*", 1},
	{"+", 0},
	{"-", 0},
}

func StringToNotation(expr_string string) []string { // Converts the given string to prefix polish notation expression

	var output = []string{}
	var stack stack
	var last_digit_is_num = false
	var cur_el string

	for i := 0; i < len(expr_string); i++ { // The main parsing cycle

		cur_el = string(expr_string[i]) // Assign the current symbol to a variable

		if ContainsString(NUMBERS_STRING, cur_el) { // Parses numbers with multiple digits
			if last_digit_is_num {
				output[len(output)-1] += cur_el
				last_digit_is_num = true
			} else {
				output = append(output, cur_el)
				last_digit_is_num = true
			}

			// Operator Logic

		} else {
			last_digit_is_num = false
			if stack.IsEmpty() || cur_el == "(" {
				stack.Push(cur_el)
			} else if cur_el == ")" {
				for stack.Peek() != "(" {
					output = append(output, stack.Peek())
					stack.Pop()
				}
				stack.Pop()
			} else if getRang(stack.Peek()) < getRang(cur_el) {
				stack.Push(cur_el)
			} else if getRang(stack.Peek()) >= getRang(cur_el) {
				for getRang(stack.Peek()) >= getRang(cur_el) {
					output = append(output, stack.Peek())
					stack.Pop()
				}
				if stack.IsEmpty() || cur_el == "(" {
					stack.Push(cur_el)
				}
			}
		}
	}

	output = append(output, stack.items...)

	return output
}

func Calc(expression []string) (float64, error) {
	calc_stack := f_stack{}
	for len(expression) != 0 {
		for i := 0; i < len(expression); i++ {
			fmt.Println("[ITER", i, " START] Expression: ", expression, "; stack: ", calc_stack)
			switch expression[0] {
			case "+":
				num1 := calc_stack.Peek()
				calc_stack.Pop()
				num2 := calc_stack.Peek()
				calc_stack.Pop()
				calc_stack.Push(num1 + num2)
				expression = PopExpression(expression)

			case "-":
				num1 := calc_stack.Peek()
				calc_stack.Pop()
				num2 := calc_stack.Peek()
				calc_stack.Pop()
				calc_stack.Push(num2 - num1)
				expression = PopExpression(expression)

			case "*":
				num1 := calc_stack.Peek()
				calc_stack.Pop()
				num2 := calc_stack.Peek()
				calc_stack.Pop()
				calc_stack.Push(num2 * num1)
				expression = PopExpression(expression)

			case "/":
				num1 := calc_stack.Peek()
				calc_stack.Pop()
				num2 := calc_stack.Peek()
				calc_stack.Pop()
				if num1 == 0 {
					return 0, errors.New("devision by zero is not allowed")
				} else {
					calc_stack.Push(num2 / num1)
				}
				expression = PopExpression(expression)

			default:
				num_to_push, err := strconv.ParseFloat(expression[0], 32)
				if err != nil {
					panic(err)
				}
				calc_stack.Push(float32(num_to_push))
				expression = PopExpression(expression)
			}
		}
	}
	out := calc_stack.Peek()

	return float64(out), nil

}

func main() {
	fmt.Println(StringToNotation("10-2*(5/2)+3"))
	fmt.Println(Calc(StringToNotation("")))
}
