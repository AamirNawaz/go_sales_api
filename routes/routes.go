package routes

import (
	"go_sales_api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	//authentication routes
	app.Post("/cashiers/:cashierId/login", controllers.Login)
	app.Post("/cashiers/:cashierId/logout", controllers.Logout)
	app.Get("/cashiers/:cashierId/passcode", controllers.Passcode)

	//Cashier routes
	app.Get("/cashiers", controllers.CashiersList)
	app.Get("/cashiers/:cashierId", controllers.GetCashierDetails)
	app.Post("/cashiers", controllers.CreateCashier)
	app.Delete("/cashiers/:cashierId", controllers.DeleteCashier)
	app.Put("/cashiers/:cashierId", controllers.UpdateCashier)

	//Category routes
	app.Get("/categories", controllers.CategoryList)
	app.Get("/categories/:categoryId", controllers.GetCategoryDetails)
	app.Post("/categories", controllers.CreateCategory)
	app.Delete("/categories/:categoryId", controllers.DeleteCategory)
	app.Put("/categories/:categoryId", controllers.UpdateCategory)

	//Products routes
	app.Get("/products", controllers.ProductList)
	app.Get("/products/:productId", controllers.GetProductDetails)
	app.Post("/products", controllers.CreateProduct)
	app.Delete("/products/:productId", controllers.DeleteProduct)
	app.Put("/products/:productId", controllers.UpdateProduct)

	//Payment routes
	app.Get("/payments", controllers.PaymentList)
	app.Get("/payments/:paymentId", controllers.GetPaymentDetails)
	app.Post("/payments", controllers.CreatePayment)
	app.Delete("/payments/:paymentId", controllers.DeletePayment)
	app.Put("/payments/:paymentId", controllers.UpdatePayment)

	//Order routes
	app.Get("/orders", controllers.OrdersList)
	app.Get("/orders/:orderId", controllers.OrderDetail)
	app.Post("/orders", controllers.CreateOrder)
	app.Post("/orders/subtotal", controllers.SubTotalOrder)
	app.Get("/orders/:orderId/download", controllers.DownloadOrder)
	app.Get("/orders/:orderId/check-download", controllers.CheckOrder)

	//reports
	app.Get("/revenues", controllers.GetRevenues)
	app.Get("/solds", controllers.GetSolds)

}
