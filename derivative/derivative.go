package main

import "fmt"

type FX func(x float64) float64

const dx = 0.000000001

func square(x float64) float64 {
	return x * x
}

func cube(x float64) float64 {
	return x * x * x
}

func derivative(fx FX) FX {
	return func(x float64) float64 {
		return (fx(x+dx) - fx(x)) / dx
	}
}

func main() {
	dxSquare := derivative(square)
	fmt.Println("f(x)=x^2")
	fmt.Printf("f(3)=%f\n", square(3))
	fmt.Printf("dxf(3)=%f\n", dxSquare(3))

	dxCube := derivative(cube)
	fmt.Println("g(x)=x^3")
	fmt.Printf("g(2)=%f\n", cube(2))
	fmt.Printf("dxg(2)=%f\n", dxCube(2))
}
