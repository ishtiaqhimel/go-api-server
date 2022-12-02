# API server Using Go

[![Go Report Card](https://goreportcard.com/badge/github.com/ishtiaqhimel/go-api-server)](https://goreportcard.com/report/github.com/ishtiaqhimel/go-api-server)

### RESTful API using [go](https://github.com/golang), [cobra CLI](https://github.com/spf13/cobra), [go-chi/chi](https://github.com/go-chi/chi), Basic Auth, [JWT Auth](https://github.com/dgrijalva/jwt-go)

--- 
API Endpoints
| Endpoint | Function | Method | StatusCode | Auth |
| -------- | -------- | ------ | ---------- | ---- |
| `/api/login` | LogIn | POST | Success - 200, Failure - 401 | Basic |
| `/api/student` | StudentGet | GET | Success - 200, Failure - 401 | JWT |
| `/api/student` | StudentPost | POST | Success - 200, Failure - 401, 409 | JWT |
| `/api/student/{id}` | StudentUpdate | PUT | Success - 200, Failure - 401, 404 | JWT |
| `/api/student/{id}` | StudentDelete | DELETE | Success - 200, Failure - 401, 404 | JWT |
| `/api/subject` | SubjectGet | GET | Success - 200, Failure - 401 | JWT |
| `/api/subject` | SubjectPost | POST | Success - 200, Failure - 401, 409 | JWT |
| `/api/subject/{id}` | SubjectUpdate | PUT | Success - 200, Failure - 401, 404 | JWT |
| `/api/subject/{id}` | SubjectDelete | DELETE | Success - 200, Failure - 401, 404 | JWT |

---
Data Model
```
package model

type Student struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Subjects  []Subject `json:"subjects"`
}

```
```
package model

type Subject struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Code  string `json:"code"`
}

```

---

Installation
* `go install github.com/ishtiaqhimel/go-api-server`

---

CLI Commands
* build the app locally `make build`
* help with the start commands `./bin/apiserver start -h` or `./bin/apiserver start --help`
---
Authentication Method
* Basic Authentication
* JWT Authentication

---
Resources:
* [sysdevbd learn GO](https://sysdevbd.com/go/)
* [A Beginner’s Guide to HTTP and REST](https://code.tutsplus.com/tutorials/a-beginners-guide-to-http-and-rest--net-16340)
* [HTTP Response Status Codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)

