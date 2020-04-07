package database

import "time"
import "gopkg.in/mgo.v2/bson"
import "github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/models"

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
