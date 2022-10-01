package model

type Harga struct {
	Base
	AdminID string `json:"admin_id,omitempty" gorm:"type:varchar(36);not null"`
	HargaAPI
}

type HargaAPI struct {
	HargaTopup   float64 `json:"harga_topup,omitempty" gorm:"type:float8;not null"`
	HargaBuyback float64 `json:"harga_buyback,omitempty" gorm:"type:float8;not null"`
}
