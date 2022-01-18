package entities

type GetUserReq struct {
	Id string
}

type CreateUserReq struct {
	Pwd         string
	Name        string
	Age         int64
	Nationality string
	Job         string
	Email       string
}

type DeleteUserReq struct {
	Id string
}
