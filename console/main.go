package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("周长计算")
	fmt.Print("请输入半径：")
	var radius float64
	fmt.Scanf("%f", &radius)

	area := math.Pi * math.Pow(radius, 2)
	fmt.Printf("半径为%.2f的园周长 = %.2f \n", radius, area)
}
