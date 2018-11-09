package base

import "fmt"

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupport this op")
	}
}

func main() {
	if resullt, error := eval(3, 4, "、"); error != nil {

		fmt.Println(error)
	} else {
		fmt.Println(resullt)
	}

}
