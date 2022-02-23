package entities

type User struct {
	Id      string `bson:"_id"`
	Age     int64
	Name    string
	Job     string
	Country string
	PwdHsh  string
	Created string
	Email   string
}
