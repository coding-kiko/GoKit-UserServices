package utils

import (
	"strings"
)

const (
	GetQueryByEmail string = "SELECT id, name, age, email, nationality, job, created FROM Users WHERE email=?"
	GetQueryById    string = "SELECT id, name, age, email, nationality, job, created FROM Users WHERE id=?"
	CreateQuery     string = "INSERT INTO Users (id, name, age, email, pwd, nationality, job, created) VALUES (?,?,?,?,?,?,?,?)"
)

func GetQuery(s string) string {
	if strings.Contains(s, "@") {
		return GetQueryByEmail
	}
	return GetQueryById
}
