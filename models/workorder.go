package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Workorder struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	ObjNr       string        `json:"objNr" bson:"objNr"`
	Adress      string        `json:"address,omitempty" bson:"address,omitempty"`
	Description string        `json:"description,omitempty" bson:"description,omitempty"`
	Start       string        `json:"start,omitempty" bson:"start,omitempty"`
	Status      string        `json:"status,omitempty" bson:"status,omitempty"`
	Invoice     string        `json:"invoice,omitempty" bson:"invoice,omitempty"`
	Worker      []struct {
		Name        string  `json:"name,omitempty" bson:"name,omitempty"`
		HoursWorked float64 `json:"hoursWorked,omitempty" bson:"hoursWorked,omitempty"`
		Description string  `json:"description,omitempty" bson:"description,omitempty"`
	} `json:"worker,omitempty" bson:"worker,omitempty"`
}
