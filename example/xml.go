package main

import (
	"log"

	"github.com/aarzilli/golua/lua"
	"github.com/mrnavastar/goluaparse"
)

const xml =	`
	<note>
		<to>Tove</to>
		<from>Jani</from>
		<heading>Reminder</heading>
		<body>Don't forget me this weekend!</body>
	</note>
	`

const code = `
	local xml = require("xml")

	local parsed, err = xml.decode(string_xml)
	if err then
		error(err)
	end

	print(parsed["note"]["body"])
	`

func main() {
	L := lua.NewState()
	defer L.Close()
	L.OpenLibs()

	L.RegisterLibrary("xml", goluaparse.XML)

	L.PushString(xml)
	L.SetGlobal("string_xml")

	if err := L.DoString(code); err != nil {
		log.Fatal(err)
	}
}
