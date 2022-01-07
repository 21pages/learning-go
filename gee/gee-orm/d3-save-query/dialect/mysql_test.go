package dialect

import (
	"reflect"
	"testing"
)

func TestDataTypeOfMysql(t *testing.T) {
	dial := &mysql{}
	cases := []struct {
		Value interface{}
		Type  string
	}{
		{"Tom", "VARCHAR"},
		{123, "BIGINT"},
		{1.2, "DOUBLE"},
		{[]int{1, 2, 3}, "BLOB"},
	}

	for _, c := range cases {
		if typ := dial.DataTypeOf(reflect.ValueOf(c.Value)); typ != c.Type {
			t.Fatalf("expect %s, but got %s", c.Type, typ)
		}
	}
}
