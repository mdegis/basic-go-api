package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Melaba!\n")
}

func ImageIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(images); err != nil {
		panic(err)
	}
}

func ImageShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var imageId int
	var err error
	if imageId, err = strconv.Atoi(vars["imageId"]); err != nil {
		panic(err)
	}
	image := RepoFindImage(imageId)
	// bulundu mu kontrol
	if image.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(image); err != nil {
			panic(err)
		}
		return
	}

	// 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

/*test:
curl \
  -F "uploadfile=@DOSYA" \
  -F "location=santa' secret shop" \
  localhost:8080/images
*/
func ImageCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	panic(err)
	if err != nil {
	}
	defer f.Close()
	io.Copy(f, file)
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	t := RepoCreateImage(Image{Location: strings.Join(r.Form["location"], ""),
		Path: dir + "/uploads/" + handler.Filename,
		Date: time.Now()})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
