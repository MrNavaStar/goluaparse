package gluaparse

import (
	"encoding/xml"

	"github.com/mrnavastar/golua/lua"
)

var XML = map[string]lua.LuaGoFunction{
	"decode": decodeXML,
}

func decodeXML(l *lua.State) int {
	var v interface{}
	if err := xml.Unmarshal([]byte(l.ToString(1)), &v); err != nil {
		l.PushNil()
		l.PushString(err.Error())
		return 2
	}
	PushGoInterface(l, v)
	return 1
}
