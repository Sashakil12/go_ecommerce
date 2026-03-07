package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func DeleteProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.SendData(w, "Invalid product ID", 400)
		return
	}
	product := database.DeleteById(idInt)
	if product != nil {
		utils.SendData(w, product, 200)
		return
	}
	utils.SendError(w, 404, "Product not found")
}
