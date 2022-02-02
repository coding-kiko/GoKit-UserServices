package entities

type GetUserReq struct {
	Id string
}

type CreateUserReq struct {
	Pwd     string
	Name    string
	Age     int64
	Country string
	Job     string
	Email   string
}

type DeleteUserReq struct {
	Id string
}

type UpdateUserReq struct {
	Pwd     string
	Name    string
	Age     int64
	Country string
	Job     string
	Email   string
}

type AuthenticateReq struct {
	Email string
	Pwd   string
}
