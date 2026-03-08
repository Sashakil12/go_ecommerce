package product

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.SendData(w, "Invalid product ID", 400)
		return
	}
	var updatedProduct repo.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedProduct)
	if err != nil {
		utils.SendError(w, 401, "Invalid json")
		return
	}
	product, err := h.productRepo.Update(idInt, updatedProduct)
	if err != nil {
		utils.SendError(w, 500, "Internal server error")
		return
	}
	if product != nil {
		utils.SendData(w, product, 200)
		return
	}
	utils.SendError(w, 404, "Product not found")
}
