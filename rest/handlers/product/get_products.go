package product

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

func (h *Handler) GetProductsHandler(w http.ResponseWriter, r *http.Request) {

	utils.SendData(w, database.List(), 200)
}
