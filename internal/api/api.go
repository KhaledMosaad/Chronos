package api

import (
	"fmt"
	"net/http"
)

func init() {

}

type Chan chan *http.Request

var ch Chan

func StartServingRestAPI() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, req *http.Request) {
		ch <- req
		fmt.Fprint(w, "notification sent")
	})
	http.ListenAndServe(":8080", nil)
}
