package main

// BMI =  weight (kg) / [height (m)]2
import "fmt"

func main() {
	fmt.Println("Hello , hope you are fine .\nThis is the BMI calculator\nPlease enter your weight , but first select the metric")
	fmt.Println("Press 1 for kg\nPress 2 for lbs")
	var weightChoice int64
	fmt.Scanln(&weightChoice)
	if weightChoice == 1 {
		fmt.Println("Enter weight in kg")
	} else {
		fmt.Println("Enter weight in lbs")
	}
	// get weight input
	var weight float64
	fmt.Scanln(&weight)
	// for converting lbs to kg
	if weightChoice == 2 {
		weight = weight * 0.453592
	}
	fmt.Println("Now , Enter your height , but first select the metric \n Press 1 for cm\n Press 2 for inches\n Press 3 for meter")
	var heightChoice int
	fmt.Scanln(&heightChoice)
	if heightChoice == 1 {
		fmt.Println("Enter Height in cm")
	} else if heightChoice == 2 {
		fmt.Println("Enter Height in inches")
	} else {
		fmt.Println("Enter Height in meters")
	}
	var height float64
	fmt.Scanln(&height)
	if heightChoice == 1 {
		height = height / 100
	} else if heightChoice == 2 {
		height = height * 0.0254
	}
	BMI := weight / (height * height)
	var howIsBmi string
	if BMI <= 18.5 {
		howIsBmi = "Unerweight"
	} else if BMI > 18.5 && BMI < 24.9 {
		howIsBmi = "Normal"
	} else if BMI >= 24.9 && BMI <= 29.9 {
		howIsBmi = "Overweight"
	} else {
		howIsBmi = "Obese"
	}
	fmt.Printf("\nYour BMI is %v and you are %v \n HAVE A NICE DAY ! ", BMI, howIsBmi)

}
