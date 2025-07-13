package task2

import "fmt"

func Bulbasaur(operand string, x, y int) (int, error) {
	switch operand {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		if y == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return x / y, nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", operand)
	}
}

func main() {
	fmt.Println(Bulbasaur("+", 10, 5))
}