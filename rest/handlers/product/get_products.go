package product

import (
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	list, err := h.productRepo.List()
	if err != nil {
		utils.SendError(w, 500, "Internal server error")
		return
	}
	utils.SendData(w, list, 200)
}
