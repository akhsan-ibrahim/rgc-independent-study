package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
	"path"
	"text/template"
	"time"

	"github.com/google/uuid"
)

func (api *API) Register(w http.ResponseWriter, r *http.Request) {
	// Read username and password request with FormValue.
	creds := model.Credentials{}
	erReg := model.ErrorResponse{}

	name := r.FormValue("username")
	pw := r.FormValue("password")

	if name == "" || pw == "" {
		erReg.Error = "Username or Password empty"
		jsonData, _ := json.Marshal(erReg)
		w.WriteHeader(400)
		w.Write(jsonData)
		return
	}

	creds.Username = name
	creds.Password = pw

	// Handle request if creds is empty send response code 400, and message "Username or Password empty"

	err := api.usersRepo.AddUser(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	filepath := path.Join("views", "status.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	var data = map[string]string{"name": creds.Username, "message": "register success!"}
	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}
}

func (api *API) Login(w http.ResponseWriter, r *http.Request) {
	// Read usernmae and password request with FormValue.
	creds := model.Credentials{} // TODO: replace this
	errLog := model.ErrorResponse{}

	name := r.FormValue("username")
	pw := r.FormValue("password")
	creds.Username = name
	creds.Password = pw

	// Handle request if creds is empty send response code 400, and message "Username or Password empty"

	if name == "" || pw == "" {
		errLog.Error = "Username or Password empty"
		jsonData, _ := json.Marshal(errLog)
		w.WriteHeader(400)
		w.Write(jsonData)
		return
	}

	err := api.usersRepo.LoginValid(creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	// Generate Cookie with Name "session_token", Path "/", Value "uuid generated with github.com/google/uuid", Expires time to 5 Hour.

	cookie := &http.Cookie{
		Name:    "session_token",
		Path:    "/",
		Value:   uuid.NewString(),
		Expires: time.Now().Add(5 * time.Hour),
	}
	http.SetCookie(w, cookie)

	session := model.Session{
		Token:    cookie.Value,
		Username: name,
		Expiry:   cookie.Expires,
	}
	err = api.sessionsRepo.AddSessions(session)

	api.dashboardView(w, r)
}

func (api *API) Logout(w http.ResponseWriter, r *http.Request) {
	//Read session_token and get Value:
	coo, _ := r.Cookie("session_token")
	sessionToken := coo.Value

	api.sessionsRepo.DeleteSessions(sessionToken)

	//Set Cookie name session_token value to empty and set expires time to Now:
	cookie := &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	}
	http.SetCookie(w, cookie)

	filepath := path.Join("views", "login.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}
}
