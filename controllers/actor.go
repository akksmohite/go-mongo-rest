package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"models"
	"net/http"
)

type (
	
	ActorController struct {
		session *mgo.Session
	}
)

func NewActorController(s *mgo.Session) *ActorController {
	return &ActorController{s}
}

//localhost:3300/actor/5a79b44b7526232590ab80c6
func (uc ActorController) GetActor(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Stub user
	u := models.Actor{}

	// Fetch user
	if err := uc.session.DB("go_rest_tutorial").C("actors").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

//localhost:3300/actor {"name": "tom Smith", "gender": "male", "age": 50}
// CreateUser creates a new user resource
func (uc ActorController) CreateActor(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Stub an user to be populated from the body
	u := models.Actor{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)

	// Add an Id
	//u.Id = "foo"
	u.Id = bson.NewObjectId()

	uc.session.DB("go_rest_tutorial").C("actors").Insert(u)

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}


func (uc ActorController) RemoveActor(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // Grab id
    id := p.ByName("id")

    // Verify id is ObjectId, otherwise bail
    if !bson.IsObjectIdHex(id) {
        w.WriteHeader(404)
        return
    }

    // Grab id
    oid := bson.ObjectIdHex(id)

    // Remove user
    if err := uc.session.DB("go_rest_tutorial").C("actors").RemoveId(oid); err != nil {
        w.WriteHeader(404)
        return
    }

    // Write status
    w.WriteHeader(200)
}
