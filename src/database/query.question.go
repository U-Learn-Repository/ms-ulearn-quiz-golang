package database

import (
	"fmt"
	"time"

	"github.com/U-Learn-Repository/ms-ulearn-quiz-golang/src/models"
	"gopkg.in/mgo.v2/bson"
)

// GetQuestions [question1, question2, ...]
func (mongo *MongoDB) GetQuestions() (questions []models.QuestionResponse, err error) {
	session := mongo.Session.Clone()
	defer session.Close()

	pipe := session.DB(mongo.Database).C(models.CollectionQuestion).Pipe([]bson.M{
		{
			"$lookup": bson.M{
				"from":         "answer",
				"localField":   "answers",
				"foreignField": "_id",
				"as":           "answers",
			},
		},
	})

	pipe.All(&questions)

	// err = session.DB(mongo.Database).C(models.CollectionQuestion).Find(bson.M{}).All(&questions)
	return questions, err
}

// https://stackoverflow.com/questions/46774424/mgo-aggregation-how-to-reuse-model-types-to-query-and-unmarshal-mixed-results/46774425

func (mongo *MongoDB) GetQuestionById(id string) (question models.QuestionResponse, err error) {
	session := mongo.Session.Clone()
	defer session.Close()

	/*
			  {
		        $match: {
		            _id: ObjectId("5ea8dc2b72db48f43f87065c")
		        }
		    },
		    {
		        "$lookup": {
		            "from": "answer",
		            "localField": "answers",
		            "foreignField": "_id",
		            "as": "answers"
		        }
		    }
	*/
	/*
		err = session.
			DB(mongo.Database).
			C(models.CollectionQuestion).
			FindId(bson.ObjectIdHex(id)).One(&question)*/

	pipe := session.DB(mongo.Database).C(models.CollectionQuestion).Pipe([]bson.M{
		{
			"$match": bson.M{
				"_id": bson.ObjectIdHex(id),
			},
		},
		{
			"$lookup": bson.M{
				"from":         "answer",
				"localField":   "answers",
				"foreignField": "_id",
				"as":           "answers",
			},
		},
	})

	pipe.One(&question)

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
