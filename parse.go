package gluaparse

import (
	"reflect"

	"github.com/mrnavastar/golua/lua"
)

func toSlice(arg interface{}) (out []interface{}, ok bool) {
	slice := reflect.ValueOf(arg)
    if slice.Kind() != reflect.Slice {
		return 
	}
    c := slice.Len()
    out = make([]interface{}, c)
    for i := 0; i < c; i++ {
        out[i] = slice.Index(i).Interface()
    }
    return out, true
}

// Pushes a representation of the given interface to the lua stack
func PushGoInterface(l *lua.State, value interface{}) {
	slice, ok := toSlice(value)
	if ok {
		value = slice
	}

	switch converted := value.(type) {
	case bool:
		l.PushBoolean(converted)
	case int:
		l.PushInteger(int64(converted))
	case float64:
		l.PushNumber(converted)
	case string:
		l.PushString(converted)
 	case []interface{}:
		l.CreateTable(len(converted), 0)
		for _, item := range converted {
			PushGoInterface(l, item)
		}
	case map[string]interface{}:
		l.CreateTable(0, len(converted))
		for key, item := range converted {
			PushGoInterface(l, item)
			l.SetField(-2, key)
		}
	}
}
