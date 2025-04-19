package users

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type UserHandler struct {
	service *UserService
}

func Setup(s *mux.Router, db *gorm.DB) {
	service := NewUserService(db)
	handler := &UserHandler{service}

	r := s.PathPrefix("/users").Subrouter()
	r.StrictSlash(true)

	r.HandleFunc("/", handler.createUser).Methods("POST")
	r.HandleFunc("/sign-in", handler.signIn).Methods("POST")
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error"))
		return
	}
	defer r.Body.Close()

	var data CreateUser
	err = json.Unmarshal(body, &data)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error"))
		return
	}

	user, err := h.service.createUser(data)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error"))
		return
	}

	w.WriteHeader(201)
	w.Write(res)
}

func (h *UserHandler) signIn(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error"))
		return
	}

	var data SignIn
	err = json.Unmarshal(body, &data)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error"))
		return
	}

	session, err := h.service.signIn(data)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	res, err := json.Marshal(session)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error"))
		return
	}

	w.WriteHeader(201)
	w.Write(res)
}
