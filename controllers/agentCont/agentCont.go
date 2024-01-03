package agentcont

import (
	"context"
	"encoding/json"

	"mongoDB/db/dbagents"
	"mongoDB/models/agents"
	"net/http"
)

var Testagent agents.Agent

func RegisterAgentnumber(res http.ResponseWriter, req *http.Request) {

	decodedJson := json.NewDecoder(req.Body).Decode(&Testagent)
	if decodedJson != nil {
		res.Write([]byte(decodedJson.Error()))
		return
	}

	//validate methods
	if !Testagent.CheckStorelength() {
		res.Write([]byte("Store Number not long enough"))
		return
	}

	if !Testagent.Ensurenonemptyfields() {
		res.Write([]byte("No field should be empty"))
		return
	}

	insertedAgent, err := dbagents.Agents.InsertOne(context.Background(), Testagent)
	if err != nil {
		res.Write([]byte(err.Error()))
		return
	} else {
		json.NewEncoder(res).Encode(map[string]interface{}{"id": insertedAgent.InsertedID})
	}

}
