package handlers

import (
	"qr-api/models"
	"qr-api/services"
	"qr-api/utils"

	"github.com/gofiber/fiber/v3"
)

// QRHandler recibe la matriz, valida su formato, la rota, calcula el QR
// y reenvía el resultado a Node.js para obtener las estadísticas.
func QRHandler(c fiber.Ctx) error {

	var req models.MatrixRequest

	// Parseamos el body
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	// Validamos que sea una matriz válida
	if err := utils.ValidateMatrix(req.Matrix); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 1. ROTAMOS LA MATRIZ
	rotatedMatrix := utils.RotateMatrix(req.Matrix)

	// 2. Calculamos Q y R usando la matriz ROTADA
	result, err := services.CalculateQR(rotatedMatrix)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Pasamos el resultado al servicio de Node
	nodeResponse, err := services.SendToNode(result)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Send(nodeResponse)
}
