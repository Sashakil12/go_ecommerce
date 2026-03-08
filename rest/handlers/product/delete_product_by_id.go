package product

import (
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.SendData(w, "Invalid product ID", 400)
		return
	}
	result, err := h.productRepo.Delete(idInt)
	if err != nil {
		utils.SendError(w, 500, "Internal server error")
		return
	}
	if result {
		utils.SendData(w, "Product deleted successfully", 200)
		return
	}
	utils.SendError(w, 404, "Product not found")
}
