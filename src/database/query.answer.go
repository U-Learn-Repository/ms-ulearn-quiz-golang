package database

import "github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/models"

func (mongo *MongoDB) InsertAnswer(newAnswer *models.Answer) (err error) {
	session := mongo.Session.Clone()
	defer session.Clone()

	err = session.DB(mongo.Database).C(models.CollectionAnswer).Insert(&newAnswer)

	return err
}