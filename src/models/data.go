package models

import "gopkg.in/mgo.v2/bson"
import "time"

const (
	CollectionQuestion = "question"
	CollectionAnswer   = "answer"
	CollectionQualification = "qualification"
)

type Question struct {
	Id        bson.ObjectId   `json:"_id,omitempty" bson:"_id,omitempty"`
	Statement string          `json:"statement,omitempty" bson:"statement,omitempty"`
	Score     int16           `json:"score,omitempty" bson:"score,omitempty"`
	UserId    int64           `json:"user_id,omitempty" bson:"user_id,omitempty"`
	CreateAt  time.Time       `json:"create_at" bson:"create_at"`
	UpdateAt  time.Time       `json:"update_at" bson:"update_at"`
	Answers   []bson.ObjectId `json:"answers" bson:"answers"`
	Qualifications []bson.ObjectId `json:"qualifications" bson:"qualifications"`
}

// Relationship in golang with mgo
// https://stackoverflow.com/questions/27659487/relationships-in-mgo/27660642

type Answer struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Context   string        `json:"context,omitempty" bson:"context,omitempty"`
	IsCorrect bool          `json:"is_correct,omitempty" bson:"is_correct,omitempty"`
}

type Qualification struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id, omitempty"`
	Value     int16         `json:"value,omitempty" form:"value" binding:"require" bson:"value,omitempty"`
	UserId    int64         `json:"user_id,omitempty" form:"user_id" binding:"require" bson:"user_id,omitempty"`
	CreateAt  time.Time     `json:"create_at,omitempty" form:"create_at" binding:"require" bson:"create_at,omitempty"`
	UpdateAt  time.Time     `json:"update_at,omitempty" form:"update_at" binding:"require" bson:"update_at,omitempty"`
}

// Default Data
// https://stackoverflow.com/questions/41907619/set-default-date-when-inserting-document-with-time-time-field