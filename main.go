package main

import (
	"hello-world/backend/config"
	"hello-world/backend/routes"
)

func main() {
	config.InitConfig()
	routes.Home()

	config.App.Listen(":3000")

}
