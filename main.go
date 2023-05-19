package main

import (
	"fmt"

	"github.com/sdamascena/gocurso/variables"
)

func main() {
	Estado, texto := variables.CoviertoaTexto(1588)
	fmt.Println(Estado)
	fmt.Println(texto)
}
