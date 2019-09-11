package main

import (
	"v2/databases"
	"v2/routes"
)

func main() {
	e := routes.Init()
	databases.Init()
	e.Logger.Fatal(e.Start(":9000"))
}
