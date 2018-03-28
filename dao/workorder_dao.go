package dao

import (
	. "github.com/marante/bravis-app/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type WorkorderDao struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	WorkorderCollection = "workorders"
)

func (w *WorkorderDao) Connect() {
	session, err := mgo.Dial(w.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(w.Database)
}

func (w *WorkorderDao) FindAll() ([]Workorder, error) {
	var workorders []Workorder
	err := db.C(WorkorderCollection).Find(bson.M{}).All(&workorders)
	return workorders, err
}
