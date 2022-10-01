package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rmwebdev/jojomic-restapi/app/model"
	"github.com/rmwebdev/jojomic-restapi/app/services"
)

func CheckHargaGet(c *fiber.Ctx) error {
	db := services.DB
	var harga []model.HargaAPI
	db.Model(&model.Harga{}).Find(&harga)

	resp := &fiber.Map{
		"error": false,
		"data":  harga,
	}
	return c.Status(200).JSON(resp)
}
