package main

import "fmt"
import "os"
import "math"
import "strconv"

func vars() (string, int, int) {
	a, b, c := "hello", 3, 7
	return a, b, c
}

func guess1() {
	if num := os.Args[1]; num == "1" {
		fmt.Println("red")
	} else if num == "2" {
		fmt.Println("orange")
	} else if num == "3" {
		fmt.Println("yellow")
	} else if num == "4" {
		fmt.Println("green")
	} else if num == "5" {
		fmt.Println("blue")
	} else {
		fmt.Println("Enter a number between 1-5")
	}
}

func guess2() {
	switch num := os.Args[2]; num {
	case "1":
		fmt.Println("red")
	case "2":
		fmt.Println("orange")
	case "3":
		fmt.Println("yellow")
	case "4":
		fmt.Println("green")
	case "5":
		fmt.Println("blue")
	default:
		fmt.Println("Enter a number between 1-5")
	}
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	max := int(math.Floor(math.Sqrt(float64(x))))
	for i := 2; i < max+1; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func primeList() {
	for i := 0; i < 100; i += 1 {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func longestIncreasingRun(x int) string {
	s := strconv.Itoa(x)
	run, maxRun := 0, 0
	index := 0
	last := '0' - 1
	for i, a := range s {
		if a > last {
			run++
		} else {
			run = 1
		}
		if run > maxRun {
			maxRun = run
			index = i
		}
		last = a
	}
	return s[index-maxRun+1 : index+1]
}

func main() {
	a, b, c := vars()
	fmt.Printf("a is %s b is %d and c is %d\n", a, b, c)

	guess1()
	guess2()

	primeList()

	num := 1232
	fmt.Println("longest run (", num, "): ", longestIncreasingRun(num))
	num = 27648923679
	fmt.Println("longest run (", num, "): ", longestIncreasingRun(num))
}
