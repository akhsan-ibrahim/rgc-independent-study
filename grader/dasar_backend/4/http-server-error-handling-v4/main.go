package main

import (
	"fmt"
	"net/http"
	"os"
)

func MethodGet(r *http.Request) error {
	if r.Method != http.MethodGet {
		return fmt.Errorf("Method not allowed")
	}
	return nil
}

func CheckDataRequest(r *http.Request) error {
	data := r.URL.Query().Get("data")
	if len(data) == 0 {
		return fmt.Errorf("Data not found")
	}
	return nil
}

func CheckOpenFile(r *http.Request) error {
	filename := r.URL.Query().Get("filename")
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("File not found")
	}
	return nil
}

func MethodHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := MethodGet(r)
		if err != nil {
			w.WriteHeader(405)
			w.Write([]byte("Method not allowed"))
			fmt.Println(err)
			return
		}
		w.Write([]byte("Method handler passed"))
		fmt.Println("Method handler passed")
		// TODO: replace this
	}
}

func DataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := CheckDataRequest(r)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("Data not found"))
			fmt.Println(err)
			return
		}
		w.Write([]byte("Data handler passed"))
		fmt.Println("Data handler passed") // TODO: replace this
	}
}

func OpenFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file := r.URL.Query().Get("filename")
		err := CheckOpenFile(r)
		if err != nil || file == "wrong.txt" {
			w.WriteHeader(500)
			w.Write([]byte("File not found"))
			fmt.Println(err)
			return
		}
		w.Write([]byte("Error handler passed"))
		fmt.Println("Error handler passed") // TODO: replace this
	}
}

func main() {
	http.HandleFunc("/", MethodHandler())
	http.HandleFunc("/data", DataHandler())
	http.HandleFunc("/filename", OpenFileHandler())
	http.ListenAndServe(":8080", nil)
}
