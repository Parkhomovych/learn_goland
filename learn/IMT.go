package learn

import (
	"fmt"
	"math"
)

func MainCalculateIMT() {
	// var vladIMT1 = CalculateIMT1(100, 1.8)
	// fmt.Println(vladIMT1)

	// vladIMT2 := CalculateIMT2(100, 1.8)
	// fmt.Println(vladIMT2)

	// fmt.Println(CalculateIMT3(100, 1.8))

	// var vladHeight float64
	// var vladWeight float64

	// fmt.Print("Enter your height in meters (example: 1.8):")
	// fmt.Scan(&vladHeight)

	// fmt.Print("Enter your weight in kilograms (example: 80):")
	// fmt.Scan(&vladWeight)

	// var vladIMT = CalculateIMT(vladWeight, vladHeight)
	// fmt.Println(vladIMT)

	// or

	// var height, weight float64
	// fmt.Println("Enter your height and weight for IMT calculation")

	// fmt.Print("Enter your height in centimeters (example: 180):")
	// fmt.Scan(&height)

	// fmt.Print("Enter your weight in kilograms (example: 80):")
	// fmt.Scan(&weight)

	// IMT := CalculateIMT(weight, height/100)
	// fmt.Print("Your IMT is: ", IMT)

	// or

	// height := GetHeight()
	// weight := GetWeight()
	// IMT := CalculateIMT(weight, height)
	// outputIMT(IMT)

	// or

	height, weight := getUserInput()
	IMT := CalculateIMT(height, weight)
	outputIMT(IMT)
}

func getUserInput() (float64, float64) {
	var height, weight float64
	fmt.Println("Enter your height and weight for IMT calculation")

	fmt.Print("Enter your height in centimeters (example: 180):")
	fmt.Scan(&height)

	fmt.Print("Enter your weight in kilograms (example: 80):")
	fmt.Scan(&weight)

	return height, weight
}

func outputIMT(imt float64) {
	fmt.Printf("Your IMT is: %.2f", imt)
}
func GetHeight() float64 {
	var height float64
	fmt.Print("Enter your height in centimeters (example: 180):")
	fmt.Scan(&height)
	return height
}
func GetWeight() float64 {
	var weight float64
	fmt.Print("Enter your weight in kilograms (example: 80):")
	fmt.Scan(&weight)
	return weight
}
func CalculateIMT(height float64, weight float64) float64 {
	// var imt = weight / math.Pow(height/100, 2)
	// return imt

	// 	imt := weight /  math.Pow(height, 2)
	// return imt

	return weight / math.Pow(height/100, 2)
}
