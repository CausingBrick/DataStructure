package main

func main() {
	println(fac(4))
}

func fac(n int64) int64 {
	if n == 1 {
		return n
	}
	return n * fac(n-1)
}
