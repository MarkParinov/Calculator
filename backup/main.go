package main

import (
	"fmt"
)

type operation struct {
	operation string
	priority  int
}

type Stack struct {
	items []string
}

var NUMBERS = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
var NUMBERS_STRING = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
var OPERATION_SIGNS = []operation{
	{"(", 0},
	{")", 1},
	{"/", 2},
	{"*", 2},
	{"+", 3},
	{"-", 3},
}

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

// Methods for stack

func (stack *Stack) Push(value string) {
	stack.items = append(stack.items, value)
}

func (stack *Stack) Pop() {
	if stack.IsEmpty() {
		return
	}
	stack.items = stack.items[:len(stack.items)-1]
}

func (stack *Stack) IsEmpty() bool {
	if len(stack.items) == 0 {
		return true
	} else {
		return false
	}
}

func (stack *Stack) Peek() string {
	return stack.items[len(stack.items)-1]
}

// Rang functions

func getRang(operator string) int {
	if operator == "(" {
		return 0
	} else if operator == "+" || operator == "-" {
		return 1
	} else if operator == "*" || operator == "/" {
		return 2
	} else if operator == ")" {
		return 3
	} else {
		return -1
	}
}

func getMaxRangOfStack(stack Stack) int {
	max_rang := 0
	for n := 0; n < len(stack.items); n++ {
		if getRang(stack.items[n]) > max_rang {
			max_rang = getRang(stack.items[n])
		}
	}
	return max_rang
}

// func isInt(s string) bool {
// 	for _, c := range s {
// 		if !unicode.IsDigit(c) {
// 			return false
// 		}
// 	}
// 	return true
// }

func StringToNotation(expr_string string) []string {

	var output = []string{}
	var stack Stack
	var last_digit_is_num = false

	for i := 0; i < len(expr_string); i++ { // The main parsing cycle

		if ContainsString(NUMBERS_STRING, string(expr_string[i])) { // Parses the number with multiple digits correctly
			if last_digit_is_num {
				output[len(output)-1] += string(expr_string[i])
			} else {
				output = append(output, string(expr_string[i]))
				last_digit_is_num = true
			}

			// -------------------------operator logic----------------------------

		} else if ContainsOperation(OPERATION_SIGNS, string(expr_string[i])) {

			last_digit_is_num = false

			if stack.IsEmpty() || getMaxRangOfStack(stack) < getRang(string(expr_string[i])) {
				stack.Push(string(expr_string[i]))
			} else if string(expr_string[i]) == ")" {
				stack.Pop()
				for {
					if stack.Peek() == "(" {
						break
					}
				}
			} else if getRang(stack.Peek()) >= getRang(string(expr_string[i])) {
				for getRang(stack.Peek()) >= getRang(string(expr_string[i])) {
					output = append(output, stack.Peek())
					stack.Pop()
				}
				if stack.IsEmpty() || getMaxRangOfStack(stack) < getRang(string(expr_string[i])) {
					stack.Push(string(expr_string[i]))
				}
			}
		}
		fmt.Println(string(expr_string[i])+":", output, stack)
	}

	output = append(output, stack.items...)

	return output

}

func Calc(expression string) (float64, error) {
	return 0, nil
}

func main() {
	fmt.Println(StringToNotation("1*2+3"))
}

/*		stack							output

1. 		['*', '(', '*', '+', ')']		[1, 2]
2.		['*']							[1, 2, '*', '+']

*/
