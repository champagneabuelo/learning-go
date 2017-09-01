package main

import "fmt"

func reverse(brands []string) {
    for i := 0; i < len(brands)/2; i ++ {
    	j := len(brands) - i - 1
    	brands[i], brands[j] = brands[j], brands[i]
    }	
}

func main() {
        brands := []string{"Honda","Yamaha","Ducati","BMW","Triumph","Polaris"}
        fmt.Println(brands)
        reverse(brands)
        fmt.Println(brands)
}