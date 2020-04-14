package main

import (
	"strconv"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "Hello World!"
}

func allContacts() []Contact {
	val, _ := svc.AllContacts()
	return val
}

func readContact(name string) Contact {
	val, _ := svc.ReadContact(name)
	return val
}

func removeContact(name string) {
	svc.RemoveContact(name)
}

func saveContact(input map[string]interface{}) {
	fn := input["Firstname"].(string)
	ln := input["Lastname"].(string)
	m, _ := strconv.Atoi(input["Mobile"].(string))
	h, _ := strconv.Atoi(input["Home"].(string))
	o, _ := strconv.Atoi(input["Office"].(string))
	contact := Contact{Firstname: fn,
		Lastname: ln,
		Mobile:   m,
		Home:     h,
		Office:   o,
	}

	// bolB, _ := json.Marshal(input)
	// log.log.Info(string(bolB))

	svc.SaveContact(contact)
}

var svc ContactsService
var log *loggingService

func main() {

	svc = NewContactService()
	log = &loggingService{}

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "First",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Bind(readContact)
	app.Bind(saveContact)
	app.Bind(removeContact)
	app.Bind(allContacts)
	app.Run()
}
