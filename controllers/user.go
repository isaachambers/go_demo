package controllers

import (
	"demo_webservice/model"
	"encoding/json"
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Welcome to the server"))
}

func (uc *userController) getAll(writer http.ResponseWriter) {
	encodeResponseAsJson(model.GetUsers(), writer)
}

func (uc *userController) get(id int, writer http.ResponseWriter) {
	user, err := model.GetUserById(id)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	encodeResponseAsJson(user, writer)
}

func (uc *userController) post(writer http.ResponseWriter, request *http.Request) {
	user, err := uc.parseRequest(request)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Failed to parse request"))
		return
	}
	user, err = model.AddUser(user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJson(user, writer)
}

func (uc *userController) put(id int, writer http.ResponseWriter, request *http.Request) {
	user, err := uc.parseRequest(request)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Failed to parse request"))
		return
	}
	if id != user.Id {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("ID in request is not same as Id in URL"))
		return
	}
	user, err = model.UpdateUser(user)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJson(user, writer)
}

func (uc *userController) delete(id int, writer http.ResponseWriter) {
	err := model.DeleteUserById(id)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte(err.Error()))
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (uc *userController) parseRequest(r *http.Request) (model.User, error) {
	dec := json.NewDecoder(r.Body)
	var user = model.User{}
	err := dec.Decode(&user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
