package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProductsHandler(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println("Error parsing json")
		http.Error(w, "Invalid json", 401)
	}
	newProduct.Id = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)
	utils.SendData(w, database.ProductList, 201)

}
