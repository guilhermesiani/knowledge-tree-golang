package bhaskara

import "math"

func CalcMaxProfit(
	priceByPassenger float64,
	startDiscount int,
	discount float64,
) int {
	a, b := findCoefficients(priceByPassenger, startDiscount, discount)
	x1, x2 := applyQuadraticEquation(a, b)

	return int((x1 + x2) / 2)
}

func findCoefficients(
	priceByPassenger float64,
	startDiscount int,
	discount float64,
) (float64, float64) {
	a := discount * -1
	b := priceByPassenger + (discount * float64(startDiscount))
	return a, b
}

func applyQuadraticEquation(a, b float64) (float64, float64) {
	c := make(chan float64)
	go quadraticEquationFormulaByOp(a, b, "-", c)
	go quadraticEquationFormulaByOp(a, b, "+", c)
	return <-c, <-c
}

func quadraticEquationFormulaByOp(a, b float64, op string, c chan float64) {
	divisor := 2 * a
	sqrtDelta := math.Sqrt(math.Pow(b, 2))
	if op == "+" {
		c <- (-b + sqrtDelta) / divisor
	}
	c <- (-b - sqrtDelta) / divisor
}