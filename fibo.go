package main

func main() {
	Fib(20)
}

func Fib(n int64) int64 {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}
