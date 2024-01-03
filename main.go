package main

import (
	helpers "mongoDB/Helpers"
	"mongoDB/db"
	"mongoDB/db/dbagents"
	"mongoDB/router"
)

func main() {
	helpers.Wg.Add(3)
	go dbagents.DBconnectionAgents()
	go db.DBconnection()
	go router.Server()
	helpers.Wg.Wait()
}
