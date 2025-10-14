package matrix

import (
	"strings"
	"fmt"
)

type Matrix struct {
	Data [][]float32
	Rows int
	Cols int
}

type CoordianteList struct {
	Values []float32
	Row    []int
	Col    []int
}

type CSR struct {
	Values    []float32
	Col       []int
	Row_index []int
}

type ELLPACK struct {
	Value [][]float32
	Index [][]int
}

func (m *Matrix) String() string {

	var sb strings.Builder

	for _, row := range m.Data {
		for _, item := range row {

			sb.WriteString(fmt.Sprintf("%4.4f", item))
		}
		
		sb.WriteString("\n")
	}

	return sb.String()
}

func NewMatrix(data [][]float32) *Matrix {

	return &Matrix {Data: data, Rows: len(data), Cols: len(data[0])}
}


func (m *Matrix) ToCoordinates() CoordianteList {

	cl := CoordianteList{}

	for i, row := range m.Data {
		for k, item := range row {

			if item != 0 {

				cl.Values = append(cl.Values, item)
				cl.Row = append(cl.Row, i)
				cl.Col = append(cl.Col, k)

			}
		}
	}

	return cl
}

func (m *Matrix) ToCSR() CSR {

	csr := CSR{}

	for _, row := range m.Data {

		csr.Row_index = append(csr.Row_index, len(csr.Values))

		for k, item := range row {

			if item != 0 {
				csr.Values = append(csr.Values, item)
				csr.Col = append(csr.Col, k)
			}
		}
	}

	return csr
}

func (m *Matrix) ToELLPACK() ELLPACK {

	ellpack := ELLPACK{}
	max_row_len := 0

	for _, row := range m.Data {

		row_len := 0

		for _, item := range row {

			if item != 0 {
				row_len++
			}
		}

		if row_len > max_row_len {
			max_row_len = row_len
		}
	}

	ellpack.Value = make([][]float32, len(m.Data))
	ellpack.Index = make([][]int, len(m.Data))

	for i := 0; i < len(m.Data); i++ {

		ellpack.Value[i] = make([]float32, max_row_len)
		ellpack.Index[i] = make([]int, max_row_len)
	}

	for i, row := range m.Data {

		new_index := 0
		for k, item := range row {

			if item > 0 {

				ellpack.Value[i][new_index] = item 
				ellpack.Index[i][new_index] = k
				new_index++
			}
		}
	}

	return ellpack
}
