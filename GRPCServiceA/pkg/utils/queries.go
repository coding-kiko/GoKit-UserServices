package utils

var GetQueryByEmail string = "SELECT id, name, age, email, nationality, job, created FROM Users WHERE email=?"
var GetQueryById string = "SELECT id, name, age, email, nationality, job, created FROM Users WHERE id=?"
var CreateQuery string = "INSERT INTO Users (id, name, age, email, pwd, nationality, job, created) VALUES (?,?,?,?,?,?,?,?)"
