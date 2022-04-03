package entity

type User struct {
	Id       string `bson:"_id,omitempty"`
	Username string `bson:"username,omitempty"`
	Password string `bson:"password,omitempty"`
}
