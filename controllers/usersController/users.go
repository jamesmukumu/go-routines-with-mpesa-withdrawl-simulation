package userscontroller

import (
	"context"
	"encoding/json"
	"fmt"

	"log"
	"mongoDB/db"
	"mongoDB/models/customers"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

//withdraw
type Withdrawamount struct{
	Amount int `json:"amount"`
	}
	
	//an instance of the amount
	var Transactamount Withdrawamount












var Customer customers.Customer

//register a new phone

func RegisterMpesa(res http.ResponseWriter,req *http.Request){
//decode json
decodedJson := json.NewDecoder(req.Body).Decode(&Customer)
if decodedJson != nil {
	log.Fatal(decodedJson.Error())
	return
}



if !Customer.CheckPhoneformat() {
res.Write([]byte("Wrong phone format"))
return
}
if !Customer.CheckPhonenumberLength() {
	res.Write([]byte("Phone number too short"))
	return
}
if !Customer.Ensurenoemptyfields() {
	res.Write([]byte("All field should be filled"))
	return
}

if !Customer.ValidatePin() {
	res.Write([]byte("Pin must be 4 charcahters long"))
	return
}





hashedPinnumber,err := bcrypt.GenerateFromPassword([]byte(Customer.Pin),20)
Customer.Pin = string(hashedPinnumber)
if err != nil {
	res.Write([]byte(err.Error()))
	return
}
insertedData,errr := db.Users.InsertOne(context.Background(),Customer)
 if errr != nil {
	res.Write([]byte(errr.Error()))
	return
 }
 


 if strings.Contains(err.Error(),"E11000 duplicate key error collection") {
	res.Write([]byte("PhoneNumber or Idnumber in use"))
	return
 }


 data := map[string]interface{}{"id":insertedData.InsertedID}
 jsonData,_ := json.Marshal(data)
 res.Write([]byte(jsonData))
 

}






//complete withdrawl
func Completewithdrawl(res http.ResponseWriter,req *http.Request){
	 var Mycustomer customers.Customer
// decodedJson := json.NewDecoder(req.Body).Decode(&Mycustomer)
// if decodedJson != nil {
// 	log.Fatal(decodedJson.Error())
// 	return
// }


	phoneNumberquery := req.URL.Query().Get("phone")
	// if !strings.HasPrefix("254",phoneNumberquery) {
	// 	res.Write([]byte("Wrong Phone Formst "))
	// 	return
	// }
	if phoneNumberquery == "" {
		res.Write([]byte("Please Provide a phoneNumber query"))
		return
	}
	
	actualphoneFilter :=  bson.M{
		"phone":phoneNumberquery,
	}
 

foundCustomer := db.Users.FindOne(context.Background(),actualphoneFilter)
decodedResult := foundCustomer.Decode(&Mycustomer)
if decodedResult != nil {
	res.Write([]byte(decodedResult.Error()))
	return
}


jsonResults,_ := json.Marshal(&Mycustomer)
fmt.Println(jsonResults)

 
Transactamount.Amount = 6
if Mycustomer.Balance < Transactamount.Amount {
	fmt.Print("Insufficient funds")
	return
}else{
//request pin
var Pinprompt customers.Customer

jsonDecodedpin := json.NewDecoder(req.Body).Decode(&Pinprompt)
if jsonDecodedpin != nil {
	log.Fatal(jsonDecodedpin.Error())

}
 

if !Pinprompt.ValidatePin() {
	res.Write([]byte("please Provide pin and it should be 4 charachters long"))
	return
}

//validate pin
matchingPin := bcrypt.CompareHashAndPassword([]byte(Mycustomer.Pin),[]byte(Pinprompt.Pin))
if matchingPin!= nil {
	res.Write([]byte("Wrong pin"))
	return
}

 


	newAmount := Mycustomer.Balance - Transactamount.Amount
	fmt.Printf("Withdraw sucess your new balance is %d",newAmount)


	update:= bson.M{
		"$set":bson.M{"bal":newAmount},
	}
	_,err2 := db.Users.UpdateOne(context.Background(),actualphoneFilter,update)
	if err2 != nil {
		res.Write([]byte(err2.Error()))
		return
	}else{
		http.Error(res,"New balance updated",http.StatusAccepted)
		return
	}
}


}