package main

import (
	"fmt"
	"../matrix"
)

func main() {

	m := matrix.NewMatrix([][]float32 {

		{1, 0, 0, 0, 2, 0},
		{0, 0, 3, 4, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 8, 0, 5},
		{0, 0, 0, 0, 0, 0},
		{0, 7, 1, 0, 0, 6}} )


	//cl := m.ToCoordinates()
	//csr := m.ToCSR()
	//ellpack := m.ToELLPACK()

	fmt.Println("Матрица: ")
	fmt.Println(m.String())
	
}