package entities

type GetUserResp struct {
	Id      string `bson:"_id"`
	Name    string
	Age     int64
	Country string
	Job     string
	Created string
	Email   string
	Error   Status
}

type CreateUserResp struct {
	Id      string `bson:"_id"`
	Created string
	Error   Status
}

type DeleteUserResp struct {
	Deleted string
	Error   Status
}

type UpdateUserResp struct {
	Updated string
	Error   Status
}

type Status struct {
	Message string
	Code    int32
}

type AuthenticateResp struct {
	Token string
	Error Status
}
