package main

import (
	"log"
	"net/http"

	"github.com/mzjp2/helm-template-preview/render"
)

func main() {
	http.HandleFunc("/template", render.HandleRenderTemplate)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
