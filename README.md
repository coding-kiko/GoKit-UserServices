# About

**This branch is the exact same as the main one but implemented with Mongodb, the database part of the diagram is inaccurate.**\
**You can appreciate how go-kit's abstraction allowed be to touch only some functions without having to refactor the whole code and by only changing the repository layer and adding some helper functions I can change from mysql to mongodb**

The aim of this project was to create a simple micro-service using the Go-Kit toolkit and with the help of other libraries in order to feel more confident with the language and apply some of my acquired knowledge in Go to this interesting project. The core idea is to expose some simple methods that have to do with the manipulation of 'Users' in a database. I learned some new concepts like GRPC, protocol buffers, Jwt, etc. However some of the things are not fully implemented and need some more work.


# The architecture
The architecture of the project follows along with the one that the [Go-Kit](https://gokit.io/) developers pictured as this abstract layered architecture, with the core GRPC service having a Transport, Endpoints and Service layer. The Transport layer exposes the methods using GRPC, the service layer takes care of the business logic, and the endpoints layer works as a powerful abstraction tool that wraps up requests and responses going in between the other two layers, and because of this the responsibilities and problems of each layer are truly and well separated.
GRPC is great, and even though it is a clear superior in terms of speed and other factors when compared to the ordinary HTTP, it still loses in one category: compatibility. So, in order to be able to expose my api to either of the protocols, I implemented [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) to work as a reverse proxy and be able to call my methods as if it was a RESTFUL Api.

<div align="center">
  <img alt="JPG" src="UserServices/readme/architecture.jpg"/>
</div>

# Methods
The following methods will be called by a command line tool **curl**, but it could be perfectly achieved with any other advanced interface tool such as Postman.


|Service Method  |HTTP method     |Route                         |
|----------------|-------------------------------|-----------------------------|
|GetUser         |GET           |/User/{Id} **or** /User/{email}            |
|CreateUser          |POST            |/User           |
|DeleteUser          |DELETE|/User/{Id} **or** /User/{email} |
|UpdateUser |  PUT |/User|
|Authenticate| POST | /Authenticate|

\
Let's CREATE a new user:\
**Notice: all of the fields are required, otherwise the handling throws an error**
   

     curl -X POST localhost:8000/User -H 'Content-Type:application/json' -d '{"name":"Cristiano Ronaldo", "age": 36,"country":"pt", "job":"Footballer", "pwd":"12345678", "email": "cristiano.ronaldo@manudt.com"}'
        {
           "created" : "02-03-2022(10:04)",
           "id" : "64738e4a-b303-483c-84a0-8181cfa1a66c,
        }
\
Let's GET an user from the database:

    curl localhost:8000/User/cristiano.ronaldo@manudt.com 
    {
       "age" : "36",
       "country" : "pt",
       "created" : "02-03-2022(10:04)",
       "email" : "cristiano.ronaldo@manudt.com",
       "id" : "64738e4a-b303-483c-84a0-8181cfa1a66c",
       "job" : "Footballer",
       "name" : "Cristiano Ronaldo"
    }
\
Let's UPDATE our user:\
**Notice: all of the fields are required, otherwise the handling throws an error**

    curl -X PUT localhost:8000/User -H 'Content-Type:application/json' -d '{"name":"Cristiano Ronaldo", "age": 37,"country":"pt", "job":"Football Manager", "pwd":"12345678", "email": "cristiano.ronaldo@manudt.com"}'
    {
       "updated" : "02-03-2022(10:09)"
    }
 \
Finally let's DELETE our user:

    curl -X DELETE localhost:8000/User/cristiano.ronaldo@manudt.com
    {
	    "deleted":"02-03-2022(10:13)"
    }
  \
The AUTHENTICATE method is not fully implemented yet, but it does generate a perfectly valid token, that the client of the Api would need to consequently add to the request header in order to achieve any of the operations shown above.
A demo of what it looks like:

    curl -X POST localhost:8000/Authenticate -H 'Content-Type:application/json' -d '{"pwd":"admin", "email": "admin.admin@admin.com"}'
    {
	     "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoiYWRtaW4uYWRtaW5AYWRtaW4uY29tIiwiZXhwIjoxNjQzODk0NjI5LCJpYXQiOjE2NDM4OTQzMjksImlzcyI6ImxvY2FsaG9zdDo4MDAwLyJ9.X6J8pBybObkyOeLWPnUamkz8V93PW4HPqTRICRm3Chk"
    }

## Errors
This is a quick guide to some of the errors you may encounter while making use of the api:

|Error  |HTTP status code     |GRPC Code                      |
|----------------|-------------------------------|-----------------------------|
|User not found         |404           |5            |
|   Unkown error       |500            |2          |
|User already exists          |409|6 |
|Invalid argument(s) |  400 |3|
|Invalid credential(s)| 422 | 7|

\
Examples:\
You missed a field when creating a new user:

    curl -X POST localhost:8000/User -H 'Content-Type:application/json' -d '{"name":"Cristiano Ronaldo", "age": 36,"country":"pt", "job":"Footballer", "pwd":"12345678"}'
    {
	    "error": {
		    "code": 400,
		    "message": "invalid argument(s)"
	    }
	}
\
You tried to create a user with an email that already exists in the database:

    curl -X POST localhost:8000/User -H 'Content-Type:application/json' -d '{"name":"Lionel Messi", "age": 35,"country":"ar", "job":"Footballer", "pwd":"12345678", "email": "cristiano.ronaldo@manudt.com"}'
    {
	     "error": {
		        "code": 409,
        		"message": "user already exists"
	     }
    }
\
Getting a user that doesn't exists:

    curl localhost:8000/User/invalid-id
    {
	    "error": {
		    "code": 404,
		    "message": "user not found"
	    }
	}
\
Trying to authenticate with invalid credentials:

    curl -X POST localhost:8000/Authenticate -H 'Content-Type:application/json' -d '{"pwd":"wrong", "email": "invalid.email@invalid.com"}'
    {
       "error" : {
          "code" : 7,
          "message" : "invalid credential(s)"
       }
    }
