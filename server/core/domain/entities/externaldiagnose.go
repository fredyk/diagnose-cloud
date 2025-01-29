package entities

import "time"

type ExternalDiagnose struct {
	Id string `json:"id,omitempty" bson:"_id,omitempty"`

	// required
	PatientId  string    `json:"patient_id" bson:"patient_id"`
	DiagnoseId string    `json:"diagnose_id" bson:"diagnose_id"`
	Date       time.Time `json:"modified" bson:"modified"`

	// optional
	Prescription string `json:"prescription" bson:"prescription,omitempty"`

	// westack defaults
	Created time.Time `json:"created,omitempty" bson:"created,omitempty"`
}
