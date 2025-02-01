package entities

import "time"

type ExternalDiagnoseDto struct {
	Id string `json:"id,omitempty"`

	// required
	PatientName string    `json:"patient_name"`
	DiagnoseId  string    `json:"diagnose_id"`
	Date        time.Time `json:"modified"`

	// optional
	Prescription string `json:"prescription"`

	// westack defaults
	Created time.Time `json:"created,omitempty"`
}
