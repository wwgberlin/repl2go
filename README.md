# repl2go
Generate valid Go code from your Go REPL instructions

## Overview

This project aims to make it easy to quickly write some Go instructions and see the output.

It could be used to:
- test the output of a Go code snippet quickly
- iteratively write Go code and get a valid file in the end

It is built upon [yaegi](https://github.com/traefik/yaegi), a interpreter for Go.

### Architecture

Repl2Go will consist of two major components:
- a web server
- a web app

The web app provides a GUI to enter Go statements and see the output per statement as well as a generated valid Go file from all statements. This is done by sending all new statements as well as the full history of previous statements to the web server.
The web server responds with the result of the latest input as well as the valid Go file.

![Repl2Go Architecture](docs/repl2go-architecture.png?raw=true "Architecture Overview")


### The GUI

![Repl2Go GUI](docs/repl2go-gui.png?raw=true "Sketch of the web app's GUI")


### The API


`POST /run`
- requestBody: 
    - **Hist**: Array of
        - **Input**: string # A Go statement
        - **Outputs**: Array of string # represents all results of interpreting the input
    - **Last**: string # A Go statement
- response:
    - **Outputs**: Array of string # ??? needs to be defined
    - **File**: string # a valid string that represents the content of a go file


## Development

Run all tests:
```go
go test
```

Start the server locally:
```go
go run
```
This will start the server on your local machine on Port 8000.

To try the functionality, you can send messages to `localhost:8000` using e.g. a tool like curl (in your shell) or [Postman](https://www.postman.com/).

### Example Calls

To send an inital message (with empty history) using curl:
```
curl ....
```

To send a message with some history using curl:
```
curl ....
```

To send an inital message (with empty history) using Postman:
- set HTTP Method to `POST`
- set URL to `localhost:8000/run`
- set the format for body to ??

Enter a valid JSON Object as the body:
```JSON
{
    "Last": "fmt.Println(\"Hello World\")",
    "Hist": []
}
```