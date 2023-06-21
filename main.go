package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"

	"golang.org/x/image/webp"
)

var imageUrl = ""

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
		}
		imageUrl = r.FormValue("imageUrl")
	}

	if r.Method == "GET" {
		fmt.Println("GET request")
		imageUrl = r.URL.Query().Get("imageUrl")
		fmt.Println(imageUrl)
	}

	fmt.Println("Image URL: " + imageUrl)
	imageResponse, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer imageResponse.Body.Close()
	convertedImage, err := webp.Decode(imageResponse.Body)
	if err != nil {
		fmt.Println(err)
		fmt.Println(err)
		return
	}

	err = png.Encode(w, convertedImage)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", `attachment;filename="image"; filename*=UTF-8''image`)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", uploadHandler)

	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatal(err)
	}
}
