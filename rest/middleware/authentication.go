package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/config"
	"ecommerce/utils"
	"fmt"
	"net/http"
	"strings"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		jwtToken := strings.Split(authHeader, " ")[1]
		tokenparts := strings.Split(jwtToken, ".")
		if len(tokenparts) != 3 {
			utils.SendError(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		jwtHeader := tokenparts[0]
		jwtPayload := tokenparts[1]
		jwtSignature := tokenparts[2]
		message := jwtHeader + "." + jwtPayload
		secret := config.GetConfig().JwtSecret
		messageByteArr := []byte(message)
		secretByteArr := []byte(secret)
		h := hmac.New(sha256.New, secretByteArr)
		h.Write(messageByteArr)
		expectedSignature := h.Sum(nil)
		expectedSignatureBase64 := utils.Base64UrlEncode(expectedSignature)
		if expectedSignatureBase64 != jwtSignature {
			utils.SendError(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		next.ServeHTTP(w, r)
		fmt.Println("jwt token", jwtToken)
	})
}
