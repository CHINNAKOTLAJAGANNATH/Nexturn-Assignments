package handlers

import (
	"assign-product-management/models"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// Home Route or Default Handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go HTTP Server")
}

// About Route Handler
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page")
}

// Contact Route Handler
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the contact page")
}

// API Handler
func APIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(map[string]string{"message": "This is the API endpoint"})
	fmt.Fprintln(w, `{"message":"This is a api endpoint"}`)
}

// Adding temporary in-memory data for testing
var products = []models.Product{
	{ID: 1, Name: "laptop", Description: "A latest model laptop", Price: 56000, Category: "electronics"},
	{ID: 2, Name: "smart-phone", Description: "A latest model samrt phone", Price: 12000, Category: "electronics"},
	{ID: 3, Name: "Mouse", Description: "A latest model mouse", Price: 40000, Category: "Wireless"},
	{ID: 4, Name: "Keyboard", Description: "A latest model samrt keyboard", Price: 30000, Category: "Wireless"},
}

// GetAllProducts Route Handler
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProductByID Route Handler
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for _, product := range products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	http.Error(w, "Product Not Found", http.StatusNotFound)
}

// AddProduct Route Handler
func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	json.NewEncoder(w).Encode(newProduct)
}

// UpdateProduct Route Handler
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	for i, product := range products {
		if product.ID == id {
			updatedProduct.ID = product.ID // Preserve the original ID
			products[i] = updatedProduct
			json.NewEncoder(w).Encode(updatedProduct)
			return
		}
	}
	http.Error(w, "Product Not Found", http.StatusNotFound)
}

// DeleteProduct Route Handler
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}
	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	http.Error(w, "Product Not Found", http.StatusNotFound)
}

// --------------------

// New Functions for filtering and Sorting

// GetProductsByCategory Route Handler
func GetProductsByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	category := r.URL.Query().Get("category")
	if category == "" {
		http.Error(w, "Category is required", http.StatusBadRequest)
		return
	}
	var productsByCategory []models.Product
	for _, product := range products {
		if strings.EqualFold(product.Category, category) {
			productsByCategory = append(productsByCategory, product)
		}
	}
	if len(productsByCategory) == 0 {
		http.Error(w, "No products found in the category", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(productsByCategory)
}

// GetProductByName Route Handler
func GetProductByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Product name is required", http.StatusBadRequest)
		return
	}
	for _, product := range products {
		if strings.EqualFold(product.Name, name) {
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

// sortByPrice sorts products by price
func sortByPrice(products *[]models.Product, order string) {
	if strings.EqualFold(order, "asc") {
		sort.Slice(*products, func(i, j int) bool {
			return (*products)[i].Price < (*products)[j].Price
		})
	} else if strings.EqualFold(order, "desc") {
		sort.Slice(*products, func(i, j int) bool {
			return (*products)[i].Price > (*products)[j].Price
		})
	}
}

// GetProducts Route Handler
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	category := r.URL.Query().Get("category")
	name := r.URL.Query().Get("name")
	sortOrder := r.URL.Query().Get("sort")
	filteredProducts := products
	// Filter by category
	if category != "" {
		filteredProducts = filterByCategory(filteredProducts, category)
	}
	// Filter by name
	if name != "" {
		filteredProducts = filterByName(filteredProducts, name)
	}
	// Sort by price
	if sortOrder != "" {
		sortByPrice(&filteredProducts, sortOrder)
		if sortOrder != "asc" && sortOrder != "desc" {
			http.Error(w, "Invalid sort order. Please use 'asc' or 'desc'.", http.StatusBadRequest)
			return
		}
	}

	if len(filteredProducts) == 0 {
		http.Error(w, "No products found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(filteredProducts)
}

// filterByCategory filters products by category
func filterByCategory(products []models.Product, category string) []models.Product {
	var filtered []models.Product
	for _, product := range products {
		if strings.EqualFold(product.Category, category) {
			filtered = append(filtered, product)
		}
	}
	return filtered
}

// filterByName filters products by name
func filterByName(products []models.Product, name string) []models.Product {
	var filtered []models.Product
	for _, product := range products {
		if strings.Contains(strings.ToLower(product.Name), strings.ToLower(name)) {
			filtered = append(filtered, product)
		}
	}
	return filtered
}
