package entities

import "time"

type InternalDiagnose struct {
	Id string `json:"id,omitempty" bson:"_id,omitempty"`

	// required
	PatientId  string    `json:"patientId" bson:"patientId"`
	DiagnoseId string    `json:"diagnoseId" bson:"diagnoseId"`
	Date       time.Time `json:"modified" bson:"modified"`

	// optional
	Prescription string `json:"prescription" bson:"prescription,omitempty"`

	// westack defaults
	Created time.Time `json:"created,omitempty" bson:"created,omitempty"`
}
