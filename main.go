package main

import (
	"news/routes"
)

func main() {
	r := routes.Register()

	r.Run()
}
