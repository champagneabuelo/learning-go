package main

import "fmt"
import "encoding/json"
import "os"

type Car struct {
	Make  string
	Model string
	Odo   int
	Year  int
	Color string
}

func main() {

	toyota := Car{"Toyota", "Camry", 30000, 2015, "blue"}
	honda := Car{"Honda", "Civic", 45000, 2015, "silver"}
	bmw := Car{"BMW", "2002", 93000, 1969, "red"}
	chevy := Car{"Chevrolet", "Express", 67000, 1969, "black"}
	merc := Car{"Mercedes", "Benz", 32000, 2011, "white"}
	nissan := Car{"Nissan", "Altima", 12000, 2011, "red"}

	years := map[int][]Car{
		2015: []Car{toyota, honda},
		1969: []Car{bmw, chevy},
		2011: []Car{merc, nissan},
	}

	makes := map[string][]Car{
		"Toyota":   []Car{toyota},
		"BMW":      []Car{bmw},
		"Mercedes": []Car{merc},
	}

	years_data, _ := json.MarshalIndent(years, "", " ")
	makes_data, _ := json.MarshalIndent(makes, "", " ")

	fmt.Printf("%s\n", years_data)
	fmt.Printf("%s\n", makes_data)

	f, _ := os.Create("json_output.txt")
	f.Write(years_data)
	f.Write(makes_data)
	f.Close()
}
