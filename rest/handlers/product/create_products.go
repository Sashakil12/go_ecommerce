package product

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateProductsHandler(w http.ResponseWriter, r *http.Request) {

	var newProduct repo.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println("Error parsing json")
		utils.SendError(w, 400, "Invalid json")
		return
	}
	allProducts, err := h.productRepo.List()
	if err != nil {
		utils.SendError(w, 500, "Internal server error")
		return
	}
	newProduct.Id = len(allProducts) + 1
	product, err := h.productRepo.Create(newProduct)
	if err != nil {
		utils.SendError(w, 500, "Internal server error")
		return
	}

	utils.SendData(w, product, 201)

}
