package controller

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rmwebdev/jojomic-restapi/app/lib"
	"github.com/rmwebdev/jojomic-restapi/app/model"
	"github.com/rmwebdev/jojomic-restapi/app/services"
)

func CheckSaldoPost(c *fiber.Ctx) error {

	api := new(model.RekeningData)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}
	db := services.DB
	id, _ := uuid.NewRandom()
	var rekening model.Rekening
	result := db.Where("norek = ?", api.Norek).First(&rekening)
	if result.RowsAffected < 1 {
		newErr := &fiber.Map{"error": true, "message": "Server not ready", "reff_id": &id}
		return c.Status(500).JSON(newErr)
	}
	saldo, _ := strconv.ParseFloat(strings.TrimSpace(rekening.Saldo), 64)

	resp := &fiber.Map{
		"error": false,
		"data":  &fiber.Map{"norek": rekening.Norek, "saldo": saldo},
	}
	return c.Status(200).JSON(resp)
}
