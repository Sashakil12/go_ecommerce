package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUsersHandler(w http.ResponseWriter, r *http.Request) {

	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println("Error parsing json", err)
		utils.SendError(w, 400, "Invalid json for users")
		return
	}
	allUsers := database.List()
	newUser.Id = len(allUsers) + 1
	newUser.Store()
	utils.SendData(w, newUser, 201)

}
