package migrations

import "github.com/rmwebdev/jojomic-restapi/app/model"

// ModelMigrations models to automigrate
var ModelMigrations = []interface{}{
	&model.Harga{},
	&model.Topup{},
	&model.Rekening{},
	&model.Transaksi{},
}
