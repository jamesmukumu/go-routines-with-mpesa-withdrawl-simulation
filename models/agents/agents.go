package agents

import (


	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Agent struct {
	AgentNumber string `json:"agentNo" bson:"agentNo"`
	StoreNumber string `json:"storeNo" bson:"storeNo"`
	AgentName   string `json:"agentname" bson:"agentname"`
	Location    string `json:"location" bson:"location"`
	Owner      primitive.ObjectID `json:"id"`
}



//methods
//1 ensure non empty field during agent registration
//2 ensure that agentNumber and store Number are of ceratin lenght 
func(agent *Agent)Ensurenonemptyfields()bool{
if agent.AgentName==""||agent.AgentNumber==""||agent.Location=="" {
	return false
}else{  
	return true
}
}

func (agent *Agent)CheckStorelength()bool{
if len(agent.StoreNumber) != 5  || len(agent.AgentNumber) !=5{
	return false
}else{
	return true
}

} 