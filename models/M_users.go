package models

import (
	"errors"

	db "example.com/topup-restapi/DB"
	utils "example.com/topup-restapi/Utils"
)

type Users struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (usr Users) Save() error {
	qry := "INSERT INTO users (email, password) VALUES (?, ?)"
	sttmt, err := db.DB.Prepare(qry)

	if err != nil {
		return err
	}

	defer sttmt.Close()
	hashedPass, err := utils.HashPass(usr.Password)
	if err != nil {
		return nil
	}
	rsl, err := sttmt.Exec(usr.Email, hashedPass)
	if err != nil {
		return err
	}
	userID, err := rsl.LastInsertId()
	usr.ID = userID

	return err
}

func (u Users) LoginCredentialValidate() error {
	qry := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(qry, u.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)

	if err != nil {
		return errors.New("Credentials is Invalid.")
	}

	passwordIsValid := utils.CheckedHashPass(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Credentials Invalid")
	}
	return nil
}
