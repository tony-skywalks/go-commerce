package controllers

import (
	"net/http"

	"github.com/tony-skywalks/my-web/pkg/models"
	"github.com/tony-skywalks/my-web/pkg/utils"
)

var User models.User

func Accounts(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		new_user := &models.User{}
		utils.ParseBody(r, new_user)
		user_obj, err := new_user.CreateUser()
		if err != nil {
			utils.SendResponse(w, "403", err.Error(), make([]string, 0), http.StatusOK)
		} else {
			utils.SendResponse(w, "200", "", user_obj, http.StatusOK)
		}
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		utils.SendResponse(w, "405", "Method Not Allowed", make([]string, 0), http.StatusMethodNotAllowed)
	}

	user := &models.User{}
	utils.ParseBody(r, user)
	user.Autheticate()

	if user.Username != "" {
		utils.SendResponse(w, "200", "", user, http.StatusOK)
	} else {
		utils.SendResponse(w, "403", "Authentication Failed", user, http.StatusOK)
	}
}
