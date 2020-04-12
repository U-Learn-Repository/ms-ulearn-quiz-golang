package database

import (
	"fmt"
	"time"

	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/models"
	"gopkg.in/mgo.v2/bson"
)

func (mongo *MongoDB) InsertQualification(newQualification *models.Qualification) (err error) {
	session := mongo.Session.Clone()
	defer session.Clone()

	newQualification.CreateAt = time.Now()
	newQualification.UpdateAt = time.Now()

	err = session.
		DB(mongo.Database).
		C(models.CollectionQualification).
		Insert(&newQualification)

	return err
}

func (mongo *MongoDB) UpdateQualification(id string, qual models.Qualification) (qualification models.Qualification, err error) {
	session := mongo.Session.Clone()
	defer session.Close()

	err = session.
		DB(mongo.Database).
		C(models.CollectionQualification).
		UpdateId(bson.ObjectIdHex(id),
			bson.M{"$set": bson.M{"value": qual.Value}})

	session.
		DB(mongo.Database).
		C(models.CollectionQualification).
		FindId(bson.ObjectIdHex(id)).One(&qualification)

	return qualification, err
}

func (mongo *MongoDB) DeleteQualification(id string) (err error) {
	session := mongo.Session.Clone()
	defer session.Close()

	fmt.Println(id)

	err = session.
		DB(mongo.Database).
		C(models.CollectionQualification).
		Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	return err
}
