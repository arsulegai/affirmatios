package management

import (
	"affirmatios/user/internal/aagent"
	"affirmatios/user/web"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Connections proxies requests to the underlying Aries agent
type Connections struct {
}

// PendingConnections structure
// Has an API to return the pending connection requests
type PendingConnections struct {
	CID string `json:"c_id"`
}

// GetAPI for accepting the pending connections
func (p *PendingConnections) GetAPI() string {
	return "/connections/accept"
}

// GetHandler for accepting the pending connections
func (p *PendingConnections) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		requestedBodyBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		err = json.Unmarshal(requestedBodyBytes, p)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		toSend, err := base64.StdEncoding.DecodeString(p.CID)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		log.Printf("%s", string(toSend))
		respBody, err := aagent.AcceptConnection(toSend)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		web.Success(writer, respBody)
	}
}

// GetMethod for PendingConnections is GET
func (p *PendingConnections) GetMethod() string {
	return http.MethodPost
}

// EstablishedConnections has an API to return the established connections
type EstablishedConnections struct{}

// GetAPI for all the connections
func (e *EstablishedConnections) GetAPI() string {
	return "/connections"
}

// GetHandler for the returning all connections
func (e *EstablishedConnections) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Call the agent and get the information
		respBody, err := aagent.GetConnections()
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		web.Success(writer, respBody)
	}
}

// GetMethod for EstablishedConnections is GET
func (e *EstablishedConnections) GetMethod() string {
	return http.MethodGet
}

// RequestConnection has an API to establish a connection
type RequestConnection struct{}

// GetAPI for all the connections
func (r *RequestConnection) GetAPI() string {
	return "/connections/request"
}

// GetHandler for the returning all connections
func (r *RequestConnection) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Call the agent and create the information
		respBody, err := aagent.CreateConnection()
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		web.Success(writer, respBody)
	}
}

// GetMethod for EstablishedConnections is GET
func (r *RequestConnection) GetMethod() string {
	return http.MethodPost
}

// GetServices returns all the services for the connection
func (c *Connections) GetServices() []web.Service {
	var services []web.Service
	p := PendingConnections{}
	e := EstablishedConnections{}
	r := RequestConnection{}
	services = append(services, &p)
	services = append(services, &e)
	services = append(services, &r)
	return services
}
