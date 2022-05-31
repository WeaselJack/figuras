package figuras

import "fmt"

type Geometria interface {
	Area() (string, float32)
	Perimetro() (string, float32)
}

func Medidas(nombre string, resultado Geometria) {

	fmt.Println(resultado.Area())
	fmt.Println(resultado.Perimetro())
}
