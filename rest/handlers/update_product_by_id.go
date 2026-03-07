package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.SendData(w, "Invalid product ID", 400)
		return
	}
	var updatedProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedProduct)
	if err != nil {
		utils.SendError(w, 401, "Invalid json")
		return
	}
	product := database.UpdateById(idInt, updatedProduct)
	if product != nil {
		utils.SendData(w, product, 200)
		return
	}
	utils.SendError(w, 404, "Product not found")
}
