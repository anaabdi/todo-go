package handler

import (
	"log"
	"net/http"

	"github.com/anaabdi/todo-go/helper"
	"github.com/anaabdi/todo-go/repository"
)

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req loginReq

	if err := parseRequest(r, &req); err != nil {
		log.Printf("err: %s", err.Error())
		Respond(w, http.StatusBadRequest, "parsing error")
		return
	}

	// Validate

	user, err := repository.GetUserByUsername(req.Username)
	if err != nil {
		log.Printf("err: %s", err.Error())
		Respond(w, http.StatusBadRequest, "failed to get user")
		return
	}

	if err := helper.VerifyPassword(user.Profile.Password, req.Password); err != nil {
		log.Printf("err: invalid password for user %s", user.Username)
		Respond(w, http.StatusBadRequest, "invalid password")
		return
	}

	/* ph, err := helper.GeneratePasswordHash(req.Password)
	if err != nil {
		log.Printf("err: %s", err.Error())
		Respond(w, http.StatusBadRequest, "failed to get user")
		return
	}

	if !bytes.Equal((*ph).Hash, user.Profile.Password) {
		log.Printf("err: invalid password for user %s", user.Username)
		Respond(w, http.StatusBadRequest, "invalid password")
		return
	} */

	accessTokenClaims := helper.JWTClaims{
		Email:      user.Profile.Email,
		Name:       user.Profile.Name,
		Phone:      user.Profile.Phone,
		UserID:     user.Profile.UserID,
		Username:   user.Username,
		Permission: user.Profile.Permission,
	}

	accessToken, err := helper.GenerateJWT(accessTokenClaims, helper.AccessTokenExpiration)
	if err != nil {
		log.Printf("err: %s", err.Error())
		Respond(w, http.StatusBadRequest, "failed to generate access token")
		return
	}

	refreshTokenClaims := helper.JWTClaims{
		Username: user.Username,
	}
	refreshToken, err := helper.GenerateJWT(refreshTokenClaims, helper.RefreshTokenExpiration)
	if err != nil {
		log.Printf("err: %s", err.Error())
		Respond(w, http.StatusBadRequest, "failed to generate refresh token")
		return
	}

	Respond(w, http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func DecodeToken() string {

}
