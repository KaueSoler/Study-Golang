package main

import (
	"fmt"
	"math"
	"strconv"

	"project.go/utility"
)

func main() {
	var i, j float64
	fmt.Println("Informe dois número por favor:")
	fmt.Scan(&i)
	fmt.Scan(&j)
	fmt.Println("Resulta:", math.Abs(i*j))
	fmt.Println(utility.Reverse(strconv.FormatFloat(math.Abs(i*j), 'f', -2, 64)))
}
