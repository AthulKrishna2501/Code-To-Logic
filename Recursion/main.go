package main

import "fmt"

// Factorial
func factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

// Fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Sum of digits
func sumofDigits(n int) int {
	if n == 0 {
		return 0
	}

	return n%10 + sumofDigits(n/10)
}

// Reverse String
func reverseString(n string) string {
	if len(n) == 0 {
		return n
	}

	return reverseString(n[1:]) + string(n[0])
}

func main() {
	result := factorial(5)
	fmt.Println("Factorial of 5 is :", result)
	fibonacci := fibonacci(5)
	fmt.Println("Fibonacci of 5 is :", fibonacci)
	sum := sumofDigits(198)
	fmt.Println("Sum of digits is :", sum)
	reverse := reverseString("hello")
	fmt.Println(reverse)
}
