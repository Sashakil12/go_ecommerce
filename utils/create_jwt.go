package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
type Payload struct {
	Sub            int    `json:"sub"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	IsShopperOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	byteArrHeader, err := json.Marshal(header)

	if err != nil {
		return "", err
	}
	headerBase64 := base64UrlEncode(byteArrHeader)
	dataArrData, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	payloadBase64 := base64UrlEncode(dataArrData)
	message := headerBase64 + "." + payloadBase64
	byteArraySecret := []byte(secret)

	byteArrMessage := []byte(message)
	h := hmac.New(sha256.New, byteArraySecret)
	h.Write(byteArrMessage)
	signature := h.Sum(nil)
	signatureBase64 := base64UrlEncode(signature)
	jwt := headerBase64 + "." + payloadBase64 + "." + signatureBase64
	return jwt, nil

}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
