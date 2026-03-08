package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductsByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.SendData(w, "Invalid product ID", 400)
		return
	}
	product := database.GetById(idInt)
	if product != nil {
		utils.SendData(w, product, 200)
		return
	}
	utils.SendError(w, 404, "Product not found")
}
