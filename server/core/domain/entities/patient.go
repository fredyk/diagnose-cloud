package entities

import "time"

type Patient struct {
	Id string `json:"id,omitempty" bson:"_id,omitempty"`

	// required
	Name           string    `json:"name" bson:"name"`
	DocumentNumber string    `json:"documentNumber" bson:"documentNumber"`
	Email          string    `json:"email" bson:"email"`
	DOB            time.Time `json:"dob" bson:"dob"` // Date of Birth

	// optional
	Phone   string `json:"phone,omitempty" bson:"phone,omitempty"`
	Address string `json:"address,omitempty" bson:"address,omitempty"`

	// westack defaults
	Created  time.Time `json:"__created,omitempty" bson:"created,omitempty"`
	Modified time.Time `json:"__modified,omitempty" bson:"modified,omitempty"`
}
