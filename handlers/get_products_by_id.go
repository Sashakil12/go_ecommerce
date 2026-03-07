package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"log"
	"net/http"
	"strconv"
)

func GetProductsByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.SendData(w, "Invalid product ID", 400)
		return
	}

	for idx, product := range database.ProductList {
		log.Println("checking product with id:", idx, product)
		log.Println("product id:", product.Id, "requested id:", idInt)
		if product.Id == idInt {
			utils.SendData(w, database.ProductList[idx], 200)
			return
		}
	}
	utils.SendData(w, "Product not found", 404)
}
