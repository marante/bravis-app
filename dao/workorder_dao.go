package dao

import (
	"github.com/marante/bravis-app/models"
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

// Connect establishes a connection with the database.
func (w *WorkorderDao) Connect() {
	session, err := mgo.Dial(w.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(w.Database)
}

// Insert inserts a workorder document in the database
func (w *WorkorderDao) Insert(order models.Workorder) error {
	err := db.C(WorkorderCollection).Insert(&order)
	return err
}

// FindAll finds all entries in the collection.
func (w *WorkorderDao) FindAll() ([]models.Workorder, error) {
	var workorders []models.Workorder
	err := db.C(WorkorderCollection).Find(bson.M{}).All(&workorders)
	return workorders, err
}

// FindById finds an entry in the collection by id.
func (w *WorkorderDao) FindById(objNr string) (models.Workorder, error) {
	var order models.Workorder
	err := db.C(WorkorderCollection).Find(bson.M{"objNr": objNr}).One(&order)
	return order, err
}

// Update updates an entry in the db.
func (w *WorkorderDao) Update(order models.Workorder) error {
	err := db.C(WorkorderCollection).Update(order.ObjNr, &order)
	return err
}

// Delete deletes the specified entry in the db.
func (w *WorkorderDao) Delete(order models.Workorder) error {
	err := db.C(WorkorderCollection).Remove(&order)
	return err
}
