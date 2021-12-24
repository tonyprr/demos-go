package main

func Sum(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Factorial(n int) int {
	fact := 1
	for i := 2; i <= n; i++ {
		fact = fact * i
	}
	return fact
}
