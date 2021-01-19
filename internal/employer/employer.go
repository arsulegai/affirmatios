package employer

import (
	"affirmatios/employer/internal/aagent"
	"affirmatios/employer/internal/credential"
	"affirmatios/employer/web"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// employer is used to issue credentials
type Employer struct {
	ID         string                `json:"connection_id"`
	Credential credential.Credential `json:"credential"`
}

// GetAPI for employer has to register a schema before it
// is sent to the receiver
func (h *Employer) GetAPI() string {
	return "/employer/issue"
}

// GetHandler returns the handler method for the request
// in this case issues the credential to the established connection
func (h *Employer) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		requestedBodyBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		// read the request body, has the connection id and the credential
		err = json.Unmarshal(requestedBodyBytes, h)
		log.Printf("%v", *h)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		// Use the credential to call the Aries agent
		// Response back from the agent
		respBody, err := aagent.IssueCredential(h.ID, h.Credential)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		web.Success(writer, respBody)
	}
}

// GetMethod returns the POST
func (h *Employer) GetMethod() string {
	return http.MethodPost
}

// ViewCredential has an API to establish a connection
type ViewCredential struct{}

// GetAPI for all the credentials
func (v *ViewCredential) GetAPI() string {
	return "/employer/view"
}

// GetHandler for the returning all credentials
func (v *ViewCredential) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Call the agent and create the information
		respBody, err := aagent.ViewCredential()
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		web.Success(writer, respBody)
	}
}

// GetMethod for EstablishedConnections is GET
func (v *ViewCredential) GetMethod() string {
	return http.MethodGet
}

// GetServices returns all the services from employer
func (h *Employer) GetServices() []web.Service {
	var services []web.Service
	v := ViewCredential{}
	services = append(services, h)
	services = append(services, &v)
	return services
}
