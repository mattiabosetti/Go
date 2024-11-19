package main

import "fmt"

func main() {
	stringa := "Sium 😊"
	fmt.Println(stringa)
	indexed := stringa[0]
	fmt.Printf("%d %T\n", indexed, indexed)

	for i, v := range stringa {
		fmt.Printf("%d %T %v\n", i, v, v)

	}
}
