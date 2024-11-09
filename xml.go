package goluaparse

import (
	"strings"

	"github.com/aarzilli/golua/lua"
	"github.com/sbabiv/xml2map"
	"github.com/stevedonovan/luar"
)

var XML = map[string]lua.LuaGoFunction{
	"decode": decodeXML,
}

func decodeXML(l *lua.State) int {
	decoder := xml2map.NewDecoder(strings.NewReader(l.ToString(1)))
	result, err := decoder.Decode()
	if err != nil {
		l.PushNil()
		l.PushString(err.Error())
		return 2
	}

	luar.GoToLua(l, result)
	return 1
}
