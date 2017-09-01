package main

import (
        "fmt"
        "os"
)

func Vars(){
    fmt.Println(os.Environ())
}

func main() {
	Vars();
}
