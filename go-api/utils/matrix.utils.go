package utils

import "errors"

// ValidateMatrix asegura que la matriz tenga datos y sea rectangular (nxm).
func ValidateMatrix(matrix [][]float64) error {
	if len(matrix) == 0 {
		return errors.New("la matriz no puede estar vacía")
	}

	cols := len(matrix[0])
	if cols == 0 {
		return errors.New("las filas de la matriz no pueden estar vacías")
	}

	for _, row := range matrix {
		if len(row) != cols {
			return errors.New("la matriz debe ser estrictamente rectangular")
		}
	}

	return nil
}

// RotateMatrix rota una matriz rectangular (n x m)
func RotateMatrix(matrix [][]float64) [][]float64 {
	if len(matrix) == 0 {
		return matrix
	}

	rows := len(matrix)
	cols := len(matrix[0])

	rotated := make([][]float64, cols)
	for i := range rotated {
		rotated[i] = make([]float64, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = matrix[i][j]
		}
	}

	return rotated
}
