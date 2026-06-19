package services

import (
	"qr-api/models"

	"gonum.org/v1/gonum/mat"
)

// CalculateQR usa gonum para aplicar la factorización QR sobre la matriz de entrada.
func CalculateQR(matrix [][]float64) (*models.MatrixResponse, error) {

	rows := len(matrix)
	cols := len(matrix[0])

	// Gonum requiere un slice 1D, así que aplanamos la matriz nativa
	data := make([]float64, 0, rows*cols)
	for _, row := range matrix {
		data = append(data, row...)
	}

	a := mat.NewDense(rows, cols, data)

	var qr mat.QR
	qr.Factorize(a)

	var q mat.Dense
	var r mat.Dense

	qr.QTo(&q)
	qr.RTo(&r)

	return &models.MatrixResponse{
		Q: denseToSlice(&q),
		R: denseToSlice(&r),
	}, nil
}

// denseToSlice convierte el tipo *mat.Dense de gonum de vuelta a [][]float64
func denseToSlice(m *mat.Dense) [][]float64 {

	rows, cols := m.Dims()
	result := make([][]float64, rows)

	for i := 0; i < rows; i++ {
		result[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			result[i][j] = m.At(i, j)
		}
	}

	return result
}
