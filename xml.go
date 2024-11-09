package goluaparse

import (
	"github.com/aarzilli/golua/lua"
	"github.com/clbanning/mxj"
	"github.com/stevedonovan/luar"
)

var XML = map[string]lua.LuaGoFunction{
	"decode": decodeXML,
}

func decodeXML(l *lua.State) int {
	result, err := mxj.NewMapXml([]byte(l.ToString(1)))
	if err != nil {
		l.PushNil()
		l.PushString(err.Error())
		return 2
	}

	luar.GoToLua(l, result)
	return 1
}
