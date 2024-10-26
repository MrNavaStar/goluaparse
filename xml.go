package gluaparse

import (
	"encoding/xml"

	"github.com/yuin/gopher-lua"
)

var api = map[string]lua.LGFunction{
	"encode": encode,
	"decode": decode,
}

func PreloadXML(l *lua.LState) {
	l.PreloadModule("xml", func(l *lua.LState) int {
		t := l.NewTable()
		l.SetFuncs(t, api)
		l.Push(t)
		return 1
	})
}

func encode(l *lua.LState) int {

	return 1
}

func decode(l *lua.LState) int {
	var v interface{}
	if err := xml.Unmarshal([]byte(l.ToString(1)), &v); err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}
	l.Push(decodeValue(l, v))
	return 1
}
