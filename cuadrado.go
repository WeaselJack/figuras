package figuras

type Cuadrado struct {
	Lado float32
}

func (res *Cuadrado) Area() (string, float32) {

	return "El area del cuadrado es: ", res.Lado * res.Lado
}

func (res *Cuadrado) Perimetro() (string, float32) {
	return "El perimetro del cuadrado es: ", 4 * res.Lado
}
