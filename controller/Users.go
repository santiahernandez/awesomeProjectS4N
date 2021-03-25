package controller

type User struct {
	LegalId    string `json:"id" bson:"_id,omitempty"`
	Name       string `json:"name"`
	Profession string `json:"profession"`
	Gender     string `json:"gender"`
}
