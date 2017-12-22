package main

import "fmt"
import "os"
import "math"
import "errors"

func writeFile() {

	f, err := os.Create("foo.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Fprintln(f, "data")
}

func quadratic(a float64, b float64, c float64) (float64, float64, error) {
	if a == 0 {
		if b != 0 {
			return c / b * -1, 0, errors.New("Division by zero")
		}
		return c, 0, errors.New("Division by zero")
	}
	if b*b-4*a*c < 0 {
		return 0, 0, errors.New("Taking the square root of a negative number")
	}
	result1 := (-1.0*b + math.Sqrt(b*b-4.0*a*c)) / (2.0 * a)
	result2 := (-1.0*b - math.Sqrt(b*b-4.0*a*c)) / (2.0 * a)
	return result1, result2, nil
}

func main() {
	writeFile()
	r1, r2, e := quadratic(1, -50, 2)
	if e != nil {
		fmt.Println("Error: ", e)
	} else {
		fmt.Println("x vals: ", r1, r2)
	}
	r1, r2, e = quadratic(0, 5, 5)
	if e != nil {
		fmt.Println("Error: ", e)
	} else {
		fmt.Println("x vals: ", r1, r2)
	}
}
