package user

import (
	"ecommerce/repo"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateUsersHandler(w http.ResponseWriter, r *http.Request) {

	var newUser repo.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println("Error parsing json", err)
		utils.SendError(w, 400, "Invalid json for users")
		return
	}

	createdUser, err := h.userRepo.Create(newUser)
	if err != nil {
		fmt.Println("Error creating user", err)
		utils.SendError(w, 500, "Error creating user")
		return
	}
	utils.SendData(w, createdUser, 201)

}
