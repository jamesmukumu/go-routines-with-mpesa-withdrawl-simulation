package router

import (
	"fmt"

	
	agentcont "mongoDB/controllers/agentCont"
	userscontroller "mongoDB/controllers/usersController"
	"mongoDB/middlewares/usersmiddleware"

	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func Server() {
	godotenv.Load()
	PORT := os.Getenv("port")

	Router := mux.NewRouter()

	Router.HandleFunc("/register/mpesa", userscontroller.RegisterMpesa).Methods("POST")
    Router.HandleFunc("/deposit/cash",usersmiddleware.ValidationPredeposit(userscontroller.Completedeposit)).Methods("GET")
	Router.HandleFunc("/withdraw/cash", usersmiddleware.ValidationPrewithdrawl(userscontroller.Completewithdrawl)).Methods("GET")
	Router.HandleFunc("/register/agent", agentcont.RegisterAgentnumber).Methods("POST")

	fmt.Printf("Server listening at %s for request", PORT)
	http.ListenAndServe(":"+PORT, Router)

}
