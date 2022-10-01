package model

import "github.com/go-openapi/strfmt"

type Transaksi struct {
	Base
	TransaksiAPI
}

type TransaksiAPI struct {
	Gram  float64 `json:"gram,omitempty" gorm:"type:float8;not null"`
	Type  string  `json:"type,omitempty" gorm:"type:varchar(36);not null"`
	Harga string  `json:"harga,omitempty" gorm:"type:varchar(36);not null"`
	Norek string  `json:"norek,omitempty" gorm:"type:varchar(36);not null"`
}

type TransaksiDataAPI struct {
	Norek     string           `json:"norek,omitempty" gorm:"type:varchar(36);not null"`
	StartDate *strfmt.DateTime `json:"start_date,omitempty" gorm:"type:timestamptz" format:"date-time" `
	EndDate   *strfmt.DateTime `json:"end_date,omitempty" gorm:"type:timestamptz" format:"date-time" `
}
type TransaksiData struct {
	Date         *strfmt.DateTime `json:"date"`
	Gram         float64          `json:"gram"`
	Type         string           `json:"type" `
	HargaTopup   float64          `json:"harga_topup"`
	HargaBuyback float64          `json:"harga_buyback"`
	Norek        string           `json:"norek"`
	Saldo        float64          `json:"saldo"`
}
