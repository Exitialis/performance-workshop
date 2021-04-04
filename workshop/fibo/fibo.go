package main

func main() {
	Fib(20)
}

// Возьмем fibo как пример того, что производительность ПК для одноядерной работы почти стоит на месте
// можно запустить на своем старом маке код и посмотреть бенчмарк для fibo, и сравнить с этим

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
