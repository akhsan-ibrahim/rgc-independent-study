package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
)

func Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			err := model.ErrorResponse{}
			err.Error = "Method is not allowed!"
			jsonData, _ := json.Marshal(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(405)
			w.Write(jsonData)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			err := model.ErrorResponse{}
			err.Error = "Method is not allowed!"
			jsonData, _ := json.Marshal(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(405)
			w.Write(jsonData)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			err := model.ErrorResponse{}
			err.Error = "Method is not allowed!"
			jsonData, _ := json.Marshal(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(405)
			w.Write(jsonData)
			return
		}
		next.ServeHTTP(w, r)
	})
}
