package main

import (
	"github.com/mrnavastar/gluaparse"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	z := gluaparse.DecodeValue(lua.NewState(), nil)
	print(z.Type().String())
}
