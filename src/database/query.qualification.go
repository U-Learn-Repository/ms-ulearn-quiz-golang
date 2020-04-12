package database

import (
	"time"

	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/models"
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
