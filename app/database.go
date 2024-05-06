package app

import (
	"database/sql"
	"golang-api-udemy/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:hwhwhwlol@tcp(localhost:33061)/tutorial_golang")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10)

	return db
}
