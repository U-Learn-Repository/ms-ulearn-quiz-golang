package database

import (
	"fmt"
	"time"

	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/models"
	"gopkg.in/mgo.v2/bson"
)

// GetQuestions [question1, question2, ...]
func (mongo *MongoDB) GetQuestions() (questions []models.Question, err error) {
	session := mongo.Session.Clone()
	defer session.Close()

	err = session.DB(mongo.Database).C(models.CollectionQuestion).Find(bson.M{}).All(&questions)
	return questions, err
}

func (mongo *MongoDB) GetQuestionById(id string) (question models.Question, err error) {
	session := mongo.Session.Clone()
	defer session.Close()

	err = session.
		DB(mongo.Database).
		C(models.CollectionQuestion).
		FindId(bson.ObjectIdHex(id)).One(&question)

	return question, err
}

// InsertQuestion
func (mongo *MongoDB) InsertQuestion(newQuestion *models.Question) (err error) {
	session := mongo.Session.Clone()
	defer session.Close()

	newQuestion.CreateAt = time.Now()
	newQuestion.UpdateAt = time.Now()

	err = session.DB(mongo.Database).C(models.CollectionQuestion).Insert(&newQuestion)
	return err
}

func (mongo *MongoDB) UpdateQuestion(id string, quest models.Question) (question models.Question, err error) {
	session := mongo.Session.Clone()
	defer session.Close()

	fmt.Println(id)

	err = session.
		DB(mongo.Database).
		C(models.CollectionQuestion).
		UpdateId(bson.ObjectIdHex(id),
			bson.M{"$set": bson.M{"statement": quest.Statement, "score": quest.Score, "update_at": time.Now()}})

	session.
		DB(mongo.Database).
		C(models.CollectionQuestion).
		FindId(bson.ObjectIdHex(id)).One(&question)

	return question, err
}

func (mongo *MongoDB) DeleteQuestion(id string) (err error) {
	session := mongo.Session.Clone()
	defer session.Close()

	fmt.Println(id)

	err = session.
		DB(mongo.Database).
		C(models.CollectionQuestion).
		Remove(bson.M{"_id": bson.ObjectIdHex(id)})

	return err
}
