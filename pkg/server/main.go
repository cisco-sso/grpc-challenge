package main

import "fmt"

func main() {
	// srv := grpc.NewServer()
	input := "One small step for man"
	fmt.Println(Reverse(input))
}

// type reverseServer struct{}

// Reverse takes a string and returns the reverse
//ex: "One small step for man" -> "nam rof pets llams enO"
func Reverse(ctx context.Context s string) string {
	a := []rune(s)
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return string(a)
}
