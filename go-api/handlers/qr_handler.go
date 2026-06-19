package handlers

import (
	"qr-api/models"
	"qr-api/services"
	"qr-api/utils"

	"github.com/gofiber/fiber/v3"
)

// QRHandler recibe la matriz, valida su formato, calcula el QR
// y reenvía el resultado a Node.js para obtener las estadísticas.
func QRHandler(c fiber.Ctx) error {

	var req models.MatrixRequest

	// Parseamos el body
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	// Validamos que sea una matriz válida (rectangular, no vacía)
	if err := utils.ValidateMatrix(req.Matrix); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Calculamos Q y R
	result, err := services.CalculateQR(req.Matrix)
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
