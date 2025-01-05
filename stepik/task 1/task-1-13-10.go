package main

import "fmt"

func main() {
	var n, sum, sum2 int
	fmt.Scan(&n)

	for n > 0 {
		sum += n % 10
		n /= 10
	}
	if sum < 10 {
		fmt.Println(sum)
	}
	if sum > 10 {
		for sum > 0 {
			sum2 += sum % 10
			sum /= 10
		}
		fmt.Println(sum2)
	}

}
