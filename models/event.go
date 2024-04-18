package models

import (
	"time"

	"ebrarcode.dev/restapi-go/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location`
	DateTime    time.Time `binding:"required" json:"date_time`
	UserID      int       `json:"user_id"`
}

var events = []Event{}

func (e Event) Save() error {

	query := "INSERT INTO events (name, description, location, date_time, user_id) VALUES (?, ?, ?, ?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err

}

func GetAllEvents() []Event {
	return events
}
