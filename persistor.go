package main

import (
	"gopkg.in/mgo.v2"
)

func getCollection() (* mgo.Collection){
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	connection := session.DB("gogreen").C("measurements")
	return connection
}

func StoreMeasurement(measurement Measurement) {
	c := getCollection()
	err := c.Insert(measurement)
	if err != nil {
		panic(err)
	}
}

func FindMeasurements() ([]Measurement){
	c := getCollection()
	var results []Measurement
	err := c.Find(nil).All(&results)

	if err != nil {
		panic(err)
	}

	return results
}

func FindMeasurement(measurement Measurement) (Measurement){
	c := getCollection();
	result := Measurement{}
	err := c.Find(measurement).One(&result)

	if err != nil {
		panic(err)
	}

	return result
}

func DeleteMeasurements(){
	getCollection().RemoveAll(nil)
}
