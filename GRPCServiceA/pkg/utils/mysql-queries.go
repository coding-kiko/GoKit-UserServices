package utils

import (
	"strings"
)

const (
	GetQueryByEmail   string = "SELECT id, name, age, email, country, job, created FROM Users WHERE email=?"
	GetQueryById      string = "SELECT id, name, age, email, country, job, created FROM Users WHERE id=?"
	CreateQuery       string = "INSERT INTO Users (id, name, age, email, pwd, country, job, created) VALUES (?,?,?,?,?,?,?,?)"
	DeleteById        string = "DELETE FROM Users WHERE id=?"
	DeleteByEmail     string = "DELETE FROM Users WHERE email=?"
	UpdateQuery       string = "UPDATE Users SET name=?, age=?, email=?, pwd=?, country=?, job=? where email=?"
	AuthenticateQuery string = "SELECT pwd FROM Users WHERE email=?"
)

func GetQuery(s string) string {
	if strings.Contains(s, "@") {
		return GetQueryByEmail
	}
	return GetQueryById
}

func DeleteQuery(s string) string {
	if strings.Contains(s, "@") {
		return DeleteByEmail
	}
	return DeleteById
}
