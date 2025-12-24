package main

import "fmt"

//Uso de estructuras

type Persona struct {
	nombre string
	edad   int
	correo string
}

func (p *Persona) diHola() {
	fmt.Println("Hola, mi nombre es:", p.nombre)

}

func main() {

	p := Persona{"Alex", 28, "alex@gmail.com"}
	p.diHola()

	/*var x int = 10
	fmt.Println(x)
	editar(&x)
	fmt.Println(x)

	//Uso de punteros y métodos
	//var x int = 10
	//var p *int = &x
	//fmt.Println(&x)
	//fmt.Println(p)

	p := Persona{"Alex", 28, "alex@gmail.com"}
	p.edad = 30
	fmt.Println(p)

	p2 := Persona{"Juan", 40, "juan@gmail.com"}
	fmt.Println(p2)

	var p Persona
	p.nombre = "Alex"
	p.edad = 28
	p.correo = "alex@gmail.com"

	fmt.Println(p)
	*/

	//Uso de mapas
	/*colores := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	for clave, valor := range colores {
		fmt.Printf("Clave: %s, Valor %s\n", clave, valor)
	}

	fmt.Println(colores["rojo"])
	colores["negro"] = "#000000"

	fmt.Println(colores)

	//valor, ok := colores["blanco"]
	if valor, ok := colores["verde"]; ok {
		fmt.Println(valor)
	} else {
		fmt.Println("No existe la clave")
	}

	delete(colores, "azul")

	fmt.Println(colores)

	//fmt.Println(valor)*/

	//Uso de rebanadas
	/*rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)

	fmt.Println(copy(rebanada2, rebanada1))

	fmt.Println(rebanada1)
	fmt.Println(rebanada2)

	nombres := make([]string, 5, 10)
	nombres[0] = "Alex"
	fmt.Println(nombres)

	var rebanada [] int
	diasSemana := []string{"Domingo", "Lunes", "Martes",
		"Miércoles", "Jueves", "Viernes", "Sábado"}

	diaRebanada := diasSemana[0:5]
	fmt.Println(diaRebanada)

	diaRebanada = append(diaRebanada, "Viernes", "Sábado", "Otro Día")
	fmt.Println(diaRebanada)

	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)
	fmt.Println(diaRebanada)

	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

	//Uso de vectores
	/*var vector = [...]int{10, 20, 30, 40, 50}
	vector[0] = 100
	vector[1] = 200

	for i := 0; i < len(vector); i++ {
		fmt.Println(vector[i])
	}

	for index, value := range vector {
		fmt.Printf("Índice: %d, Valor: %d\n", index, value)
	}


	//Uso de Matrices
	var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matriz)
	*/

}

//Uso de punteros y métodos
func editar(x *int) {
	*x = 20
}
