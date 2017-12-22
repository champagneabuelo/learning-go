// fibonacci.go: receives user-submitted input as a string, and prints the fibonacci sequence up until the given number
package main

import "fmt"

func fibonacci(n int) {
	p := 0
	q := 1

	for p+q <= n {
		if p < q {
			p += q
			fmt.Printf("%d", p)
		} else {
			q += p
			fmt.Printf("%d", q)
		}

		if p+q <= n {
			fmt.Printf(", ")
		} else {
			fmt.Println()
		}
	}
}


func main() {
	var n int
	fmt.Print("Enter number to calculate: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if n < 1 {
		fmt.Println("Given number is too low, no sequence can be generated.")
	} else {
		fibonacci(n)
	}
}
