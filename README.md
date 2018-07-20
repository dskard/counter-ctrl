# Counter Controller
## Golang server with API interface for controlling an external Counter program

### Installing Go

It might be easiest to download a Golang Docker image and use that for development:

```
docker pull golang:1.10
```

You can find other available tags on the [docker hub page](https://hub.docker.com/_/golang/)

### Launching the Docker container

Set `HOME` so we can use the local cache in `~/.cache/go-build`
```
chex -i golang:1.10 -a "-e HOME=/opt/shared/${USER}" bash
```

### Dependencies

We need the mux library to route requests
```
go get github.com/gorilla/mux
```

### Building and Running the Binary

```
make all
make run
```

### Starting and Stopping

Use `curl` to send `start` and `stop` commands to the `go` server process

```
$ curl -v -s -X POST -H "Content-Type: application/json" -d '{"start":"8"}' localhost:4723/start
*   Trying ::1...
* Connected to localhost (::1) port 4723 (#0)
> POST /start HTTP/1.1
> Host: localhost:4723
> User-Agent: curl/7.47.0
> Accept: */*
> Content-Type: application/json
> Content-Length: 13
> 
* upload completely sent off: 13 out of 13 bytes
< HTTP/1.1 201 Created
< Content-Type: application/json
< Date: Fri, 20 Jul 2018 04:18:43 GMT
< Content-Length: 11
< 
* Connection #0 to host localhost left intact
{"pid":645}
$
$
$
$ curl -v -s -X GET localhost:4723/stop
*   Trying ::1...
* Connected to localhost (::1) port 4723 (#0)
> GET /stop HTTP/1.1
> Host: localhost:4723
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Fri, 20 Jul 2018 04:19:10 GMT
< Content-Length: 17
< 
* Connection #0 to host localhost left intact
{"status":"done"}
```

### Useful articles:

0. [Learning Go: the tour](https://tour.golang.org/list)
1. [Building a RESTful API with Golang](https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo)
2. [Starting a command: func (*Cmd) Start](https://golang.org/pkg/os/exec/#Cmd.Start)
3. [Killing a child process and all of its children in Go](https://medium.com/@felixge/killing-a-child-process-and-all-of-its-children-in-go-54079af94773)
4. [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)


