package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rmwebdev/jojomic-restapi/app/lib"
	"github.com/rmwebdev/jojomic-restapi/app/model"
	"github.com/rmwebdev/jojomic-restapi/app/services"
)

func CheckMutasiPost(c *fiber.Ctx) error {

	api := new(model.TransaksiDataAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}
	db := services.DB
	id, _ := uuid.NewRandom()
	var transaksi []model.TransaksiData
	res := db.Model(&model.Transaksi{}).
		Select(`
			"transaksi".id,
			"transaksi".created_at AS date,
			"transaksi".gram,
			"transaksi".harga AS harga_topup,
			"transaksi".harga AS harga_buyback,
			"transaksi".type,
			"rekening".norek,
			"rekening".saldo
	`).
		Joins(`LEFT JOIN (
		SELECT id, saldo, norek FROM 
		"rekening" r
	)rekening ON "transaksi".norek = rekening.norek`).
		Where(`"transaksi".norek = ?`, api.Norek).Find(&transaksi)

	if res.RowsAffected == 0 {
		newErr := &fiber.Map{"error": true, "message": "Server not ready", "reff_id": &id}
		return c.Status(500).JSON(newErr)
	}
	resp := &fiber.Map{
		"error": false,
		"data":  &transaksi,
	}

	return c.Status(200).JSON(resp)
}
