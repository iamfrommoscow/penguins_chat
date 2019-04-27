package db

import (
	"github.com/jackc/pgx"
	"chat/models"
	"fmt"
)

func Query(sql string, args ...interface{}) (*pgx.Rows, error) {
	if connection == nil {
		return nil, pgx.ErrDeadConn
	}

	tx, err := connection.Begin()
	if err != nil {
		return nil, err
	}

	rows, err := tx.Query(sql, args...)
	return rows, err
}

func Exec(sql string, args ...interface{}) (commandTag pgx.CommandTag, err error) {
	if connection == nil {
		return "", pgx.ErrDeadConn
	}

	tx, err := connection.Begin()
	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	tag, err := tx.Exec(sql, args...)
	if err != nil {
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		return "", err
	}

	return tag, nil
}

func RowsToMessages(rows *pgx.Rows) []models.Message {
	messages := []models.Message{}
	for rows.Next() {
		entry := models.Message{}
		if err := rows.Scan(&entry.ID, &entry.From, &entry.Text); err != nil {
			fmt.Println(err)
		}
		messages = append(messages, entry)
	}
	return messages
}
