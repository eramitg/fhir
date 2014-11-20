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

func QuestionnaireAnswersIndexHandler(rw http.ResponseWriter, r *http.Request) {
	var result []models.QuestionnaireAnswers
	c := Database.C("questionnaireanswerss")
	iter := c.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	var bundle models.QuestionnaireAnswersBundle
	bundle.Type = "Bundle"
	bundle.Title = "QuestionnaireAnswers Index"
	bundle.Id = bson.NewObjectId().Hex()
	bundle.Updated = time.Now()
	bundle.TotalResults = len(result)
	bundle.Entries = result

	log.Println("Setting questionnaireanswers search context")
	context.Set(r, "QuestionnaireAnswers", result)
	context.Set(r, "Resource", "QuestionnaireAnswers")
	context.Set(r, "Action", "search")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(bundle)
}

func QuestionnaireAnswersShowHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("questionnaireanswerss")

	result := models.QuestionnaireAnswers{}
	err := c.Find(bson.M{"_id": id.Hex()}).One(&result)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting questionnaireanswers read context")
	context.Set(r, "QuestionnaireAnswers", result)
	context.Set(r, "Resource", "QuestionnaireAnswers")
	context.Set(r, "Action", "read")

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(rw).Encode(result)
}

func QuestionnaireAnswersCreateHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	questionnaireanswers := &models.QuestionnaireAnswers{}
	err := decoder.Decode(questionnaireanswers)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("questionnaireanswerss")
	i := bson.NewObjectId()
	questionnaireanswers.Id = i.Hex()
	err = c.Insert(questionnaireanswers)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting questionnaireanswers create context")
	context.Set(r, "QuestionnaireAnswers", result)
	context.Set(r, "Resource", "QuestionnaireAnswers")
	context.Set(r, "Action", "create")

	host, err := os.Hostname()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Add("Location", "http://"+host+":8080/QuestionnaireAnswers/"+i.Hex())
}

func QuestionnaireAnswersUpdateHandler(rw http.ResponseWriter, r *http.Request) {

	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	decoder := json.NewDecoder(r.Body)
	questionnaireanswers := &models.QuestionnaireAnswers{}
	err := decoder.Decode(questionnaireanswers)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	c := Database.C("questionnaireanswerss")
	questionnaireanswers.Id = id.Hex()
	err = c.Update(bson.M{"_id": id.Hex()}, questionnaireanswers)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	log.Println("Setting questionnaireanswers update context")
	context.Set(r, "QuestionnaireAnswers", result)
	context.Set(r, "Resource", "QuestionnaireAnswers")
	context.Set(r, "Action", "update")
}

func QuestionnaireAnswersDeleteHandler(rw http.ResponseWriter, r *http.Request) {
	var id bson.ObjectId

	idString := mux.Vars(r)["id"]
	if bson.IsObjectIdHex(idString) {
		id = bson.ObjectIdHex(idString)
	} else {
		http.Error(rw, "Invalid id", http.StatusBadRequest)
	}

	c := Database.C("questionnaireanswerss")

	err := c.Remove(bson.M{"_id": id.Hex()})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Setting questionnaireanswers delete context")
	context.Set(r, "QuestionnaireAnswers", id.Hex())
	context.Set(r, "Resource", "QuestionnaireAnswers")
	context.Set(r, "Action", "delete")
}
