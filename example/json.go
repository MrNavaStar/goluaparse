package main

import (
	"log"

	"github.com/aarzilli/golua/lua"
	"github.com/mrnavastar/goluaparse"
)

const json =	`
	{
		"note": {
			"to": "Trove",
			"from": "Jani",
			"heading": "Reminder",
			"body": "Don't forget me this weekend!"
		}
	}
	`

const jsonCode = `
	local json = require("json")

	local parsed, err = json.decode(string_json)
	if err then
		error(err)
	end

	print(parsed["note"]["body"])
	`

func main() {
	L := lua.NewState()
	defer L.Close()
	L.OpenLibs()

	L.RegisterLibrary("json", goluaparse.JSON)

	L.PushString(json)
	L.SetGlobal("string_json")

	if err := L.DoString(jsonCode); err != nil {
		log.Fatal(err)
	}
}
