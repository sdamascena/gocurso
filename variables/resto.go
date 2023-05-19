package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Salario float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Maria"
	Estado = true
	Salario = 1000.88
	Fecha = time.Now()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Salario)
	fmt.Println(Fecha)

}

func CoviertoaTexto(numero int) (bool, string) {
	var texto string
	texto = strconv.Itoa(numero)
	return true, texto
}
