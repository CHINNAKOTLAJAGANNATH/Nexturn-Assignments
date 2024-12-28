package main

import (
	db "go-sqlite-crud-project-assignment/config"
	"go-sqlite-crud-project-assignment/controller"
	"go-sqlite-crud-project-assignment/repository"
	"go-sqlite-crud-project-assignment/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db.InitializeDatabase()

	// Create repository, service, and controller
	productRepo := repository.NewProductRepository(db.GetDB())
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	// Initialize Gin router
	r := gin.Default()

	// Routes
	r.POST("/products", productController.CreateProduct)
	r.GET("/products/:id", productController.GetProduct)
	r.GET("/products", productController.GetAllProducts)
	r.PUT("/products/:id", productController.UpdateProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)

	// Start server
	r.Run(":8080")

}

// model-repo-service-controller
// http://localhost:8080/products
// http://localhost:8080/products/1
