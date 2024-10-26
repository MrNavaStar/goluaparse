package gluaparse

import (
	"encoding/json"

	"github.com/yuin/gopher-lua"
)

func preload(l *lua.LState, name string, api map[string]lua.LGFunction) {
	l.PreloadModule(name, func(l *lua.LState) int {
		l.Push(l.SetFuncs(l.NewTable(), api))
		return 1
	})
}

func DecodeValue(l *lua.LState, value interface{}) lua.LValue {
	switch converted := value.(type) {
	case bool:
		return lua.LBool(converted)
	case float64:
		return lua.LNumber(converted)
	case string:
		return lua.LString(converted)
	case json.Number:
		return lua.LString(converted)
	case []interface{}:
		arr := l.CreateTable(len(converted), 0)
		for _, item := range converted {
			arr.Append(DecodeValue(l, item))
		}
		return arr
	case map[string]interface{}:
		tbl := l.CreateTable(0, len(converted))
		for key, item := range converted {
			tbl.RawSetH(lua.LString(key), DecodeValue(l, item))
		}
		return tbl
	case nil:
		return lua.LNil
	}
	return lua.LNil
}
