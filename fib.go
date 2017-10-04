package main

import "fmt"

func main() {
	fmt.Println(fiblist(15))
}

func fib(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fiblist(n int) []int {
	nums := make([]int, n, n+1)
	for i := 1; i <= n; i++ {
		nums[i-1] = fib(i)
	}
	return nums
}
