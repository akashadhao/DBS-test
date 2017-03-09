package main

import (
        "fmt"
		"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type restaurants struct {
        _id string
        name string
        foodtype string
        location string
        imgurl string
        mapurl string
}

func main() {
	
        session, err := mgo.Dial("localhost:27017")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("restMDB").C("restaurants")
        
        result := restaurants{}
        err = c.Find(bson.M{"name": "yelp1"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Phone:", result.name)
}
