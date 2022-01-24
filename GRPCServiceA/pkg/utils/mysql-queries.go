package utils

import (
	"strings"
)

const (
	GetQueryByEmail string = "SELECT id, name, age, email, nationality, job, created FROM Users WHERE email=?"
	GetQueryById    string = "SELECT id, name, age, email, nationality, job, created FROM Users WHERE id=?"
	CreateQuery     string = "INSERT INTO Users (id, name, age, email, pwd, nationality, job, created) VALUES (?,?,?,?,?,?,?,?)"
	DeleteById      string = "DELETE FROM Users WHERE id=?"
	DeleteByEmail   string = "DELETE FROM Users WHERE email=?"
	UpdateByEmail   string = "UPDATE Users SET name=?, age=?, email=?, pwd=?, nationality=?, job=? where email=?"
	UpdateById      string = "UPDATE Users SET name=?, age=?, email=?, pwd=?, nationality=?, job=? where id=?"
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

func UpdateQuery(s string) string {
	if strings.Contains(s, "@") {
		return UpdateByEmail
	}
	return UpdateById
}
