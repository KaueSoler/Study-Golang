package main

import (
	"fmt"
	"math"
)

func main() {
	var i, j float64

	fmt.Println("Informe dois número por favor:")
	fmt.Scan(&i, &j)
	fmt.Println("Resulta:", math.Abs(i*j))
}
