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
	switch converted := value.(type) {
	case bool:
		l.PushBoolean(converted)
		return
	case int:
		l.PushInteger(int64(converted))
		return
	case float64:
		l.PushNumber(converted)
		return
	case string:
		l.PushString(converted)
		return
	case []byte:
		l.PushBytes(converted)
		return
	}

	slice, ok := toSlice(value)
	if ok {
		value = slice
	}

	switch converted := value.(type) {
 	case []interface{}:
		l.CreateTable(len(converted), 0)
		for i, item := range converted {
			PushGoInterface(l, item)
			l.RawSeti(-2, i+1)
		}
	case map[string]interface{}:
		l.CreateTable(0, len(converted))
		for key, item := range converted {
			PushGoInterface(l, item)
			l.SetField(-2, key)
		}
	}
}
