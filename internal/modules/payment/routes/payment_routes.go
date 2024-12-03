package routes

import (
	"github.com/gofiber/fiber/v2"
	mwMngr "github.com/jumayevgadam/tsu-toleg/internal/gateway/middleware"
	"github.com/jumayevgadam/tsu-toleg/internal/infrastructure/database"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payment/handler"
	"github.com/jumayevgadam/tsu-toleg/internal/modules/payment/service"
)

func Routes(f fiber.Router, mw *mwMngr.MiddlewareManager, dataStore database.DataStore) {
	// Init Service.
	Service := service.NewPaymentService(dataStore)
	// Init Handler.
	Handler := handler.NewPaymentHandler(Service)
	// put Route
	paymentGroup := f.Group("/payment")
	{
		paymentGroup.Post("/add", mwMngr.RoleBasedMiddleware(mw, "add:payment", dataStore), Handler.AddPayment())
	}

}
