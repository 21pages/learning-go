package schema

import (
	"geeorm/dialect"
	"go/ast"
	"reflect"
)

// 字段
type Field struct {
	Name string
	Type string //数据库类型
	Tag  string
}

// 表
type Schema struct {
	Model      interface{}       //结构体
	Name       string            //表名
	Fields     []*Field          //字段集合
	FieldNames []string          //字段名集合
	fieldMap   map[string]*Field //字段map, 内部
}

func (s *Schema) GetField(name string) *Field {
	return s.fieldMap[name] //maybe nil
}

// 结构体各成员值, dest为pointer
func (s *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest)) //Indirect 如果指针转值, 如果值不变
	fieldValues := []interface{}{}
	for _, field := range s.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface()) //获取结构体相应成员名的值
	}
	return fieldValues
}

// 如果结构体实现该接口, 就使用TableName指定的表名
type ITableName interface {
	TableName() string
}

//将结构体转为表
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()
	var tableName string
	if t, ok := dest.(ITableName); ok {
		tableName = t.TableName()
	} else {
		tableName = modelType.Name()
	}
	schema := &Schema{
		Model:      dest,
		Name:       tableName,
		Fields:     []*Field{},
		FieldNames: []string{},
		fieldMap:   make(map[string]*Field),
	}
	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if !p.Anonymous && ast.IsExported(p.Name) { // 或者p.IsExported()
			field := &Field{
				Name: p.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))), //转换数据类型
			}
			if tag, ok := p.Tag.Lookup("geeorm"); ok { //eg `geeorm:"val1" key2:"val2"`
				field.Tag = tag
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, field.Name)
			schema.fieldMap[field.Name] = field
		}
	}

	return schema
}
