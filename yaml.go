package gluaparse

import (
	"github.com/mrnavastar/golua/lua"
	"gopkg.in/yaml.v3"
)

var YAML = map[string]lua.LuaGoFunction{
	"decode": decodeYAML,
}

func decodeYAML(l *lua.State) int {
	var v interface{}
	if err := yaml.Unmarshal([]byte(l.ToString(1)), &v); err != nil {
		l.PushNil()
		l.PushString(err.Error())
		return 2
	}
	PushGoInterface(l, v)
	return 1
}
