package main

import (
	"bufio"
	"fmt"
	"os"
)

type checkMake func(string) bool

func BikePrinter(fn checkMake, bikes ...string) {
	for _, bike := range bikes {
		if fn(bike) {
			fmt.Printf(bike + " ")
		}
	}
	fmt.Println()
}

func Italian(moto string) bool {
	italian := []string{"Ducati", "Aprilia", "Moto", "Guzzi"}
	for _, make := range italian {
		if moto == make {
			return true
		}
	}
	return false
}

func British(moto string) bool {
	british := []string{"BSA", "Triumph", "Norton"}
	for _, make := range british {
		if moto == make {
			return true
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var moto string
	motoList := []string{}

	for moto != "q" {
		fmt.Print("Input Motorcycle (q to quit): ")
		scanner.Scan()
		moto = scanner.Text()
		if moto != "q" {
			fmt.Println("Motorcycle entered: ", moto)
			motoList = append(motoList, moto)
		}
	}
	fmt.Printf("Italian: ")
	BikePrinter(Italian, motoList...)
	fmt.Printf("British: ")
	BikePrinter(British, motoList...)
	fmt.Printf("German: ")
	BikePrinter(func(moto string) bool {
		german := []string{"BMW", "MZ", "Sommer"}
		for _, make := range german {
			if moto == make {
				return true
			}
		}
		return false
	}, motoList...)
}
