package entities

type GetUserResp struct {
	Id          string
	Name        string
	Age         int64
	Nationality string
	Job         string
	Created     string
	Email       string
}

type CreateUserResp struct {
	Id      string
	Created string
}
