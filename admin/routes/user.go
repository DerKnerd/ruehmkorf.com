package routes

import (
	"net/http"
	"ruehmkorf.com/database/models"
	httpUtils "ruehmkorf.com/utils/http"
)

func UserList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		users, err := models.FindAllUsers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		httpUtils.RenderAdmin("admin/templates/user/overview.gohtml", users, w)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

type userData struct {
	Message   string
	Name      string
	Email     string
	Activated bool
}

func UserNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		httpUtils.RenderAdmin("admin/templates/user/new.gohtml", userData{Activated: true}, w)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/user/new.gohtml", userData{Message: err.Error()}, w)
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		activated := r.FormValue("activated") == "on"

		user := models.User{
			Name:      name,
			Email:     email,
			Password:  password,
			Activated: activated,
		}

		err = models.CreateUser(user)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/user/new.gohtml", userData{
				Message:   "Benutzer konnte nicht gespeichert werden",
				Name:      name,
				Email:     email,
				Activated: activated,
			}, w)
			return
		}

		http.Redirect(w, r, "/admin/user", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func UserEdit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		user, err := models.FindUserById(r.URL.Query().Get("id"))
		data := userData{}
		if err != nil {
			data.Message = "Benutzer nicht gefunden"
		} else {
			data.Name = user.Name
			data.Email = user.Email
			data.Activated = user.Activated
		}
		httpUtils.RenderAdmin("admin/templates/user/edit.gohtml", data, w)
	} else if r.Method == http.MethodPost {
		user, err := models.FindUserById(r.URL.Query().Get("id"))
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/user/edit.gohtml", userData{Message: "Benutzer nicht gefunden"}, w)
			return
		}

		err = r.ParseForm()
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/user/edit.gohtml", userData{Message: err.Error()}, w)
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		activated := r.FormValue("activated") == "on"
		password := r.FormValue("password")

		user.Name = name
		user.Email = email
		user.Activated = activated
		user.Password = password

		err = models.UpdateUser(*user, password != "")
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/user/edit.gohtml", userData{
				Message:   "Benutzer konnte nicht gespeichert werden",
				Name:      name,
				Email:     email,
				Activated: activated,
			}, w)
			return
		}

		http.Redirect(w, r, "/admin/user", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		user, err := models.FindUserById(id)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/user/delete.gohtml", userData{
				Message: "Benutzer nicht gefunden",
			}, w)
			return
		}

		httpUtils.RenderAdmin("admin/templates/user/delete.gohtml", userData{
			Name: user.Name,
		}, w)
	} else if r.Method == http.MethodPost {
		id := r.URL.Query().Get("id")
		_, err := models.FindUserById(id)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/user/delete.gohtml", userData{
				Message: "Benutzer nicht gefunden",
			}, w)
			return
		}

		err = models.DeleteUser(id)
		if err != nil {
			httpUtils.RenderAdmin("admin/templates/user/delete.gohtml", userData{
				Message: "Benutzer nicht gefunden",
			}, w)
			return
		}

		http.Redirect(w, r, "/admin/user", http.StatusFound)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
