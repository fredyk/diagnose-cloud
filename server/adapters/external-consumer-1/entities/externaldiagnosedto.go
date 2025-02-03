package entities

import "time"

type ExternalDiagnoseDto struct {
	Id string `json:"id,omitempty"`

	// required
	PatientName string    `json:"patient_name"`
	Diagnose    string    `json:"diagnose"`
	Date        time.Time `json:"modified"`

	// optional
	Prescription string `json:"prescription"`

	// westack defaults
	Created time.Time `json:"created,omitempty"`
}
