package handlers

import (
	"encoding/json"
	"qr-api/models"
	"qr-api/services"
	"qr-api/utils"

	"github.com/gofiber/fiber/v3"
)

func QRHandler(c fiber.Ctx) error {

	var req models.MatrixRequest

	// leemos los datos que nos envían
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "request inválido",
		})
	}

	// verificamos que la matriz esté bien armada
	if err := utils.ValidateMatrix(req.Matrix); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// rotamos 90 grados a la derecha
	rotatedMatrix := utils.RotateMatrix(req.Matrix)

	// calculamos Q y R con la matriz ya rotada
	result, err := services.CalculateQR(rotatedMatrix)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// enviamos las matrices a node para sacar las estadísticas
	nodeResponse, err := services.SendToNode(result)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "falló la comunicación con node: " + err.Error(),
		})
	}

	// leemos el JSON que devuelve Node y revisamos que no traiga errores o texto raro
	var stats map[string]interface{}
	if err := json.Unmarshal(nodeResponse, &stats); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "respuesta inválida de node: " + string(nodeResponse),
		})
	}

	return c.JSON(fiber.Map{
		"rotatedMatrix": rotatedMatrix,
		"statistics":    stats,
	})
}
