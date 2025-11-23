package models

import (
	"time"

	db "example.com/topup-restapi/DB"
)

type TopUp struct {
	ID          int64
	UserID      int
	TitleGame   string    `json:"title_game" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Price       string    `json:"price" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
}

var top_ups = []TopUp{}

func (e TopUp) Save() error {
	quer := `INSERT INTO topups 
	(titlegame, description, price, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	//statemnt, err := db.DB.Prepare(quer)
	statemnt, err := db.DB.Prepare(quer)
	if err != nil {
		return err
	}

	defer statemnt.Close()

	results, err := statemnt.Exec(e.TitleGame, e.Description, e.Price, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := results.LastInsertId()
	e.ID = id
	return err
	//topups = append(topups, e)
}

func GetAllItem() ([]TopUp, error) {
	qry := "SELECT * FROM topups"
	rows, err := db.DB.Query(qry)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var Topups []TopUp
	for rows.Next() {
		var topups TopUp
		err := rows.Scan(&topups.ID, &topups.TitleGame, &topups.Description, &topups.Price, &topups.DateTime, &topups.UserID)

		if err != nil {
			return nil, err
		}

		Topups = append(Topups, topups)
	}
	return Topups, nil
}

func GetTopUpByID(id int64) (*TopUp, error) {
	quer := "SELECT * FROM topups WHERE id = ?"
	rows := db.DB.QueryRow(quer, id)

	var topups TopUp
	err := rows.Scan(&topups.ID, &topups.TitleGame, &topups.Description, &topups.Price, &topups.DateTime, &topups.UserID)
	if err != nil {
		return nil, err
	}
	return &topups, err
}

func (topup TopUp) Update() error {
	quer := `
	UPDATE topups SET
	titlegame = ?, description = ?, price = ?, dateTime = ?
	WHERE id = ? 
	`
	statemnt, err := db.DB.Prepare(quer)
	if err != nil {
		return err
	}
	defer statemnt.Close()
	_, err = statemnt.Exec(topup.TitleGame, topup.Description, topup.Price, topup.DateTime, topup.ID)
	return err

}

func (topup TopUp) Delete() error {
	quer := "DELETE FROM topups WHERE id = ?"
	sttmnt, err := db.DB.Prepare(quer)
	if err != nil {
		return nil
	}

	defer sttmnt.Close()
	_, err = sttmnt.Exec(topup.ID)
	return err
}
