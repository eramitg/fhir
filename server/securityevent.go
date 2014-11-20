package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gitlab.mitre.org/intervention-engine/fhir/models"
	"gopkg.in/mgo.v2/bson"
)

func SecurityEventIndexHandler(rw http.ResponseWriter, r *http.Request) {
	var result []models.SecurityEvent
	c := Database.C("securityevents")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var bundle models.SecurityEventBundle
	bundle.Type = "Bundle"
	bundle.Title = "SecurityEvent Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entries = result

	log.Println("Setting securityevent search context")
	context.Set(r, "SecurityEvent", result)
	context.Set(r, "Resource", "SecurityEvent")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func SecurityEventShowHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("securityevents")

	result := models.SecurityEvent{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting securityevent read context")
	context.Set(r, "SecurityEvent", result)
	context.Set(r, "Resource", "SecurityEvent")
	context.Set(r, "Action", "read")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(result)
}

func SecurityEventCreateHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	securityevent := &models.SecurityEvent{}
	err := decoder.Decode(securityevent)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("securityevents")
	i := bson.NewObjectId()
	securityevent.Id = i.Hex()
	err = c.Insert(securityevent)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting securityevent create context")
	context.Set(r, "SecurityEvent", result)
	context.Set(r, "Resource", "SecurityEvent")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":8080/SecurityEvent/"+i.Hex())
}

func SecurityEventUpdateHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	securityevent := &models.SecurityEvent{}
	err := decoder.Decode(securityevent)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("securityevents")
	securityevent.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, securityevent)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting securityevent update context")
	context.Set(r, "SecurityEvent", result)
	context.Set(r, "Resource", "SecurityEvent")
	context.Set(r, "Action", "update")
}

func SecurityEventDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("securityevents")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting securityevent delete context")
	context.Set(r, "SecurityEvent", id.Hex())
	context.Set(r, "Resource", "SecurityEvent")
	context.Set(r, "Action", "delete")
}
