package middleware

import "net/http"

func CorsWithPreflight(next http.Handler) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Method", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Habib")
		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	}
	return http.HandlerFunc(handleAllReq)
}
