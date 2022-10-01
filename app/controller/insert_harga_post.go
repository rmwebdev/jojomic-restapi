package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rmwebdev/jojomic-restapi/app/lib"
	"github.com/rmwebdev/jojomic-restapi/app/model"
	"github.com/rmwebdev/jojomic-restapi/app/services"
)

func InsertHargaPost(c *fiber.Ctx) error {
	api := new(model.Harga)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB

	if err := db.Create(&api).Error; err != nil {
		newErr := &fiber.Map{"error": true, "message": "Server not ready", "reff_id": api.ID}
		return c.Status(500).JSON(newErr)
	}
	resp := &fiber.Map{"error": false, "reff_id": api.ID}
	return c.Status(200).JSON(resp)
}
