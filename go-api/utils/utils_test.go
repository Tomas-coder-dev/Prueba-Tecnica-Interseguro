package utils

import (
	"reflect"
	"testing"
)

// TestValidateMatrix prueba los casos de éxito y de error al validar la matriz
func TestValidateMatrix(t *testing.T) {
	// Caso 1: Matriz válida
	validMatrix := [][]float64{{1, 2}, {3, 4}}
	if err := ValidateMatrix(validMatrix); err != nil {
		t.Errorf("Esperaba nil (sin error), pero obtuve: %v", err)
	}

	// Caso 2: Matriz vacía
	emptyMatrix := [][]float64{}
	if err := ValidateMatrix(emptyMatrix); err == nil {
		t.Errorf("Esperaba un error por matriz vacía, pero obtuve nil")
	}

	// Caso 3: Matriz deforme (no rectangular)
	unevenMatrix := [][]float64{{1, 2}, {3}}
	if err := ValidateMatrix(unevenMatrix); err == nil {
		t.Errorf("Esperaba un error por matriz no rectangular, pero obtuve nil")
	}
}

// TestRotateMatrix verifica que la matemática de la rotación sea exacta
func TestRotateMatrix(t *testing.T) {
	// Matriz de entrada 3x3
	input := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Lo que matemáticamente esperamos recibir (90° horario)
	expected := [][]float64{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}

	result := RotateMatrix(input)

	// reflect.DeepEqual compara que todos los valores internos de los arrays sean idénticos
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Rotación incorrecta.\nEsperaba: %v\nObtuve: %v", expected, result)
	}
}
