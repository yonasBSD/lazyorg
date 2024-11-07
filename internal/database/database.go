package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/HubertBel/lazyorg/internal/calendar"
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
            description TEXT,
            location TEXT,
            time DATETIME NOT NULL,
            duration REAL NOT NULL,
            frequency INTEGER,
            occurence INTEGER
        )`)
	if err != nil {
		return err
	}
	_, err = database.Db.Exec(
		`CREATE TABLE IF NOT EXISTS notes (
            id INTEGER NOT NULL PRIMARY KEY,
            content TEXT NOT NULL,
            updated_at DATETIME NOT NULL
        )`)
	if err != nil {
		return err
	}

	return nil
}

func (database *Database) AddRecurringEvents(e *calendar.Event) ([]int, error) {
	events := e.GetReccuringEvents()
	ids := make([]int, 0, len(events))

	tx, err := database.Db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(`INSERT INTO events (name, description, location, time, duration, frequency, occurence) VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()

	for _, event := range events {
		result, err := stmt.Exec(
			event.Name,
			event.Description,
			event.Location,
			event.Time,
			event.DurationHour,
			event.FrequencyDay,
			event.Occurence,
		)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		id, err := result.LastInsertId()
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		ids = append(ids, int(id))
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return ids, nil
}

func (database *Database) GetEventsByDate(date time.Time) ([]*calendar.Event, error) {
	formattedDate := fmt.Sprintf("%04d-%02d-%02d", date.Year(), date.Month(), date.Day())

	var events []*calendar.Event
	rows, err := database.Db.Query(
		`SELECT * FROM events WHERE date(time) = ?;`, formattedDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event calendar.Event
		if err := rows.Scan(
			&event.Id, &event.Name, &event.Description, &event.Location, &event.Time, &event.DurationHour, &event.FrequencyDay, &event.Occurence,
		); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	return events, nil
}

func (database *Database) DeleteEvent(id int) error {
	result, err := database.Db.Exec("DELETE FROM events WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("event with id %d not found", id)
	}

	return nil
}

func (database *Database) DeleteEventsByName(name string) error {
	tx, err := database.Db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec("DELETE FROM events WHERE name = ?", name)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("no events found with the name: %s", name)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (database *Database) SaveNote(content string) error {
    tx, err := database.Db.Begin()
    if err != nil {
        return err
    }

    _, err = tx.Exec("DELETE FROM notes")
    if err != nil {
        tx.Rollback()
        return err
    }

    _, err = tx.Exec(
        "INSERT INTO notes (content, updated_at) VALUES (?, datetime('now'))",
        content,
    )
    if err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit()
}

func (database *Database) GetLatestNote() (string, error) {
    var content string
    err := database.Db.QueryRow(
        "SELECT content FROM notes ORDER BY updated_at DESC LIMIT 1",
    ).Scan(&content)
    
    if err == sql.ErrNoRows {
        return "", nil
    }
    return content, err
}
