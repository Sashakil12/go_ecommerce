package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateProductsHandler(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println("Error parsing json")
		utils.SendError(w, 400, "Invalid json")
		return
	}
	allProducts := database.List()
	newProduct.Id = len(allProducts) + 1
	product := database.Store(newProduct)
	utils.SendData(w, product, 201)

}
