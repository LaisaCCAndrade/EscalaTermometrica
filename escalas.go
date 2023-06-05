package main

import "fmt"

const ebulicaoK = 273

func main() {
	tempK := 310
	tempC := tempK - ebulicaoK

	fmt.Printf("A temperatura de ebulição da agua em °C = %v", tempC)
}
