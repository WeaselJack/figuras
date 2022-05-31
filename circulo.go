package figuras

import "math"

type Circulo struct {
	//Tambien hay que poner los atributos en mayusculas
	// para hacerlos publicos. Igual que con estructuras, funciones e interfaces
	//Mayusculas en todos los sitios. Incluido la funcion de abajo
	Radio float32
}

func (res *Circulo) Area() (string, float32) {
	return "El area del circulo es: ", math.Pi * (res.Radio * res.Radio)
}

func (res *Circulo) Perimetro() (string, float32) {
	return "El perimetro del circulo es: ", 2 * math.Pi * res.Radio
}
