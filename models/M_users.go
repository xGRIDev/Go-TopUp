package models

import db "example.com/topup-restapi/DB"

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
	rsl, err := sttmt.Exec(usr.Email, usr.Password)
	if err != nil {
		return err
	}
	userID, err := rsl.LastInsertId()
	usr.ID = userID

	return err
}
