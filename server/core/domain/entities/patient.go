package entities

type Patient struct {
	Id string `json:"id,omitempty" bson:"_id,omitempty"`

	// required
	Name           string `json:"name" bson:"name"`
	DocumentNumber string `json:"documentNumber" bson:"documentNumber"`
	Email          string `json:"email" bson:"email"`

	// optional
	Phone   string `json:"phone,omitempty" bson:"phone,omitempty"`
	Address string `json:"address,omitempty" bson:"address,omitempty"`
}
