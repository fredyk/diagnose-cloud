package entities

import "time"

type Diagnose struct {
	Id string `json:"id,omitempty" bson:"_id,omitempty"`

	// required
	PatientId string    `json:"patientId" bson:"patientId"`
	Date      time.Time `json:"date" bson:"date"`
	Diagnose  string    `json:"diagnose" bson:"diagnose"`

	// optional
	Prescription string `json:"prescription,omitempty" bson:"prescription,omitempty"`

	// westack defaults
	Created  time.Time `json:"__created,omitempty" bson:"created,omitempty"`
	Modified time.Time `json:"__modified,omitempty" bson:"modified,omitempty"`
}
