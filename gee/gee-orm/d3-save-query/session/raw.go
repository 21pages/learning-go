package session

import (
	"database/sql"
	"geeorm/dialect"
	"geeorm/schema"
	"strings"

	"geeorm/log"
)

type Session struct {
	db       *sql.DB         //通用
	sql      strings.Builder //query
	sqlVars  []interface{}   //arg
	dialect  dialect.Dialect //数据库差异
	refTable *schema.Schema  //转换
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect}
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.db.Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// query at most more row
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.db.QueryRow(s.sql.String(), s.sqlVars...)
}

// query all rows
func (s *Session) QueryRows(rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.db.Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error("Query:", err)
	}
	return
}

func (s *Session) Raw(sql string, args ...interface{}) *Session {
	s.Clear()
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, s.sqlVars...)
	return s
}
