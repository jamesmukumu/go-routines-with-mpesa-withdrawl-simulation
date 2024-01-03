package customers

import "strings"

//create a class customer

type Customer struct{
Firstname string `json:"firstname" bson:"firstname"`
Secondname string `json:"secondname" bson:"secondname"`
Phonenumber string `json:"phone" bson:"phone"`
IdNumber  string `json:"idno" bson:"idno"`
Balance int `json:"bal" bson:"bal"`
Pin string `json:"pin" bson:"pin"`

}
func(customer *Customer)Ensurenoemptyfields()bool{
if customer.Pin=="" ||customer.Firstname=="" || customer.Secondname==""||customer.IdNumber=="" ||customer.Phonenumber==""||customer.Balance==0 {

	
	return false
}else{
	return true
}
	
}

//ensure phonenumber is ten characters long

func(customer *Customer)CheckPhonenumberLength()bool{
if len(customer.Phonenumber) != 12 {
	return false
}else{
	return true
}


}


//check if number starts with 254
func (customer *Customer)CheckPhoneformat()bool{
if strings.HasPrefix(customer.Phonenumber,"254") {
	return true
}else{
	return false
}

}

//check that pin number is four charachters
func(customer *Customer)ValidatePin()bool{
if len(customer.Pin) != 4 || customer.Pin==""{
	return false
}else{
	return true
}


}

