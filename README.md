# Building CMPE273-lab2 project

```
cd CMPE273-lab2
export GOPATH=$PWD # Set GO PATH
go get -d -v ./... # Get dependencies 
go build -v ./...  # Build the project
go build src/httprouter/Httprouter.go # Build Server 
```

# Run the server using
```
go run src/httprouter/Server.go
```
or 
```
./Httprouter 
```

# API calls
## Get 'Hello, user'

```
curl -X GET http://localhost:8080/hello/Renu
```

Response is
```
Hello, Renu!
```

## Create (POST) a greeting 

```
curl -X POST -H 'Content-Type: application/json' -d '{"name": "Renu"}' http://localhost:8080/hello
```

Response is
```
{"Greeting":"Hello, Renu!"}
```

