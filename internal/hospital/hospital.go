package hospital

import (
	"affirmatios/hospital/internal/aagent"
	"affirmatios/hospital/internal/credential"
	"affirmatios/hospital/web"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Hospital is used to issue credentials
type Hospital struct {
	ConnectionID string
	Credential   credential.Credential
}

// GetAPI for Hospital has to register a schema before it
// is sent to the receiver
func (h *Hospital) GetAPI() string {
	return "/hospital/issue"
}

// GetHandler returns the handler method for the request
// in this case issues the credential to the established connection
func (h *Hospital) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		requestedBodyBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		// read the request body, has the connection id and the credential
		err = json.Unmarshal(requestedBodyBytes, h)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		// Use the credential to call the Aries agent
		// Response back from the agent
		respBody, err := aagent.IssueCredential(h.ConnectionID, h.Credential)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		web.Success(writer, respBody)
	}
}

// GetMethod returns the POST
func (h *Hospital) GetMethod() string {
	return http.MethodPost
}

// GetServices returns all the services from hospital
func (h *Hospital) GetServices() []web.Service {
	var services []web.Service
	services = append(services, h)
	return services
}
