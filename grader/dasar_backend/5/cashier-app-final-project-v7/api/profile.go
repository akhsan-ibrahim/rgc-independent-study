package api

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func (api *API) ImgProfileView(w http.ResponseWriter, r *http.Request) {
	// View with response image `img-avatar.png` from path `assets/images`
	imgByte, err := ioutil.ReadFile("./assets/images/img-avatar.png")
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("File not found"))
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(imgByte)
}

func (api *API) ImgProfileUpdate(w http.ResponseWriter, r *http.Request) {
	// Update image `img-avatar.png` from path `assets/images`
	if err := r.ParseMultipartForm(1024); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	imgUp, _, _ := r.FormFile("file-avatar")
	defer imgUp.Close()

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fileLoc := filepath.Join(dir, "./assets/images/img-avatar.png")
	targetFile, err := os.OpenFile(fileLoc, os.O_WRONLY|os.O_CREATE, 0646)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, imgUp); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	api.dashboardView(w, r)
}
