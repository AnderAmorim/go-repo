package areago

import (
	"math"
)

const Pi = 3.1415

func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

func Rect(base, altura float64) float64 {
	return base * altura
}

func triangulo(base, altura float64) float64 {
	return base * altura / 2
}
