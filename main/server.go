package main

import (  
    // Standard library packages
    "net/http"
	"log"
    // Third party packages
    "github.com/julienschmidt/httprouter"
    "gopkg.in/mgo.v2"
    "controllers"
)

func getSession() *mgo.Session {  
    // Connect to our local mongo
    s, err := mgo.Dial("mongodb://localhost")

    // Check if connection error, is mongo running?
    if err != nil {
        panic(err)
    }
    return s
}

func main() {  
    // Instantiate a new router
    r := httprouter.New()

    // Get a UserController instance
    uc := controllers.NewActorController(getSession())

    // Get a user resource
    r.GET("/actor/:id", uc.GetActor)

    r.POST("/actor", uc.CreateActor)

    r.DELETE("/actor/:id", uc.RemoveActor)

	log.Print("Server started on 3300")
    // Fire up the server
    http.ListenAndServe("localhost:3300", r)
}