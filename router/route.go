package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mufti-shiddiq/go-test-api/handler"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	v1 := api.Group("/product")
	
	// routes
	v1.Get("/", handler.GetAllProducts)
	v1.Get("/:id", handler.GetSingleProduct)
	v1.Post("/", handler.CreateProduct)
	v1.Put("/:id", handler.UpdateProduct)
	v1.Delete("/:id", handler.DeleteProductByID)
}