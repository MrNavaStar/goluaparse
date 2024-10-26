package gluaparse

import (
	"github.com/yuin/gopher-lua"
	"gopkg.in/yaml.v3"
)

var yamlApi = map[string]lua.LGFunction{
	"decode": decodeYAML,
}

func PreloadYAML(l *lua.LState) {
	preload(l, "yaml", yamlApi)
}

func decodeYAML(l *lua.LState) int {
	var v interface{}
	if err := yaml.Unmarshal([]byte(l.ToString(1)), &v); err != nil {
		l.Push(lua.LNil)
		l.Push(lua.LString(err.Error()))
		return 2
	}
	l.Push(decodeValue(l, v))
	return 1
}
