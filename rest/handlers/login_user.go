package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var loginUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginUser)
	if err != nil {
		fmt.Println("Error parsing json")
		utils.SendError(w, 400, "Invalid json for users")
		return
	}
	user := loginUser.Find(loginUser.Email, loginUser.Password)
	if user != nil {
		utils.SendData(w, user, 200)
		return
	}
	utils.SendError(w, 401, "Unauthorized")
}
