package main

import (
	"github.com/isaque-vieira2019/challenge-bravo/database"
	"github.com/isaque-vieira2019/challenge-bravo/routes"
)

func main() {
	database.DbConnect()
	routes.HandleRequest()
}
