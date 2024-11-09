package goluaparse

import (
	"github.com/aarzilli/golua/lua"
	"github.com/stevedonovan/luar"
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
	luar.GoToLua(l, v)
	return 1
}
