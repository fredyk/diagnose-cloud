package entities

type ExternalPatient struct {

	// Not present in external API
	Id string `json:"-" bson:"-"`

	// required
	Name           string `json:"name" bson:"name"`
	DocumentNumber string `json:"document_number" bson:"document_number"`
	Email          string `json:"email" bson:"email"`

	// optional
	Phone   string `json:"phone,omitempty" bson:"phone,omitempty"`
	Address string `json:"address,omitempty" bson:"address,omitempty"`
}
