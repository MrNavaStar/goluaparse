package goluaparse

import (
	"encoding/xml"

	"github.com/aarzilli/golua/lua"
	"github.com/mrnavastar/lunatico"
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
	lunatico.PushAny(l, v)
	return 1
}
