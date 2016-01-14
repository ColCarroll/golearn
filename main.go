package main

import (
	"fmt"
	"golearn/linear"
)

func main() {
	x, y, alpha := linear.GenerateLinearData(0.1, 5, 1000)
	model := linear.LinearModel{}
	fmt.Printf("A =\n%v\ny=\n%v\n", x, y)
	model.Fit(x, y)
	fmt.Printf("Actual x=\n%v\ncomputed x=\n%v\n", alpha, model.Coefs)
	fmt.Printf("Mean Squared Error: %.3f", model.Error)
}
