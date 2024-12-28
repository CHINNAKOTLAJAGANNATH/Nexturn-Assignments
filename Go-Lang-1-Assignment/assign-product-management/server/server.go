package server

import (
	"assign-product-management/handlers"
	"fmt"
	"net/http"
)

func Start() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/contact", handlers.ContactHandler)
	// Request URL: http://localhost:8080/Products
	http.HandleFunc("/products", handlers.GetAllProducts)
	// Request URL: http://localhost:8080/Product/name?name=laptop
	http.HandleFunc("/product/name", handlers.GetProductByName)
	// Request URL: http://localhost:8080/product/category?category=electronics
	http.HandleFunc("/product/category", handlers.GetProductsByCategory)

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetProductByID(w, r)
		case http.MethodPost:
			handlers.AddProduct(w, r)
		case http.MethodPut:
			handlers.UpdateProduct(w, r)
		case http.MethodDelete:
			handlers.DeleteProduct(w, r)
		default:
			http.Error(w, "Invalid Request method", http.StatusMethodNotAllowed)
		}
	})

	// http://localhost:8080/filterandsort?sort=asc
	http.HandleFunc("/filterandsort", handlers.GetProducts)

	fmt.Println("server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("error", err)
	}
}
