package repository

import (
	"encoding/json"
	"fmt"

	"github.com/anaabdi/todo-go/helper"
	"github.com/lib/pq"
)

type User struct {
	ID            int64       `json:"id"`
	Username      string      `json:"username"`
	DeactivatedAt pq.NullTime `json:"deactivated_at"`
	Profile       Profile     `json:"profile"`
}

type Profile struct {
	UserID     string `json:"user_id"`
	Salt       []byte `json:"salt"`
	Password   []byte `json:"password"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Permission string `json:"permission"`
}

func (user *User) SetPassword(password string) {
	/* ph, err := helper.GeneratePasswordHash(password, nil)
	if err != nil {
		return
	}
	user.Profile.Salt = (*ph).Salt
	user.Profile.Password = (*ph).Hash */

	hashedPassword, err := helper.GeneratePasswordHash(password)
	if err != nil {
		return
	}
	user.Profile.Password = hashedPassword
}

func (user *User) SetUserID() {
	user.Profile.UserID = fmt.Sprintf("U%08d", user.ID)
}

func GetUserByUsername(username string) (*User, error) {
	var (
		user    User
		profile []byte
	)
	row := db.QueryRow(queryUserByUsername, username)
	if err := row.Scan(&user.ID, &user.Username, &profile, &user.DeactivatedAt); err != nil {
		return nil, err
	}

	user.readProfileBlob(profile)
	return &user, nil
}

func AddNewUser(u *User) error {
	profile, err := u.createProfileBlob()
	if err != nil {
		return err
	}

	row := db.QueryRow(queryInsert, u.Username, profile, "admin", "admin")
	if err := row.Scan(&u.ID); err != nil {
		return err
	}

	u.SetUserID()
	profile, err = u.createProfileBlob()
	if err != nil {
		return err
	}

	resUpdate := db.MustExec(queryUpdateUserID, profile, u.ID)
	if _, err := resUpdate.RowsAffected(); err != nil {
		return err
	}

	return nil
}

func (user *User) createProfileBlob() (profile []byte, err error) {
	return json.Marshal(user.Profile)
}

func (user *User) readProfileBlob(profile []byte) (err error) {
	return json.Unmarshal(profile, &user.Profile)
}
