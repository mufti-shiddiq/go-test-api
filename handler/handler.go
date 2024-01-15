package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mufti-shiddiq/go-test-api/database"
	"github.com/mufti-shiddiq/go-test-api/model"
)

//Create a product
func CreateProduct(c *fiber.Ctx) error {
	db := database.DB.Db
	product := new(model.Product)
	
	// Store the body in the product and return error if encountered
	err := c.BodyParser(product)
	if err != nil {
		return c.Status(406).JSON(fiber.Map{"status": "error", "message":  "Something's wrong with your input", "data": err})
	}

	err = db.Create(&product).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message":  "Could not create product", "data": err})
	}

	// Return the created product
	return c.Status(201).JSON(fiber.Map{"status": "success", "message":  "Product has created", "data": product})
}

// Get All Products from db
func GetAllProducts(c *fiber.Ctx) error {
	db := database.DB.Db
	var products []model.Product

	// find all products in the database
	db.Find(&products)

	// If no product found, return an error
	if len(products) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Products not found", "data": nil})
	}

	// return products
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Products Found", "data": products})
}

// GetSingleProduct from db
func GetSingleProduct(c *fiber.Ctx) error {
	db := database.DB.Db

	// get id params
	id := c.Params("id")
	var product model.Product

	// find single product in the database by id
	db.Find(&product, "id = ?", id)
	if product.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Product Found", "data": product})
}

// update a product in db
func UpdateProduct(c *fiber.Ctx) error {
	type updateProduct struct {
		Name string `json:"name"`
		Price int `json:"price"`
		URLImage string `json:"url_image"`
	}

	db := database.DB.Db
	var product model.Product

	// get id params
	id := c.Params("id")

	// find single product in the database by id
	db.Find(&product, "id = ?", id)
	if product.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found", "data": nil})
	}

	var updateProductData updateProduct
	err := c.BodyParser(&updateProductData)
	if err != nil {
		return c.Status(406).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	product.Name = updateProductData.Name
	product.Price = updateProductData.Price
	product.URLImage = updateProductData.URLImage

	// Save the Changes
	db.Save(&product)

	// Return the updated product
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Success update product", "data": product})
}

// delete product in db by ID
func DeleteProductByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var product model.Product

	// get id params
	id := c.Params("id")

	// find single product in the database by id
	db.Find(&product, "id = ?", id)
	if product.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found", "data": nil})
	}

	err := db.Delete(&product, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete product", "data": nil})
	}
	
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Product deleted"})
}