package gluaparse

import (
	"encoding/json"

	"github.com/yuin/gopher-lua"
)

var jsonApi = map[string]lua.LGFunction{
	"decode": decodeJSON,
}

func PreloadJSON(l *lua.LState) {
	preload(l, "json", jsonApi)
}

func decodeJSON(l *lua.LState) int {
	var v interface{}
	if err := json.Unmarshal([]byte(l.ToString(1)), &v); err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}
	l.Push(DecodeValue(l, v))
	return 1
}
