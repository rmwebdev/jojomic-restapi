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

func TopupPost(c *fiber.Ctx) error {
	api := new(model.TopupAPI)
	if err := lib.BodyParser(c, api); nil != err {
		return lib.ErrorBadRequest(c, err)
	}

	db := services.DB

	var data model.Topup
	lib.Merge(api, &data)

	var harga model.Harga
	result := db.Last(&harga)
	if result.RowsAffected < 1 {
		return lib.ErrorNotFound(c, "Price not available")
	}
	hargaInput, _ := strconv.ParseFloat(strings.TrimSpace(data.Harga), 64)
	saldoInput, _ := strconv.ParseFloat(strings.TrimSpace(data.Gram), 64)

	if harga.HargaTopup != hargaInput {
		newErr := &fiber.Map{"error": true, "message": "Harga yang dimasukan harus sama dengan harga topup", "reff_id": data.ID}
		return c.Status(422).JSON(newErr)
	}
	tx := db.Begin()

	if err := tx.Create(&data).Error; err != nil {
		newErr := &fiber.Map{"error": true, "message": "Server not ready", "reff_id": data.ID}
		tx.Rollback()
		return c.Status(500).JSON(newErr)
	}

	rek := model.Rekening{}
	res := db.Where("norek = ?", data.Norek).Find(&rek)
	oldSaldo, _ := strconv.ParseFloat(strings.TrimSpace(rek.Saldo), 64)
	saldo := saldoInput + oldSaldo
	newSaldo := fmt.Sprintf("%v", saldo)

	if res.RowsAffected > 0 {
		if err := tx.Model(&rek).Where("norek = ?", data.Norek).Update("saldo", newSaldo).Error; err != nil {
			tx.Rollback()
			return err
		}
	} else {
		rekening := model.Rekening{}
		rekening.Norek = data.Norek
		rekening.Saldo = data.Gram
		if err := db.Create(&rekening).Error; err != nil {
			db.Rollback()
			return err
		}
	}
	transaksi := model.Transaksi{}
	transaksi.Gram = saldoInput
	transaksi.Harga = data.Harga
	transaksi.Norek = data.Norek
	transaksi.Type = "topup"
		if err := db.Create(&transaksi).Error; err != nil {
			db.Rollback()
			return err
		}

	resp := &fiber.Map{"error": false, "reff_id": data.ID}
	tx.Commit()
	return c.Status(200).JSON(resp)
}
