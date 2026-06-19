package handlers

import (
	"encoding/json"
	"qr-api/models"
	"qr-api/services"
	"qr-api/utils"

	"github.com/gofiber/fiber/v3"
)

// QRHandler maneja el request de la matriz, valida, rota y saca el QR
func QRHandler(c fiber.Ctx) error {

	var req models.MatrixRequest

	// mapeamos el body a la estructura
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "request invalido",
		})
	}

	// check para que no manden matrices vacias o deformes
	if err := utils.ValidateMatrix(req.Matrix); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// paso 1: rotamos a la derecha (90 grados)
	rotatedMatrix := utils.RotateMatrix(req.Matrix)

	// paso 2: sacamos las matrices Q y R usando la que ya esta rotada
	result, err := services.CalculateQR(rotatedMatrix)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// le pasamos el paquete a node para que calcule max, min, etc
	nodeResponse, err := services.SendToNode(result)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// desempaquetamos el json que nos devuelve node
	var stats map[string]interface{}
	json.Unmarshal(nodeResponse, &stats)

	// por ultimo devolvemos todo junto al front
	return c.JSON(fiber.Map{
		"rotatedMatrix": rotatedMatrix,
		"statistics":    stats,
	})
}
