package models

type MatrixRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

type MatrixResponse struct {
	Q [][]float64 `json:"q"`
	R [][]float64 `json:"r"`
}

type StatisticsResponse struct {
	Max               float64 `json:"max"`
	Min               float64 `json:"min"`
	Average           float64 `json:"average"`
	Sum               float64 `json:"sum"`
	HasDiagonalMatrix bool    `json:"hasDiagonalMatrix"`
}
