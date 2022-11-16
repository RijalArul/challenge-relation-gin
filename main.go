package main

import (
	"challenge-relation-gin/config"
	"challenge-relation-gin/routes"
)

func main() {
	config.StartDB()
	routes.Routes()
}
