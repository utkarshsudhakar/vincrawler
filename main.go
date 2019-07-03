package main

import (
	"net/http"

	"./service"
)

func main() {

	service.Init()
	http.ListenAndServe(":4041", nil)

}
