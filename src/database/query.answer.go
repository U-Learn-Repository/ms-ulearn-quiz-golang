package database

import (
	"fmt"

	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/models"
	"gopkg.in/mgo.v2/bson"
)

func (mongo *MongoDB) InsertAnswer(newAnswer *models.Answer) (err error) {
	session := mongo.Session.Clone()
	defer session.Clone()

	err = session.DB(mongo.Database).C(models.CollectionAnswer).Insert(&newAnswer)

	return err
}

func (mongo *MongoDB) UpdateAnswer(id string, ans models.Answer) (answer models.Answer, err error) {
	session := mongo.Session.Clone()
	defer session.Close()

	fmt.Println(id)

	err = session.
		DB(mongo.Database).
		C(models.CollectionAnswer).
		UpdateId(bson.ObjectIdHex(id),
			bson.M{"$set": bson.M{"context": ans.Context, "is_correct": ans.IsCorrect}})

	session.
		DB(mongo.Database).
		C(models.CollectionAnswer).
		FindId(bson.ObjectIdHex(id)).One(&answer)

	return answer, err
}
