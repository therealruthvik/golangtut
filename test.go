package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// adding comments for changing the commit

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(rand.Intn(10))
	fmt.Println(math.Sqrt(15))
	fmt.Println(math.Pi)
	fmt.Println(add(200, 345))
	a, b := swap("ruthvik", "garlapati")
	fmt.Println(a, b)
	fmt.Println(nakedreturn(10))
	var i int
	var c, python, java bool
	fmt.Println(i, c, python, java)

	var k, j int = 1, 2
	fmt.Println(k, j)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println(sumAndDivide(23, 123))
	n := 13
	if n > 10 {
		fmt.Println("The number should be less than 10")
	} else {
		fmt.Println("The number is greater than 10")
	}
}

func sumAndDivide(x, y int) float64 {
	var temp int
	for i := x; i <= y; i++ {
		temp += i
	}
	a := float64(x)
	b := float64(y)

	return float64(temp) / (a + b)
}

func swap(x, y string) (string, string) {
	return x, y
}

func nakedreturn(sum int) (x, y int) {
	x, y = sum*4, sum*5
	return
}
