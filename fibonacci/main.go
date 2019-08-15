package main

func fibonacci(n int32) int32 {
	if n == 0 || n == 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	println(fibonacci(40))
}
