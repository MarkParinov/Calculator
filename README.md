# Golang calculator

The 'main.go' contains functions and data structures used for math expressions evaluation. The functions
are capable of parsing any mathematical expression to postfix polish notation and then calculating the result.

# Usage 

* The 'StringToNotation' function takes a math expression as an argument and returns the same expression in postfix polish notation.
* The 'Calc' function takes a math expression, written in **!postfix!** polish notation and returns the result of the expression and an
error, in case one of those occurs.

# Specification:

This method of calculating returns a float32 value. It doesn't allow devision by zero and it also works with brackets, allowing more complex
expressions to be calculated.

# Please note,

that this model only works with numbers, and does nt have the ability to operate with letters or variable names. If you want to see this feature, you can complete this project with your own implementation, or just star this repo to let me know that you want to see it in the module.