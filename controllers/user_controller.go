package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nicopereiran7/first-api-go/common"
	"github.com/nicopereiran7/first-api-go/models"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	users := []models.User{}
	db := common.GetConnection()
	defer db.Close()

	db.Find(&users)

	json, _ := json.Marshal(users)

	common.SendResponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	id := mux.Vars(request)["id"]

	db := common.GetConnection()
	defer db.Close()

	db.Find(&user, id)

	if user.ID > 0 {
		json, _ := json.Marshal(user)
		common.SendResponse(writer, http.StatusOK, json)
	} else {
		common.SendError(writer, http.StatusNotFound)
	}
}

func Save(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	db := common.GetConnection()
	defer db.Close()

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		log.Fatal(err)
		common.SendError(writer, http.StatusBadRequest)
		return
	}

	err = db.Create(&user).Error

	if err != nil {
		log.Fatal(err)
		common.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(user)

	common.SendResponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	user := models.User{}

	db := common.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&user, id)

	if user.ID > 0 {
		db.Delete(writer, http.StatusNotFound)
	} else {
		common.SendError(writer, http.StatusNotFound)
	}

}
