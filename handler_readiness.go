package main

import "net/http"



func handlerRediness(w http.ResponseWriter, req *http.Request) {

	responseWithJson(w,200,struct{
		Success bool `json:"success"`
	}{
	Success: true,
	})

}