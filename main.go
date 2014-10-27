/*
usgae usage ./main -H=name-of-mongodb-host
TODO:
put in  some comments
Write tests
Phase 2: make output pretty and use template/http
Change mongo connection so uses envirnment variables so can pass to Container
*/
package main

import (
	"flag"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

var H string // mongodb hostname

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

	// flag( name, default value, description)
	flag.StringVar(&H, "H", "localhost", "MongoDB hostname")
	flag.Parse()

	http.HandleFunc("/", myHandler)
	log.Println("Listening on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func myHandler(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial(H)
	defer session.Close()

	if err != nil {
		fmt.Fprintf(w, "could not connect to mongodb")
		log.Fatal(err)
	}

	var elements []element

	session.DB("demo").C("periodictable").Find(nil).All(&elements)

	for _, element := range elements {
		fmt.Fprintf(w, "Element %s has name  %s  and atomic number  %s \n", element.Symbol, element.Name, element.Z)
	}
}
