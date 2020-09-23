package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()
	//Handle method accepts 2 parameters, 1. String, another is an object that implements the Handler interface.
	// In Go. No need to explicitly define that all just like in Java, In Golang, as long as the method signature exists on the object, it automatically qualifies
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJson(data interface{}, writer io.Writer) {
	enc := json.NewEncoder(writer)
	enc.Encode(data)
}
