package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"regexp"

	"github.com/anaskhan96/soup"
)

type Body struct {
	VIN []string
}

func RespondWithJSON(vin []string, w http.ResponseWriter, r *http.Request) {

	body := Body{VIN: vin}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBody)

}

func getVin(res http.ResponseWriter, w *http.Request) {
	var vin []string

	url := w.URL.Query().Get("url")
	pageLen, _ := strconv.Atoi(w.URL.Query().Get("pages"))

	for j := 1; j <= pageLen; j++ {

		newurl := url + "?page=" + strconv.Itoa(j)

		resp, err := soup.Get(newurl)
		if err != nil {
			fmt.Println("error in fetching the data from url ")
		}
		//fmt.Println(resp)
		r, _ := regexp.Compile("vin\":\"(\\w.+?)\"")
		f := r.FindAllStringSubmatch(resp, -1)
		for i := range f {
			vin = append(vin, f[i][1])

		}
	}
	RespondWithJSON(vin, res, w)
}
