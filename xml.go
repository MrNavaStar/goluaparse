package gluaparse

import (
	"encoding/xml"

	"github.com/yuin/gopher-lua"
)

var xmlApi = map[string]lua.LGFunction{
	"decode": decodeXML,
}

func PreloadXML(l *lua.LState) {
	preload(l, "xml", xmlApi)
}

func decodeXML(l *lua.LState) int {
	var v interface{}
	if err := xml.Unmarshal([]byte(l.ToString(1)), &v); err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}
	l.Push(DecodeValue(l, v))
	return 1
}
