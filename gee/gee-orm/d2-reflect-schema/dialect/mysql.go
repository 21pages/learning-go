package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type mysql struct{}

func init() {
	RegisterDialect("mysql", &mysql{})
}

func (m *mysql) DataTypeOf(typ reflect.Value) string {
	switch typ.Kind() {
	case reflect.Bool, reflect.Int8, reflect.Uint8:
		return "TINYINT"
	case reflect.Int16, reflect.Uint16:
		return "SMALLINT"
	case reflect.Int32, reflect.Uint32:
		return "INTEGER" //"INT"
	case reflect.Int, reflect.Uint, reflect.Uintptr, reflect.Int64, reflect.Uint64:
		return "BIGINT"
	case reflect.Float32:
		return "FLOAT"
	case reflect.Float64:
		return "DOUBLE"
	case reflect.String:
		return "VARCHAR"
	case reflect.Array, reflect.Slice:
		return "BLOB"
	case reflect.Struct:
		if _, ok := typ.Interface().(time.Time); ok {
			return "DATETIME"
		}
	}
	panic(fmt.Sprintf("invalid sql type(%s) %s", typ.Type().Name(), typ.Kind()))
}

func (m *mysql) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "select TABLE_NAME from INFORMATION_SCHEMA.TABLES where TABLE_NAME=?", args
}
