package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	//Si no especificas el nombre de la variable, se crearía automáticamente una con el mismo nombre del tipo
	//En este caso, contactInfo. Equivalente a poner "contactInfo contactInfo"
	contact contactInfo
}

func main() {
	// Método 1 (Demasiado acomplamiento, dependiente del orden de los parámetros)
	// canita := person{"Cañita", "Brava"}
	// Método 2, un poco verboso, y ocupa mucho espacio
	// var canita person
	// canita.firstName = "Cañita"
	// canita.lastName = "Brava"
	// Método 3, el mejor IMHO:
	canita := person{
		firstName: "Cañita",
		lastName: "Brava",
		contact: contactInfo{
			email: "canita@brava.com",
			zipCode: 28080,
		},
	}
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v", p) //Este método pinta también el nombre de los argumentos.
}
