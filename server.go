package main

import (
	"github.com/rafimuhammad01/learn-go-echo-rest/db"
	"github.com/rafimuhammad01/learn-go-echo-rest/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
