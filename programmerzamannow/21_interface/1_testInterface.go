package main

import (
	"bfcTest/packages/helper"
	"fmt"
	"math"
)

type shape interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	r float64
}
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (p persegi) keliling() float64 {
	return 4 * p.sisi
}

func (l lingkaran) luas() float64 {
	return math.Pi * l.r * l.r
}

func (l lingkaran) keliling() float64 {
	return 2 * math.Pi * l.r
}

func printShapeInfo(s shape) {
	fmt.Printf("Keliling %T = %.2f \n", s, s.keliling())
	fmt.Printf("Luas %T = %.2f \n", s, s.luas())
}

func main() {
	l := lingkaran{r: 10}
	printShapeInfo(l)
	helper.PrintBreak("")

	shapes := []shape{
		persegi{sisi: 12},
		lingkaran{r: 10},
		persegi{sisi: 10},
		lingkaran{r: 100},
	}

	fmt.Println(shapes)

	for _, val := range shapes {
		printShapeInfo(val)
		helper.PrintBreak("")
	}

}
