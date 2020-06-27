# GO Microservices

## Setting up server
* clone the repo with UTL `https://github.com/AjayKumarMP/GO_microservices_boilerplate.git`
* `cd GO_microservices_boilerplate/`
* Install GO, protoc-3.12.3-linux-x86_64(as per OS) & set to PATH variable
* Set `GOPATH` & `GOROOT` and add it PATH env variable
* Packages Gin, protobuf, google.golan.org/grpc, github.com/golang/protobuf/protoc-gen-go. use command ge get -u `packa_name`
* Run this command to generate `*.pb.go` file(`GRPC`)
`protoc --proto_path=services/authService/proto --proto_path=third_party --go_out=plugins=grpc:services/authService/proto authService.proto`
* Switch to `cd services/authService` run `go run main.go`. This will start the service in TCP as GRPC service(tcp://8080).
* To Run the REST server switch to `cd /api` and run `go run main.go`. After executing thi REST server will up & running.(http://localhost:8080)
* To test Api's call the API `http://localhost:8080/auth/register` a POST method which will return {success: bool, message: string} as response.

## Happy Coding!!