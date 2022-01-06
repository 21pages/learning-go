package geeorm

import (
	"database/sql"

	"geeorm/log"

	"geeorm/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) *Engine {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error("Open:", err)
		return nil
	}
	if err := db.Ping(); err != nil {
		log.Error("Ping:", err)
		db.Close()
		return nil
	}
	log.Info("connect success!", driver, source)
	return &Engine{db: db}
}

func (e *Engine) Close() error {
	if err := e.db.Close(); err != nil {
		log.Error("Close:", err)
		return err
	}
	return nil
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db)
}
