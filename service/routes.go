package service

import "net/http"

func Init() {
	http.HandleFunc("/getVin", getVin)
}
