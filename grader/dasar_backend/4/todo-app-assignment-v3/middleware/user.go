package middleware

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"time"

	"a21hc3NpZ25tZW50/model"
)

func isExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := model.ErrorResponse{}
		_, errCoo := r.Cookie("session_token")
		// fmt.Println(Coo)
		if errCoo != nil {
			if errCoo == http.ErrNoCookie {
				err.Error = "http: named cookie not present"
				jsonData, _ := json.Marshal(err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(401)
				w.Write(jsonData)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
