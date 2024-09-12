package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/HubertBel/go-organizer/cmd/types"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Db *sql.DB
}

func (database *Database) InitDatabase(path string) error {
	var err error

	database.Db, err = sql.Open("sqlite3", path)
	if err != nil {
		return err
	}

	_, err = database.Db.Exec(
		`CREATE TABLE IF NOT EXISTS events (
            id INTEGER NOT NULL PRIMARY KEY,
            name TEXT NOT NULL,
            time DATETIME NOT NULL,
            duration REAL
        )`)
	if err != nil {
		return err
	}

	return nil
}

func (database *Database) AddEvent(e *types.Event) (int, error) {
	result, err := database.Db.Exec(
		`INSERT INTO events (name, time, duration) VALUES(?, ?, ?);`,
		e.Name,
		e.Time,
		e.DurationHour)

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return int(id), nil
}

func (database *Database) GetEventsByDate(date time.Time) ([]types.Event, error) {
    formattedDate := fmt.Sprintf("%04d-%02d-%02d", date.Year(), date.Month(), date.Day())

	var events []types.Event
	rows, err := database.Db.Query(
		`SELECT * FROM events WHERE date(time) = ?;`, formattedDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event types.Event
        var id int // TODO handle id
		if err := rows.Scan(
			&id, &event.Name, &event.Time, &event.DurationHour,
		); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, err
}
