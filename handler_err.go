package main

import "net/http"

func handlerError(w http.ResponseWriter, req *http.Request) {

	responseWithError(w, 400, "Something Went Wrong")

}