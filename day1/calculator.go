package day1

import "math"

func addition(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func subtraction(num1 float64, num2 float64) float64 {
	return num1 - num2
}

func multiplication(num1 float64, num2 float64) float64 {
	return num1 * num2
}

func division(num1 float64, num2 float64) float64 {
	return num1 / num2
}

func sin(num float64) float64 {
	return math.Sin(num)
}

func cos(num float64) float64 {
	return math.Cos(num)
}

func tan(num float64) float64 {
	return math.Tan(num)
}

func sqrt(num float64) float64 {
	return math.Sqrt(num)
}
