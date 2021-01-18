package management

import (
	"affirmatios/hospital/web"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Connections proxies requests to the underlying Aries agent
type Connections struct {
}

// PendingConnections structure
// Has an API to return the pending connection requests
type PendingConnections struct {
	ConnectionID string
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
		// TODO: accept the connections
		err = json.Unmarshal(requestedBodyBytes, p)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		respBody, err := web.StructToBytes(p)
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
		// TODO: Query and send the connections
		respBody, err := web.StructToBytes(e)
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

// GetServices returns all the services for the connection
func (c *Connections) GetServices() []web.Service {
	var services []web.Service
	p := PendingConnections{}
	e := EstablishedConnections{}
	services = append(services, &p)
	services = append(services, &e)
	return services
}
