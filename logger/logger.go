package logger

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	pglorus "gopkg.in/gemnasium/logrus-postgresql-hook.v1"
)

func GOGOGO() {
	db, err := sql.Open("postgres", "user=postgres dbname=postgres host=postgres sslmode=disable")
	if err != nil {
		fmt.Sprintf("Can't connect to postgresql database: %s", err)
	}
	defer db.Close()

	hook := pglorus.NewHook(db, map[string]interface{}{"this": "is logged every time"})
	hook.InsertFunc = func(db *sql.DB, entry *logrus.Entry) error {
		jsonData, err := json.Marshal(entry.Data)
		if err != nil {
			return err
		}

		_, err = db.Exec("INSERT INTO another_logs_table(level, message, message_data, created_at) VALUES ($1,$2,$3,$4);", entry.Level, entry.Message, jsonData, entry.Time)
		return err
	}

	logrus.AddHook(hook)

}
