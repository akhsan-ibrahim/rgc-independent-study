package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/model"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	bodyData := model.Credentials{}
	err := model.ErrorResponse{}

	errBody := json.NewDecoder(r.Body).Decode(&bodyData)
	if errBody != nil {
		err.Error = "Internal Server Error"
		jsonData, _ := json.Marshal(err)
		w.WriteHeader(400)
		w.Write(jsonData)
		return
	}
	if bodyData.Username == "" || bodyData.Password == "" {
		err.Error = "Username or Password empty"
		jsonData, _ := json.Marshal(err)
		w.WriteHeader(400)
		w.Write(jsonData)
		return
	}
	if _, ok := db.Users[bodyData.Username]; ok {
		err.Error = "Username already exist"
		jsonData, _ := json.Marshal(err)
		w.WriteHeader(409)
		w.Write(jsonData)
		return
	}

	success := model.SuccessResponse{
		Username: bodyData.Username,
		Message:  "Register Success",
	}
	jsonData, _ := json.Marshal(success)
	w.Write(jsonData)

	db.Users[bodyData.Username] = bodyData.Password
}

func Login(w http.ResponseWriter, r *http.Request) {
	bodyData := model.Credentials{}
	err := model.ErrorResponse{}
	errBody := json.NewDecoder(r.Body).Decode(&bodyData)

	if errBody != nil {
		err.Error = "Internal Server Error"
		jsonData, _ := json.Marshal(err)
		w.WriteHeader(400)
		w.Write(jsonData)
		return
	}
	if bodyData.Username == "" || bodyData.Password == "" {
		err.Error = "Username or Password empty"
		jsonData, _ := json.Marshal(err)
		w.WriteHeader(400)
		w.Write(jsonData)
		return
	}

	if _, ok := db.Users[bodyData.Username]; !ok || bodyData.Password != db.Users[bodyData.Username] {
		err.Error = "Wrong User or Password!"
		jsonData, _ := json.Marshal(err)
		w.WriteHeader(401)
		w.Write(jsonData)
		return
	}

	idSession := uuid.NewString()
	exp := time.Now().Add(5 * time.Hour)
	cookie := &http.Cookie{
		Name:    "session_token",
		Value:   idSession,
		Expires: exp,
	}
	http.SetCookie(w, cookie)

	db.Sessions[idSession] = model.Session{
		Username: bodyData.Username,
		Expiry:   exp,
	}

	success := model.SuccessResponse{
		Username: bodyData.Username,
		Message:  "Login Success",
	}
	jsonData, _ := json.Marshal(success)
	w.Write(jsonData)
}

func AddToDo(w http.ResponseWriter, r *http.Request) {
	err := model.ErrorResponse{}
	coo, errCoo := r.Cookie("session_token")
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

	type ToDo struct {
		Task string `json:"task"`
		Done bool   `json:"done"`
	}

	todo := ToDo{}
	errBody := json.NewDecoder(r.Body).Decode(&todo)
	if errBody != nil {
		err.Error = "Internal Server Error"
		jsonData, _ := json.Marshal(err)
		w.WriteHeader(400)
		w.Write(jsonData)
		return
	}

	sesion := db.Sessions[coo.Value]
	success := model.SuccessResponse{
		Username: sesion.Username,
		Message:  "Task " + todo.Task + " added!",
	}
	jsonData, _ := json.Marshal(success)
	w.Write(jsonData)

	doModel := model.Todo{
		Id:   uuid.NewString(),
		Task: todo.Task,
		Done: todo.Done,
	}
	db.Task[sesion.Username] = append(db.Task[sesion.Username], doModel)
}

func ListToDo(w http.ResponseWriter, r *http.Request) {
	coo, _ := r.Cookie("session_token")
	sesion := db.Sessions[coo.Value]
	lsTodo := db.Task[sesion.Username]
	if db.Task == nil || len(db.Task) == 0 {
		err := model.ErrorResponse{}
		err.Error = "Todolist not found!"
		jsonData, _ := json.Marshal(err)
		w.WriteHeader(404)
		w.Write(jsonData)
		return
	}

	w.WriteHeader(200)
	jsonData, _ := json.Marshal(lsTodo)
	w.Write(jsonData)
}

func ClearToDo(w http.ResponseWriter, r *http.Request) {
	err := model.ErrorResponse{}
	coo, errCoo := r.Cookie("session_token")
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

	sesion := db.Sessions[coo.Value]
	success := model.SuccessResponse{
		Username: sesion.Username,
		Message:  "Clear ToDo Success",
	}
	jsonData, _ := json.Marshal(success)
	w.Write(jsonData)

	db.Task[sesion.Username] = []model.Todo{}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// username := fmt.Sprintf("%s", r.Context().Value("username"))
	coo, _ := r.Cookie("session_token")
	sesion := db.Sessions[coo.Value]
	success := model.SuccessResponse{
		Username: sesion.Username,
		Message:  "Logout Success",
	}
	jsonData, _ := json.Marshal(success)
	w.Write(jsonData)

	delete(db.Sessions, coo.Value)
}

func ResetToDo(w http.ResponseWriter, r *http.Request) {
	db.Task = map[string][]model.Todo{}
	w.WriteHeader(http.StatusOK)
}

type API struct {
	mux *http.ServeMux
}

func NewAPI() API {
	mux := http.NewServeMux()
	api := API{
		mux,
	}

	mux.Handle("/user/register", middleware.Post(http.HandlerFunc(Register)))
	mux.Handle("/user/login", middleware.Post(http.HandlerFunc(Login)))
	mux.Handle("/user/logout", middleware.Get(middleware.Auth(http.HandlerFunc(Logout))))
	mux.Handle("/todo/create", middleware.Post(http.HandlerFunc(AddToDo)))
	mux.Handle("/todo/read", middleware.Get(middleware.Auth(http.HandlerFunc(ListToDo))))
	mux.Handle("/todo/clear", middleware.Delete(http.HandlerFunc(ClearToDo)))
	mux.Handle("/todo/reset", http.HandlerFunc(ResetToDo))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}

func main() {
	mainAPI := NewAPI()
	mainAPI.Start()
}
