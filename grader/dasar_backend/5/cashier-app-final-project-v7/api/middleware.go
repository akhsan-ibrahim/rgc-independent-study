package api

import (
	"a21hc3NpZ25tZW50/model"
	"context"
	"encoding/json"
	"net/http"
)

func (api *API) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		coo, errCoo := r.Cookie("session_token")
		// sessionToken := coo.Value // TODO: replace this

		// if errCoo != nil {
		if errCoo == http.ErrNoCookie {
			err := model.ErrorResponse{Error: "http: named cookie not present"}
			jsonData, _ := json.Marshal(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(401)
			w.Write(jsonData)
			return
		}
		// }

		sessionFound, err := api.sessionsRepo.CheckExpireToken(coo.Value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "username", sessionFound.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) Get(next http.Handler) http.Handler {
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
		// TODO: answer here
		next.ServeHTTP(w, r)
	})
}

func (api *API) Post(next http.Handler) http.Handler {
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
		// TODO: answer here
		next.ServeHTTP(w, r)
	})
}
