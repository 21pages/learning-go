package geeorm

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("mysql", "sun:root@tcp(127.0.0.1:3306)/gee")
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
}
