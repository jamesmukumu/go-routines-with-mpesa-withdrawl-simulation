package main

import (


	"mongoDB/db"
	"mongoDB/db/dbagents"
	"mongoDB/router"
)

func main() {

	 dbagents.DBconnectionAgents()
	 db.DBconnection()
	 router.Server()
  
}



