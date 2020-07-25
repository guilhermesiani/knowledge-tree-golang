package bhaskara

import "math"

func CalcMaxProfit(
	priceByPassenger float64,
	startDiscount int,
	discount float64,
) int {
	// Find coefficients
	a := discount * -1
	b := priceByPassenger + (discount * float64(startDiscount))

	// Apply Bhaskara method
	x1, x2 := applyBhaskara(a, b)

	return int((x1 + x2) / 2)
}

func applyBhaskara(a, b float64) (float64, float64) {
	c := make(chan float64)
	go bhaskaraFormuleByOp(a, b, "-", c)
	go bhaskaraFormuleByOp(a, b, "+", c)
	return <-c, <-c
}

func bhaskaraFormuleByOp(a, b float64, op string, c chan float64) {
	divisor := 2 * a
	sqrtDelta := math.Sqrt(math.Pow(b, 2))
	if op == "+" {
		c <- (-b + sqrtDelta) / divisor
	}
	c <- (-b - sqrtDelta) / divisor
}