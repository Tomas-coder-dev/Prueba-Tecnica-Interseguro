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
