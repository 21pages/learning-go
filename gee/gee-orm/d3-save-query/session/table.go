package session

import (
	"fmt"
	"geeorm/log"
	"geeorm/schema"
	"reflect"
	"strings"
)

// Model set session refTable
func (s *Session) Model(value interface{}) *Session {
	if s.refTable == nil || reflect.TypeOf(s.refTable) != reflect.TypeOf(value) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("refTable is nil")
		return nil
	}
	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.RefTable()
	var columns []string
	for _, v := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", v.Name, v.Type, v.Tag)) //"Age INTEGER NOT NULL"
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("create table %s (%s)", table.Name, desc)).Exec()
	if err != nil {
		log.Error("CreateTable:", err)
	}
	return err
}

func (s *Session) DropTable() error {
	table := s.RefTable()
	_, err := s.Raw(fmt.Sprintf("drop table %s if exists", table.Name)).Exec()
	if err != nil {
		log.Error("DropTable:", err)
	}
	return err
}

func (s *Session) HasTable() bool {
	table := s.RefTable()
	row := s.Raw(s.dialect.TableExistSQL(table.Name)).QueryRow()
	var dest string
	_ = row.Scan(&dest)
	return dest == table.Name
}
