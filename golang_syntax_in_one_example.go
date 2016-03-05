package main

// This was written by clover_rui

import (
	"fmt"
	"math"
)

var x int = 1
var java = "java"

const Pi = 3.14

const (
	big   = 1 << 100
	small = big >> 99
)

func main() {
	const key = "world"
	fmt.Println(Pi, "Hello", key)
	fmt.Println("Hello,GO!")
	fmt.Println("Happy", math.Pi, "Birthday!")
	fmt.Println(math.Nextafter(3, 2))
	fmt.Println(math.Pi)
	fmt.Println("求和：", add(2, 3))
	fmt.Println(mutil(5, 5))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(45))

	g := "go"
	c := 1
	fmt.Println(x, java, g, c)

	fmt.Println()

	fmt.Println(needInt(small))
	//溢出
	//fmt.Println(needInt(big))
	fmt.Println(needFloat(small))
	fmt.Println(needFloat(big))

	count := 0
	for i := 0; i < 100; i++ {
		count += i
	}
	fmt.Println(count)

	sum := 1
	// for sum < 100 {
	// sum += sum
	// }
	//go语言中没有while语句，java别的语言中的while语句在go语言中都是通过for语句实现
	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)
	//死循环
	for {
	}
	for {
	}
}

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 { return x * 0.1 }

func split(sum int) (x int) {
	var y int
	x = sum*3/2 + y
	return
}

func swap(a, b string) (string, string) {
	return a, b
}

func add(x int, y int) int {
	return x + y
}

func mutil(x, y int) int {
	return x * y
}
