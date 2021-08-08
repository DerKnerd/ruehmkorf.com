package routes

import (
	"encoding/json"
	"github.com/lib/pq"
	"net/http"
	"ruehmkorf.com/database/models"
)

func UserAction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		if id != "" {
			userDetails(w, id)
		} else {
			userList(w)
		}
	} else if r.Method == http.MethodPost {
		userNew(w, r)
	} else if r.Method == http.MethodPut {
		userEdit(w, r)
	} else if r.Method == http.MethodDelete {
		userDelete(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func userDetails(w http.ResponseWriter, id string) {
	user, err := models.FindUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func userList(w http.ResponseWriter) {
	users, err := models.FindAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(users)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type userData struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Activated bool   `json:"activated"`
	Password  string `json:"password"`
}

func userNew(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data userData
	err := decoder.Decode(&data)

	user := models.User{
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		Activated: data.Activated,
	}

	err = models.CreateUser(user)
	if conv, ok := err.(*pq.Error); ok == true && conv.Code == "23505" {
		w.WriteHeader(http.StatusConflict)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func userEdit(w http.ResponseWriter, r *http.Request) {
	user, err := models.FindUserById(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var bodyData userData
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&bodyData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Name = bodyData.Name
	user.Email = bodyData.Email
	user.Activated = bodyData.Activated
	user.Password = bodyData.Password

	err = models.UpdateUser(*user, bodyData.Password != "")
	if conv, ok := err.(*pq.Error); ok == true && conv.Code == "23505" {
		w.WriteHeader(http.StatusConflict)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func userDelete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := models.FindUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = models.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
