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

func MedicationDispenseIndexHandler(rw http.ResponseWriter, r *http.Request) {
	var result []models.MedicationDispense
	c := Database.C("medicationdispenses")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var bundle models.MedicationDispenseBundle
	bundle.Type = "Bundle"
	bundle.Title = "MedicationDispense Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entries = result

	log.Println("Setting medicationdispense search context")
	context.Set(r, "MedicationDispense", result)
	context.Set(r, "Resource", "MedicationDispense")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func MedicationDispenseShowHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("medicationdispenses")

	result := models.MedicationDispense{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting medicationdispense read context")
	context.Set(r, "MedicationDispense", result)
	context.Set(r, "Resource", "MedicationDispense")
	context.Set(r, "Action", "read")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(result)
}

func MedicationDispenseCreateHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	medicationdispense := &models.MedicationDispense{}
	err := decoder.Decode(medicationdispense)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("medicationdispenses")
	i := bson.NewObjectId()
	medicationdispense.Id = i.Hex()
	err = c.Insert(medicationdispense)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting medicationdispense create context")
	context.Set(r, "MedicationDispense", result)
	context.Set(r, "Resource", "MedicationDispense")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":8080/MedicationDispense/"+i.Hex())
}

func MedicationDispenseUpdateHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	medicationdispense := &models.MedicationDispense{}
	err := decoder.Decode(medicationdispense)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("medicationdispenses")
	medicationdispense.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, medicationdispense)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting medicationdispense update context")
	context.Set(r, "MedicationDispense", result)
	context.Set(r, "Resource", "MedicationDispense")
	context.Set(r, "Action", "update")
}

func MedicationDispenseDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("medicationdispenses")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting medicationdispense delete context")
	context.Set(r, "MedicationDispense", id.Hex())
	context.Set(r, "Resource", "MedicationDispense")
	context.Set(r, "Action", "delete")
}
