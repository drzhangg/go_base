package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

type User struct {
	Id, Name string
}

type UserResource struct {
	users map[string]User
}

func (u UserResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/{user-id}").To(u.findUser))
	ws.Route(ws.POST("").To(u.updateUser))
	ws.Route(ws.PUT("/{user-id}").To(u.createUser))
	ws.Route(ws.DELETE("/{user-id}").To(u.removeUser))

	container.Add(ws)
}

// GET http://localhost:8090/users/1
//
func (u UserResource) findUser(request *restful.Request, response *restful.Response) {

	id := request.PathParameter("user-id")
	usr, ok := u.users[id]
	if !ok {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteAsJson(usr)
	}
}

// POST http://localhost:8090/users
// <User><Id>1</Id><Name>Melissa Raspberry</Name></User>
//
func (u *UserResource) updateUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
	err := request.ReadEntity(&usr)
	if err == nil {
		u.users[usr.Id] = *usr
		response.WriteEntity(usr)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// PUT http://localhost:8090/users/1
// <User><Id>1</Id><Name>Melissa</Name></User>
//
func (u *UserResource) createUser(request *restful.Request, response *restful.Response) {
	usr := User{Id: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	if err == nil {
		u.users[usr.Id] = usr
		response.WriteHeader(http.StatusCreated)
		response.WriteEntity(usr)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// DELETE http://localhost:8090/users/1
//
func (u *UserResource) removeUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	delete(u.users, id)
}

func RunRestfulCurlyRouterServer() {
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	u := UserResource{map[string]User{}}
	u.Register(wsContainer)

	log.Print("start listening on localhost:8090")
	server := &http.Server{Addr: ":8090", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

func waitForServerUp(serverURL string) error {
	for start := time.Now(); time.Since(start) < time.Minute; time.Sleep(5 * time.Second) {
		_, err := http.Get(serverURL + "/")
		if err == nil {
			return nil
		}
	}
	return errors.New("waiting for server timed out")
}

func main() {
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	u := UserResource{map[string]User{}}
	u.Register(wsContainer)

	log.Printf("start listening on localhost:8089")
	server := &http.Server{Addr: ":8089", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}