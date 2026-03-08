package handlers

import (
	"ecommerce/config"
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
		accessToken, err := utils.CreateJwt(config.GetConfig().JwtSecret, utils.Payload{
			Sub:       user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		})
		if err != nil {
			fmt.Println("Error creating jwt", err)
			utils.SendError(w, 500, "Error creating access token")
			return
		}
		utils.SendData(w, map[string]string{
			"access_token": accessToken,
		}, 200)
		return
	}
	utils.SendError(w, 401, "Unauthorized!")
}
