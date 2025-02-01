package entities

import "time"

type Diagnose struct {
	Id string `json:"id,omitempty" bson:"_id,omitempty"`

	// required
	PatientId  string    `json:"patientId" bson:"patientId"`
	DiagnoseId string    `json:"diagnoseId" bson:"diagnoseId"`
	Date       time.Time `json:"modified" bson:"modified"`
	Diagnose   string    `json:"diagnose" bson:"diagnose"`

	// optional
	Prescription string `json:"prescription,omitempty" bson:"prescription,omitempty"`

	// westack defaults
	Created time.Time `json:"created,omitempty" bson:"created,omitempty"`
}
