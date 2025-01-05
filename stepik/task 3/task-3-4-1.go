package main

import "fmt"

func main() {
	value1, value2, operation := readTask()
	x, ok := value1.(float64)
	if !ok {
		fmt.Printf("value=%v: %T\n", value1, value1)
		return
	}

	y, ok := value2.(float64)
	if !ok {
		fmt.Printf("value=%v: %T\n", value2, value2)
		return
	}

	op, ok := operation.(string)
	if !ok {
		fmt.Println("неизвестная операция")
		return
	}

	switch op {
	case "+":
		result := x + y
		fmt.Printf("%.4f", result)
	case "-":
		result := x - y
		fmt.Printf("%.4f", result)
	case "*":
		result := x * y
		fmt.Printf("%.4f", result)
	case "/":
		result := x / y
		fmt.Printf("%.4f", result)

	default:
		fmt.Println("неизвестная операция")
	}

}
