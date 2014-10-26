/*
TODO:
Put in some messagsing when http service starts
put in  some comments
Write tests
Phase 2: make pretty by using one of the framework libraries
*/
package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type element struct {
	ID                     bson.ObjectId `bson:"_id,omitempty"`
	Z                      string        `bson:"Z"`
	Symbol                 string        `bson:"Symbol"`
	Name                   string        `bson:"Name"`
	A                      string        `bson:"A"`
	N                      string        `bson:"N"`
	Period                 string        `bson:"Period"`
	IUPAC_Group            string        `bson:"IUPAC_Group"`
	Old_IUPAC              string        `bson:"Old_IUPAC"`
	CAS_Group              string        `bson:"CAS_Group"`
	Category               string        `bson:"Category"`
	Standard_Atomic_Weight string        `bson:"Standard_Atomic_Weight"`
	Last_Digit             string        `bson:"Last_Digit"`
	Ionisation_eV          string        `bson:"Ionisation_eV"`
	Normal_State           string        `bson:"Normal_State"`
	Density_Kg_m3_20C      string        `bson:"Density_Kg_m3_20C"`
	Melting_point_C        string        `bson:"Melting_point_C"`
	Melting_point_K        string        `bson:"Melting_point_K"`
	Boiling_point_C        string        `bson:"Boiling_point_C"`
	Boiling_point_K        string        `bson:"Boiling_point_K"`
	Year_of_discovery      string        `bson:"Year_of_discovery"`
	Discoverer             string        `bson:"Discoverer"`
}

func main() {

	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("localhost")
	defer session.Close()

	if err != nil {
		fmt.Fprintf(w, "could not connect to mongodb")
	}

	var elements []element

	session.DB("demo").C("periodictable").Find(nil).All(&elements)

	for _, element := range elements {
		fmt.Fprintf(w, "Element %s has name  %s  and atomic number  %s \n", element.Symbol, element.Name, element.Z)
	}
}
