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

func CompositionIndexHandler(rw http.ResponseWriter, r *http.Request) {
	var result []models.Composition
	c := Database.C("compositions")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var bundle models.CompositionBundle
	bundle.Type = "Bundle"
	bundle.Title = "Composition Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entries = result

	log.Println("Setting composition search context")
	context.Set(r, "Composition", result)
	context.Set(r, "Resource", "Composition")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func CompositionShowHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("compositions")

	result := models.Composition{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting composition read context")
	context.Set(r, "Composition", result)
	context.Set(r, "Resource", "Composition")
	context.Set(r, "Action", "read")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(result)
}

func CompositionCreateHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	composition := &models.Composition{}
	err := decoder.Decode(composition)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("compositions")
	i := bson.NewObjectId()
	composition.Id = i.Hex()
	err = c.Insert(composition)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting composition create context")
	context.Set(r, "Composition", result)
	context.Set(r, "Resource", "Composition")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":8080/Composition/"+i.Hex())
}

func CompositionUpdateHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	composition := &models.Composition{}
	err := decoder.Decode(composition)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("compositions")
	composition.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, composition)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting composition update context")
	context.Set(r, "Composition", result)
	context.Set(r, "Resource", "Composition")
	context.Set(r, "Action", "update")
}

func CompositionDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("compositions")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting composition delete context")
	context.Set(r, "Composition", id.Hex())
	context.Set(r, "Resource", "Composition")
	context.Set(r, "Action", "delete")
}
