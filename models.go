package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type Users struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	DOB    string             `json:"dob" bson:"dob,omitempty"`
	Mobile string             `json:"mobile" bson:"mobile,omitempty"`
	Email  string 			  `json:"email" bson: "email,omitempty"`
	Timestamp string          `json:"timestamp" bson: "timestamp,omitempty"`
}

type Contact struct {
	UserIdOne string `json:"idone,omitempty" bson:"idone,omitempty"`
	UserIdTwo string `json:"idtwo,omitempty" bson:"idtwo,omitempty"`
	Timestamp string `json:"timestamp" bson: "timestamp,omitempty"`
}