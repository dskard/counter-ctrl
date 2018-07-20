## A Few Things I Learned About Golang

### Useful links

1. [Effective Go](https://golang.org/doc/effective_go.html)

### The Things

1. Case matters for field names inside a struct, and struct names. Uppercase means public, lowercase means private?


```
type Cmdargs struct {
  start string `json:"start"`
}

var cargs Cmdargs

log.Printf("startVal: %v", cargs.start)
```

vs

```
type Cmdargs struct {
  Start string `json:"start"`
}

var cargs Cmdargs

log.Printf("startVal: %v", cargs.Start)
```

2. Gorilla Mux - building a RESTful API server is easy


3. Object oriented programming: https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

```
// app.go

package main

import (
    "github.com/gorilla/mux"
)

type App struct {
    Router *mux.Router
    DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) { }

func (a *App) Run(addr string) { }
```

Highlight the `(a *App)` part of the function.

4. Retrieving packages using `go get …`

5. `if` statements need `{` on the same line
https://stackoverflow.com/a/22948654
