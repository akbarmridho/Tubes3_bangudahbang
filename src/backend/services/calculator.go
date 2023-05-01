package services

import (
	"errors"
	// "fmt"
	"strconv"
)

func Precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 0
	}
}

func ApplyOp(a float64, b float64, op rune) float64 {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	default:
		return 0
	}
}

func PopVal(stack *[]float64) (float64, error) {
	if len(*stack) == 0 {
		return 0, errors.New("there's a mistake in the math expression entered")
	}
	ret := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return float64(ret), nil
}

func PopOp(stack *[]rune) (rune, error) {
	if len(*stack) == 0 {
		return 0, errors.New("there's a mistake in the math expression entered")
	}
	ret := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return ret, nil
}

func Calculate(input string) (string, error) {
	var (
		values []float64
		ops    []rune
	)

	for i := 0; i < len(input); i++ {
		var c rune = rune(input[i])
		if c == ' ' {
			continue
		} else if c == '(' {
			ops = append(ops, c)
		} else if c >= '0' && c <= '9' {
			var val float64 = 0
			j := i
			for j < len(input) && input[j] >= '0' && input[j] <= '9' {
				val = float64(val*10 + float64(input[j]-'0'))
				j++
			}
			values = append(values, val)
			i = j - 1
		} else if c == ')' {
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				val2, err := PopVal(&values)
				if err != nil {
					return "There's a mistake in the math expression entered", errors.New("invalid expression")
				}
				val1, err := PopVal(&values)

				op, err1 := PopOp(&ops)
				if err != nil || err1 != nil {
					return "There's a mistake in the math expression entered", errors.New("invalid expression")
				}

				values = append(values, ApplyOp(val1, val2, op))
			}
			if len(ops) > 0 {
				PopOp(&ops)
			}
		} else {
			for len(ops) > 0 && Precedence(ops[len(ops)-1]) >= Precedence(c) {
				val2, err := PopVal(&values)

				if err != nil {
					return "There's a mistake in the math expression entered", errors.New("invalid expression")
				}

				val1, err := PopVal(&values)

				op, err1 := PopOp(&ops)

				if err != nil || err1 != nil {
					return "There's a mistake in the math expression entered", errors.New("invalid expression")
				}
				if val2 == 0 && op == '/' {
					return "Undefined", errors.New("invalid expression")
				}
				values = append(values, ApplyOp(val1, val2, op))
			}
			ops = append(ops, c)
		}
	}
	for len(ops) > 0 {
		val2, err := PopVal(&values)
		val1, err := PopVal(&values)
		op, err1 := PopOp(&ops)

		if err != nil || err1 != nil {
			return "There's a mistake in the math expression entered", errors.New("invalid expression")
		}

		values = append(values, ApplyOp(val1, val2, op))
	}
	if len(values) > 1 || len(ops) > 1 {
		return "There's a mistake in the math expression entered", errors.New("invalid expression")
	}
	return strconv.FormatFloat(values[len(values)-1], 'f', -1, 32), nil
}

// func main() {
// 	fmt.Println(Calculate("6 * 2 + (5 - 3) * 3 - 8"))
// 	fmt.Println(6*2 + (5-3)*3 - 8)
// 	fmt.Println(Calculate("(3 + 4) + 7 * 2 - 1 - 9"))
// 	fmt.Println((3 + 4) + 7*2 - 1 - 9)
// 	fmt.Println(Calculate("5 - 2 + 4 * (8 - (5 + 1)) + 9"))
// 	fmt.Println(5 - 2 + 4*(8-(5+1)) + 9)
// 	fmt.Println(Calculate("(8 - 1 + 3) * 6 - ((3 + 7) * 2)"))
// 	fmt.Println((8-1+3)*6 - ((3 + 7) * 2))
// }
