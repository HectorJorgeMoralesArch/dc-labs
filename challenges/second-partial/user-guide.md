User Guide
==========
>! Created by
>! Hector Jorge Morales Arch as JarlArchJernRauda
>! and Juan Pablo as PuPumPa

## Requirements

- Hove installed Goland 
- Download the folder *second-partial*
- Open a terminal and type the next command:
$ go get github.com/gorilla/mux

# Using the API

# Starting the Server
For the function of this API you must have open two terminals.
The first terminal will run the server, which will receive the requests from the second terminal, therefore you must type the following command:
$ go run api.go

# Client
In the other terminal you can put 4 types of comands
   - `/ Login`
   - `/ Logout`
   - `/ Upload`
   - `/ Status`

# Login

The __*Login*__ command will receive an username and a password, and returns an access token.
$ curl -u username:password http://localhost:8080/login

It displays the information as:

{
	"message": "Hi username, welcome to the DPIP System",
	"token" "OjIE89GzFw"
}

# Logout

The __*Logout*__ command will receive the generated token from the user and erases it.
$ curl -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/logout

It displays the information as:

{
	"message": "Bye username, your token has been revoked"
}

# Upload

The __*Upload*__ command will receive an image, the path to the file, sent by the client and return the filename and the size.
$ curl -F 'data=@/path/to/local/image.png' -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/upload

It displays the information as:

{
	"message": "An image has been successfully uploaded",
	"filename": "image.png",
	"size": "500kb"
}


# Status

the __*Status*__ command will receive the token of the user and show a message with the time.
$ curl -H "Authorization: Bearer <ACCESS_TOKEN>" http://localhost:8080/status

It displays the information as:

{
	"message": "Hi username, the DPIP System is Up and Running"
	"time": "2015-03-07 11:06:39"
}