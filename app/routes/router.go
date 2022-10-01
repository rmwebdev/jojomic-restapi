package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rmwebdev/jojomic-restapi/app/controller"
	"github.com/rmwebdev/jojomic-restapi/app/services"
	"github.com/spf13/viper"
)

// Handle all request to route to controller
func Handle(app *fiber.App) {
	app.Use(cors.New())
	services.InitDatabase()

	api := app.Group(viper.GetString("ENDPOINT"))
	// fmt.Println(api)
	api.Post("/input-harga", controller.InsertHargaPost)
	api.Get("/check-harga", controller.CheckHargaGet)
	api.Post("/topup", controller.TopupPost)
	api.Post("/saldo", controller.CheckSaldoPost)
	api.Post("/buyback", controller.BuybackPost)
	api.Post("/mutasi", controller.CheckMutasiPost)

}
