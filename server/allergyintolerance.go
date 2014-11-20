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

func AllergyIntoleranceIndexHandler(rw http.ResponseWriter, r *http.Request) {
	var result []models.AllergyIntolerance
	c := Database.C("allergyintolerances")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var bundle models.AllergyIntoleranceBundle
	bundle.Type = "Bundle"
	bundle.Title = "AllergyIntolerance Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entries = result

	log.Println("Setting allergyintolerance search context")
	context.Set(r, "AllergyIntolerance", result)
	context.Set(r, "Resource", "AllergyIntolerance")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func AllergyIntoleranceShowHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("allergyintolerances")

	result := models.AllergyIntolerance{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting allergyintolerance read context")
	context.Set(r, "AllergyIntolerance", result)
	context.Set(r, "Resource", "AllergyIntolerance")
	context.Set(r, "Action", "read")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(result)
}

func AllergyIntoleranceCreateHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	allergyintolerance := &models.AllergyIntolerance{}
	err := decoder.Decode(allergyintolerance)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("allergyintolerances")
	i := bson.NewObjectId()
	allergyintolerance.Id = i.Hex()
	err = c.Insert(allergyintolerance)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting allergyintolerance create context")
	context.Set(r, "AllergyIntolerance", result)
	context.Set(r, "Resource", "AllergyIntolerance")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":8080/AllergyIntolerance/"+i.Hex())
}

func AllergyIntoleranceUpdateHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	allergyintolerance := &models.AllergyIntolerance{}
	err := decoder.Decode(allergyintolerance)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("allergyintolerances")
	allergyintolerance.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, allergyintolerance)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting allergyintolerance update context")
	context.Set(r, "AllergyIntolerance", result)
	context.Set(r, "Resource", "AllergyIntolerance")
	context.Set(r, "Action", "update")
}

func AllergyIntoleranceDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("allergyintolerances")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting allergyintolerance delete context")
	context.Set(r, "AllergyIntolerance", id.Hex())
	context.Set(r, "Resource", "AllergyIntolerance")
	context.Set(r, "Action", "delete")
}
