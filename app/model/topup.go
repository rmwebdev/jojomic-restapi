package model

type Topup struct {
	Base
	TopupAPI
}

type TopupAPI struct {
	Gram  string `json:"gram,omitempty" gorm:"type:varchar(36);not null"`
	Harga string `json:"harga,omitempty" gorm:"type:varchar(36);not null"`
	Norek string `json:"norek,omitempty" gorm:"type:varchar(36);not null"`
}
