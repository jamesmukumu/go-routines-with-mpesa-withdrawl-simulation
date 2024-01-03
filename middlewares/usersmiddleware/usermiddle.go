package usersmiddleware

import (
	"context"
	"mongoDB/db"
	"mongoDB/db/dbagents"
	"mongoDB/models/agents"
	"mongoDB/models/customers"
	"net/http"


	"go.mongodb.org/mongo-driver/bson"
)







//middlware before actual Withrawal to handle
var Customer customers.Customer

var Testagent agents.Agent
func ValidationPrewithdrawl(next http.HandlerFunc)http.HandlerFunc{
return func(res http.ResponseWriter, req *http.Request) {
//validate First that ephoneNumber exists
var AccountDb customers.Customer
phoneNumberquery := req.URL.Query().Get("phone")
// if !strings.HasPrefix("254",phoneNumberquery) {
// 	res.Write([]byte("Wrong Phone Format "))
// 	return
// }
if phoneNumberquery == "" {
	res.Write([]byte("Please Provide a phoneNumber query"))
	return
}


actualphoneFilter := bson.M{
	"phone":phoneNumberquery,
}

foundPhonumber := db.Users.FindOne(context.Background(),actualphoneFilter).Decode(&AccountDb)
if foundPhonumber != nil {
	http.Error(res,foundPhonumber.Error(),http.StatusAccepted)
	return
}



//validate Agentnumber and store Number

agentNoquery := req.URL.Query().Get("agentNo")
storeNoquery := req.URL.Query().Get("storeNo")

if agentNoquery== "" || storeNoquery=="" {
	http.Error(res,"Please provide both store number and agent number",http.StatusAccepted)
	return
}
if len(agentNoquery) != 5 || len(storeNoquery) != 5 {
	http.Error(res,"Storenumber and agent number format wrong",http.StatusAccepted)
	return
}

actualMpesafilter := bson.M{
	"storeNo":storeNoquery,
	"agentNo":agentNoquery, 
}
result:= dbagents.Agents.FindOne(context.Background(),actualMpesafilter).Decode(&Testagent)
if result != nil {
	res.Write([]byte(result.Error()))
	return
}else{

	del := req.URL.Query()
     del.Del("agentNo")
	 del.Del("storeNo")
	 req.URL.RawQuery = del.Encode()

	next.ServeHTTP(res,req)
}




}



}