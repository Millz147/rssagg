package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, msg string){
	if code >499{
		log.Println("Responding with 5XX error",msg)
	}
type ErrorResponse struct{
		Error string `json:"error"`
	}
	responseWithJson(w,code,ErrorResponse{
		Error: msg,
	})
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {

	data,err:= json.Marshal(payload)

	if err!=nil{
		w.WriteHeader(500)
		return
	}
w.Header().Add("Content-Type","application/json")
w.WriteHeader(code)
w.Write(data)

}