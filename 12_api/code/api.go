package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"

// Planet is a planet type
type Planet struct {
	Name       string `json:"name"`
	Terrain    string `json:"population"`
	Population string `json:"terrain"`
}

// Person is a person type
type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}

// AllPeople is a collection of Person types
type AllPeople struct {
	People []Person `json:"results"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome home!")
}

// Calls to get data for home world
func (p *Person) getHomeworld() {
	res, err := http.Get(p.HomeworldURL)
	if err != nil {
		log.Print("Error requesting home world", err)
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		log.Print("Error reading homeworld response body", err)
	}

	json.Unmarshal(bytes, &p.Homeworld)
}

// Calls to get the people data from the external API
func getPeople(w http.ResponseWriter, r *http.Request) {

	//Grab the data from the API
	res, err := http.Get(BaseURL + "people")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request people data")
	}
	fmt.Println(res)

	//Parse the data from the body of the request
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse request body")
	}

	var people AllPeople
	fmt.Println(string(bytes))

	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	fmt.Println(people)

	//Do the secondary request to get the homeworld info
	for _, pers := range people.People {
		pers.getHomeworld()
		fmt.Println(pers)
	}

}

func main() {
	fmt.Println(BaseURL)
	http.HandleFunc("/", home)
	//create a route for people
	http.HandleFunc("/people/", getPeople)
	fmt.Println("Server is running on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
