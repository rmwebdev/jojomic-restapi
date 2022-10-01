package controller

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rmwebdev/jojomic-restapi/app/lib"
	"github.com/rmwebdev/jojomic-restapi/app/model"
	"github.com/rmwebdev/jojomic-restapi/app/services"
)

func BuybackPost(c *fiber.Ctx) error {
	api := new(model.TransaksiAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB

	var data model.Transaksi
	lib.Merge(api, &data)

	var topUp model.Topup
	result := db.Where("norek = ?", data.Norek).Find(&topUp)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c, "Saldo not available")
	}
	oldSaldo, _ := strconv.ParseFloat(strings.TrimSpace(topUp.Gram), 64)

	if oldSaldo < data.Gram {
		newErr := &fiber.Map{"error": true, "message": "saldo buyback terlalu rendah", "reff_id": data.ID}
		return c.Status(422).JSON(newErr)
	}
	tx := db.Begin()
	data.Type = "buyback"

	if err := tx.Create(&data).Error; err != nil {
		newErr := &fiber.Map{"error": true, "message": "Server not ready", "reff_id": data.ID}
		tx.Rollback()
		return c.Status(500).JSON(newErr)
	}
	rek := model.Rekening{}
	res := db.Where("norek = ?", data.Norek).Find(&rek)
	oldSaldos, _ := strconv.ParseFloat(strings.TrimSpace(rek.Saldo), 64)
	saldo := oldSaldos - data.Gram
	newSaldo := fmt.Sprintf("%v", saldo)

	if res.RowsAffected > 0 {
		if err := tx.Model(&rek).Where("norek = ?", data.Norek).Update("saldo", newSaldo).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	resp := &fiber.Map{"error": false, "reff_id": data.ID}
	tx.Commit()
	return c.Status(200).JSON(resp)
}
