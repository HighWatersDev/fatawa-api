package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Fatwa struct {
	Title     string    `json:"title,omitempty" bson:"title"`
	Question  string    `json:"question" bson:"question"`
	Answer    string    `json:"answer" bson:"answer"`
	Link      string    `json:"link,omitempty" bson:"link,omitempty"`
	Author    string    `json:"author" bson:"author"`
	Topic     string    `json:"topic" bson:"topic"`
	Lang      string    `json:"lang" bson:"lang"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type FatwaDb struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title"`
	Question  string             `json:"question" bson:"question"`
	Answer    string             `json:"answer" bson:"answer"`
	Link      string             `json:"link,omitempty" bson:"link,omitempty"`
	Author    string             `json:"author" bson:"author"`
	Topic     string             `json:"topic" bson:"topic"`
	Lang      string             `json:"lang" bson:"lang"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
