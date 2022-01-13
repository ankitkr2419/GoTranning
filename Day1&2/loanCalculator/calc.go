package main

import (
	loan "calc/calculate"
	"fmt"
)

func main() {
	var principle float32
	var rate float32
	var time int
	fmt.Scanf("%f", &principle)
	fmt.Scanf("%f", &rate)
	fmt.Scanf("%d", &time)

	fmt.Println(loan.Calculate(principle, rate, time))
}
