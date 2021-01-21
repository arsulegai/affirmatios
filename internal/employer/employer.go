package employer

import (
	"affirmatios/employer/internal/aagent"
	"affirmatios/employer/internal/credential"
	"affirmatios/employer/internal/ledger"
	"affirmatios/employer/web"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

// VerifyCredential has an API to establish a connection
type VerifyCredential struct {
	CredentialExchangeID string `json:"credential_exchange_id"`
}

// GetAPI for all the credentials
func (v *VerifyCredential) GetAPI() string {
	return "/employer/verify"
}

// GetHandler for the returning all credentials
func (v *VerifyCredential) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		requestedBodyBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		// read the request body, has the exchange id and the credential
		err = json.Unmarshal(requestedBodyBytes, v)
		log.Printf("%v", v)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		// Call the agent and create the information
		respBody, err := aagent.VerifyCredential(v.CredentialExchangeID)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		web.Success(writer, respBody)
	}
}

// GetMethod for EstablishedConnections is GET
func (v *VerifyCredential) GetMethod() string {
	return http.MethodPost
}

// ListPresentedCred has an API to establish a connection
type ListPresentedCred struct{}

// GetAPI for all the credentials
func (v *ListPresentedCred) GetAPI() string {
	return "/employer/presented-proofs"
}

// GetHandler for the returning all credentials
func (v *ListPresentedCred) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Call the agent and create the information
		respBody, err := aagent.PPList()
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		web.Success(writer, respBody)
	}
}

// GetMethod for EstablishedConnections is GET
func (v *ListPresentedCred) GetMethod() string {
	return http.MethodGet
}

// GetServices returns all the services from employer
func (h *Employer) GetServices() []web.Service {
	var services []web.Service
	v := ViewCredential{}
	p := RequestProof{}
	f := VerifyCredential{}
	l := ListPresentedCred{}
	services = append(services, h)
	services = append(services, &v)
	services = append(services, &p)
	services = append(services, &f)
	services = append(services, &l)
	return services
}

type RequestProof struct {
	ConnectionId string `json:"connection_id"`
	RecordType   string `json:"record_type"`
}

func (p *RequestProof) GetAPI() string {
	return "/employer/request-proof"
}

func (p *RequestProof) GetMethod() string {
	return http.MethodPost
}

func (p *RequestProof) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		requestedBodyBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		// read the request body, has the connection id and the credential
		err = json.Unmarshal(requestedBodyBytes, p)
		log.Printf("%v", *p)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		credId, body := GetCredId(p.RecordType)

		body = strings.ReplaceAll(body, "$$CONNECTIONID$$", p.ConnectionId)
		body = strings.ReplaceAll(body, "$$CREDID$$", credId)
		// Response back from the agent
		respBody, err := aagent.RequestProof([]byte(body))
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		web.Success(writer, respBody)
	}
}

func GetCredId(recordType string) (string, string) {
	credId := ledger.GetCredentialDefID() //this is the employer record
	body := employerRecord
	switch recordType {
	case "Hospital":
		credId = ledger.GetHealthCredentialDefID()
		body = healthRecord
	case "Degree":
		credId = ledger.GetDegreeCredentialDefID()
		body = degreeRecord
	}
	return credId, body
}
