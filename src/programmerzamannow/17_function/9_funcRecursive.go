package main

import "fmt"

func factorialLoop(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func factorialLoopRecursive(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorialLoopRecursive(n-1)
	}

}

func main() {
	multiplyLoop := factorialLoop
	fmt.Println(multiplyLoop(5))
	multiplyLoopRecursive := factorialLoopRecursive
	fmt.Print(multiplyLoopRecursive(5))
}
