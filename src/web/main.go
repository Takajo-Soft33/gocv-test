package main

import (
	"fmt"
	"ioutil"
	"log"
	"net/http"
	"os"
)

func uploadFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("baseimage")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()



		//TODO: 人物の切り抜き、テクスチャの貼り付け、など
	}
}

func main() {
	http.HandleFunc("/upload", uploadFunc)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
