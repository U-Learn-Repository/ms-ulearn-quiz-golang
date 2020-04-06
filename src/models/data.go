package models

import "gopkg.in/mgo.v2/bson"
import "time"

const (
	CollectionQuestion = "question"
	CollectionAnswer   = "answer"
	CollectionQualification = "qualification"
)

type Question struct {
	Id        bson.ObjectId   `json:"_id,omitempty" bson:"_id, omitempty"`
	Statement string          `json:"statement,omitempty" form:"statement" binding:"require" bson:"statement,omitempty"`
	Score     int16           `json:"score,omitempty" form:"score" binding:"require" bson:"score,omitempty"`
	UserId    int64           `json:"user_id,omitempty" form:"user_id" binding:"require" bson:"user_id,omitempty"`
	CreateAt  time.Time       `json:"create_at,omitempty" form:"create_at" binding:"require" bson:"create_at,omitempty"`
	UpdateAt  time.Time       `json:"update_at,omitempty" form:"update_at" binding:"require" bson:"update_at,omitempty"`
	Answers   []bson.ObjectId `json:"answers" form:"answers" binding:"require" bson:"answers"`
	Qualification []bson.ObjectId `json:"qualification" form:"qualification" binding:"require" bson:"qualification"`
}

// Relationship in golang with mgo
// https://stackoverflow.com/questions/27659487/relationships-in-mgo/27660642

type Answer struct {
	Id        bson.ObjectId `json:"_id,omitempty" bson:"_id, omitempty"`
	Contest   string        `json:"contest,omitempty" form:"contest" binding:"require" bson:"contest,omitempty"`
	IsCorrect bool          `json:"is_correct,omitempty" form:"is_correct" binding:"require" bson:"is_correct,omitempty"`
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