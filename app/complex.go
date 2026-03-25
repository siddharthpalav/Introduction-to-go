package main

type Complex struct {
	real      float64
	imaginary float64
}

func NewComplex(real float64, imaginary float64) *Complex {
	c := Complex{
		real:      real,
		imaginary: imaginary,
	}
	return &c
}
