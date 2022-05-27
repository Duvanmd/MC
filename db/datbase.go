package db

import (
	"adn/models"
	"fmt"

	"gopkg.in/mgo.v2"
)

var collection = getSession().DB("Adn").C("adn-mutan")

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

func SaveAdn(adn *models.Adn) bool {
	err := collection.Insert(adn)
	if err != nil {
		panic(err)
	} else {
		return false
	}
}

func GetAdn() []models.Adn {
	var results []models.Adn
	err := collection.Find(nil).Sort("_id").All(&results)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(results)
		return results
	}
}
