package main

import "fmt"

func main() {

	const ponto_ebulicaoK = 373
	var ponto_ebulicaoC = (ponto_ebulicaoK - 273)

	fmt.Println("temperatura de ebulição em kelvin ºK: ", ponto_ebulicaoK)
	fmt.Println("temperatura de ebulição em celsio ºC: ", ponto_ebulicaoC)
}
