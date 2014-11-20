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

func DocumentReferenceIndexHandler(rw http.ResponseWriter, r *http.Request) {
	var result []models.DocumentReference
	c := Database.C("documentreferences")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var bundle models.DocumentReferenceBundle
	bundle.Type = "Bundle"
	bundle.Title = "DocumentReference Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entries = result

	log.Println("Setting documentreference search context")
	context.Set(r, "DocumentReference", result)
	context.Set(r, "Resource", "DocumentReference")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func DocumentReferenceShowHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("documentreferences")

	result := models.DocumentReference{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting documentreference read context")
	context.Set(r, "DocumentReference", result)
	context.Set(r, "Resource", "DocumentReference")
	context.Set(r, "Action", "read")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(result)
}

func DocumentReferenceCreateHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	documentreference := &models.DocumentReference{}
	err := decoder.Decode(documentreference)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("documentreferences")
	i := bson.NewObjectId()
	documentreference.Id = i.Hex()
	err = c.Insert(documentreference)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting documentreference create context")
	context.Set(r, "DocumentReference", result)
	context.Set(r, "Resource", "DocumentReference")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":8080/DocumentReference/"+i.Hex())
}

func DocumentReferenceUpdateHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	documentreference := &models.DocumentReference{}
	err := decoder.Decode(documentreference)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("documentreferences")
	documentreference.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, documentreference)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting documentreference update context")
	context.Set(r, "DocumentReference", result)
	context.Set(r, "Resource", "DocumentReference")
	context.Set(r, "Action", "update")
}

func DocumentReferenceDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("documentreferences")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting documentreference delete context")
	context.Set(r, "DocumentReference", id.Hex())
	context.Set(r, "Resource", "DocumentReference")
	context.Set(r, "Action", "delete")
}
