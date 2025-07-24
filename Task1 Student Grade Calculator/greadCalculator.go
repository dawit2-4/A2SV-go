package main

import (
	"fmt"
)

func average(arr []int) float64 {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr))
}

func main() {
	var name string
	var size int
	fmt.Println("Whats your name?")
	fmt.Scan(&name)
	fmt.Println("How many subjects?")
	fmt.Scan(&size)

	subjects := make([]string, size)
	scores := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Enter the %d subject's name: ", i+1)
		fmt.Scan(&subjects[i])

		var grade int
		for {
			fmt.Printf("Enter the %d subject's score(0-100): ", i+1)
			fmt.Scan(&grade)

			if grade >= 0 && grade <= 100 {
				break
			}
			fmt.Println("Invalid grade! Please enter a value between 0 and 100.")
		}
		scores[i] = grade
	}

	avg := average(scores)
	fmt.Println("\n--- Grade Report ---")
	fmt.Printf("Student Name: %s\n", name)
	for i := 0; i < size; i++ {
		fmt.Printf("%s: %d\n", subjects[i], scores[i])
	}
	fmt.Printf("Average Grade: %.2f\n", avg)
}
