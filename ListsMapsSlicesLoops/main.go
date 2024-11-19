package main

import (
	"fmt"
)

func main() {
	intArray := [3]int32{1, 2, 3}
	fmt.Println(intArray[0])
	fmt.Println(intArray[1:3])

	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 4)
	fmt.Println(intSlice)
	fmt.Printf("La lunghezza dello slice è %v, mentre la sua capcaità è %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{5, 6, 7}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// Maps
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Alessandro": 23, "Sara": 17}
	fmt.Println(myMap2)
	fmt.Printf("Alessandro ha %v anni\n", myMap2["Alessandro"])

	delete(myMap2, "Alessandro")
	var eta, ok = myMap2["Adams"]
	if ok {
		fmt.Printf("Età: %v\n", eta)
	} else {
		fmt.Println("Nome non valido")
	}

	for nome, eta := range myMap2 {
		fmt.Printf("Nome: %v\n", nome)
		fmt.Printf("Età: %v\n", eta)
	}
	for indice, value := range intSlice {
		fmt.Printf("Indice: %v, Value: %v\n", indice, value)
	}

	for i := 0; i < 10; i++ {
		println(i)
	}
}
