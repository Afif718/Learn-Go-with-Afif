package main

import "fmt"

const pi = 3.1416

func mmain() {

	var radius float64

	//lets find Circumference of circle 2pir
	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&radius);


	circleResult := Circumference(radius)
	fmt.Printf(`Circumference of the Circle is: %0.1f`, circleResult)
}

func Circumference(radius float64) float64{
	circum := 2 * pi * radius

	return circum
}