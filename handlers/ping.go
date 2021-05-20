package handlers

import "net/http"

func Ping(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("pong"))
}
