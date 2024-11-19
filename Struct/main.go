package main

import "fmt"

type Persona struct {
	nome  string
	eta   uint8
	amici uint8
}

func main() {
	var persona Persona = Persona{"Alessandro", 17, 40}
	fmt.Println(persona.nome, persona.eta, persona.amici)
}
