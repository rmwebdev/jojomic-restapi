package model

type Rekening struct {
	Base
	RekeningAPI
}

type RekeningAPI struct {
	Saldo string `json:"saldo,omitempty" gorm:"type:varchar(36);not null"`
	Norek string `json:"norek,omitempty" gorm:"type:varchar(36);not null"`
}

type RekeningData struct {
	Norek string `json:"norek,omitempty" gorm:"type:varchar(36);not null"`
}
