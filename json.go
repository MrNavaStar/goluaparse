package goluaparse

import (
	"encoding/json"

	"github.com/aarzilli/golua/lua"
	"github.com/mrnavastar/lunatico"
)

var JSON = map[string]lua.LuaGoFunction{
	"decode": decodeJSON,
}

func decodeJSON(l *lua.State) int {
	var v interface{}
	if err := json.Unmarshal([]byte(l.ToString(1)), &v); err != nil {
		l.PushNil()
		l.PushString(err.Error())
		return 2
	}
	lunatico.PushAny(l, v)
	return 1
}
