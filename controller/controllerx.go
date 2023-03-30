package controller

import (
	"GoTools/model"
	"encoding/json"
	"log"
	"net/http"
)

func LoginEndpoint(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		showResponse(w, 400, "Parse Error", "")
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	var user model.User
	if err := db.QueryRow("SELECT id, name, email, password from users where email = ? AND password = ?", email, password).Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
		log.Println(err.Error())
		showResponse(w, 400, "Login Failed", user)
		return
	}

	startScheduler(user)
	generateToken(w, user)
	showResponse(w, 200, "Login Success", user)
}

func LogoutEndpoint(w http.ResponseWriter, r *http.Request) {
	stopScheduler()
	resetUserToken(w)
	showResponse(w, 200, "Logout Success", "")
}

func GetUserListEndpoint(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	users = GetUsers()

	if users == nil {
		db := connect()
		defer db.Close()

		query := "SELECT id, name, email, password from users"

		rows, err := db.Query(query)
		if err != nil {
			log.Println(err)
			showResponse(w, 400, "Query Failed", "")
			return
		}

		var user model.User
		for rows.Next() {
			if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
				log.Println(err.Error())
				showResponse(w, 400, "Get Failed", "")
				return
			} else {
				users = append(users, user)
			}
		}
		SetUsers(users)
	}

	showResponse(w, 200, "Get User List Success", users)
}

func showResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	response := map[string]interface{}{
		"status":  status,
		"message": message,
		"data":    data,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}