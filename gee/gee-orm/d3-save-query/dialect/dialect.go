package dialect

import (
	"geeorm/log"
	"reflect"
)

var dialectsMap = map[string]Dialect{} //不同的数据库类型

type Dialect interface {
	DataTypeOf(typ reflect.Value) string                    //数据类型映射
	TableExistSQL(tableName string) (string, []interface{}) //返回查询表是否存在的语句
}

func RegisterDialect(driver string, dialect Dialect) {
	dialectsMap[driver] = dialect
}

func GetDialect(driver string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[driver]
	if !ok {
		log.Error("GetDialect failed!", driver)
	}
	return
}
