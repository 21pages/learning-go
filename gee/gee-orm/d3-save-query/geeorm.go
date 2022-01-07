package geeorm

import (
	"database/sql"
	"fmt"
	"geeorm/dialect"

	"geeorm/log"

	"geeorm/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (*Engine, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error("Open:", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Error("Ping:", err)
		db.Close()
		return nil, err
	}
	// make sure dialect exists
	dialect, ok := dialect.GetDialect(driver)
	if !ok {
		log.Error("no dialect of", driver)
		db.Close()
		return nil, fmt.Errorf("no dialect of %s", driver)
	}

	log.Info("connect success!", driver, source)
	return &Engine{db: db, dialect: dialect}, nil
}

func (e *Engine) Close() error {
	if err := e.db.Close(); err != nil {
		log.Error("Close:", err)
		return err
	}
	return nil
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}
