package handler

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/anaabdi/todo-go/repository"
)

type registerUser struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Permission string `json:"permission"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req registerUser

	if err := parseRequest(r, &req); err != nil {
		log.Printf("err: %s", err.Error())
		Respond(w, http.StatusBadRequest, "parsing error")
		return
	}

	// Validate Input

	user := &repository.User{
		Username: req.Username,
		Profile: repository.Profile{
			Name:       req.Name,
			Email:      req.Email,
			Permission: req.Permission,
			Phone:      req.Phone,
		},
	}

	user.SetPassword(req.Password)

	if err := repository.AddNewUser(user); err != nil {
		log.Printf("err: %s", err.Error())
		Respond(w, http.StatusBadRequest, "failed to store user")
		return
	}

	Respond(w, http.StatusCreated, map[string]interface{}{
		"id":         user.ID,
		"user_id":    user.Profile.UserID,
		"username":   user.Username,
		"name":       user.Profile.Name,
		"permission": user.Profile.Permission,
	})

	log.Printf("successfully adding new user: %#v", user)
}

func getUserFromCtx(ctx context.Context) (*repository.User, error) {
	userCtx := ctx.Value("user")
	if u, ok := userCtx.(*repository.User); ok {
		return u, nil
	}

	return nil, errors.New("failed to get user from context")
}

func Me(w http.ResponseWriter, r *http.Request) {
	user, err := getUserFromCtx(r.Context())
	if err != nil {
		log.Printf("err: %s", err.Error())
		Respond(w, http.StatusBadRequest, "failed to get user")
		return
	}

	Respond(w, http.StatusOK, user)
	log.Printf("successfully get profile")
}
