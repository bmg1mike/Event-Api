package models

import (
	"fmt"
	"time"

	"events.com/db"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

// var events = []Event{}

func (e Event) Save() error {
	// Save the event to the database
	result, err := db.DB.Exec("INSERT INTO events(name, description, location, date_time, user_id) VALUES(?,?,?,?,?)", e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	// events = append(events, e)
	// db.DB.Close()
	if err != nil {
		return err
	}
	value, err := result.RowsAffected()

	if err != nil {
		return err
	}

	fmt.Println("the number of rows affected are" + string(value))
	return nil
}

func GetAllEvents() ([]Event, error) {
	rows, err := db.DB.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := []Event{}
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil

}

func GetEventById(id int) (Event, error) {
	row := db.DB.QueryRow("SELECT * FROM events WHERE id = ?", id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}

func UpdateEvent(event Event) error {
	_, err := db.DB.Exec("UPDATE events SET name = ?, description = ?, location = ?, date_time = ?, user_id = ? WHERE id = ?", event.Name, event.Description, event.Location, event.DateTime, event.UserId, event.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEvent(id int) error {
	_, err := db.DB.Exec("DELETE FROM events WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
