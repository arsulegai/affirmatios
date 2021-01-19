package user

import (
	"affirmatios/user/internal/aagent"
	"affirmatios/user/internal/credential"
	"affirmatios/user/web"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// user is used to issue credentials
type User struct {
	ID         string                `json:"connection_id"`
	Credential credential.Credential `json:"credential"`
}

// GetAPI for user has to register a schema before it
// is sent to the receiver
func (h *User) GetAPI() string {
	return "/user/issue"
}

// GetHandler returns the handler method for the request
// in this case issues the credential to the established connection
func (h *User) GetHandler() http.HandlerFunc {
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
func (h *User) GetMethod() string {
	return http.MethodPost
}

// ViewCredential has an API to establish a connection
type ViewCredential struct{}

// GetAPI for all the credentials
func (v *ViewCredential) GetAPI() string {
	return "/user/view"
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

type StoreCredential struct {
	ID string `json:"exchange_id"`
}

func (s *StoreCredential) GetAPI() string {
	return "/user/store"
}

func (s *StoreCredential) GetMethod() string {
	return http.MethodPost
}

//Handler for storing the credentials
func (s *StoreCredential) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Call the agent and create the information
		requestedBodyBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		log.Printf(string(requestedBodyBytes))
		err = json.Unmarshal(requestedBodyBytes, s)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		log.Printf("%v", *s)
		respBody, err := aagent.StoreCredentials(s.ID)

		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		web.Success(writer, respBody)
	}
}

type GetCredential struct{}

func (g *GetCredential) GetAPI() string {
	return "/credentials"

}
func (g *GetCredential) GetMethod() string {
	return http.MethodGet
}
func (g *GetCredential) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println("/credential Get")
		respBody, err := aagent.GetCredentials()
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		web.Success(writer, respBody)
	}
}

type GetCredentialById struct{}

func (g *GetCredentialById) GetAPI() string {
	return "/credentialsbyid"
}
func (g *GetCredentialById) GetMethod() string {
	return http.MethodGet
}
func (g *GetCredentialById) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseForm()
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		log.Println(request.Form)
		id := request.Form.Get("id")
		log.Println(id)
		respBody, err := aagent.GetCredentialsById(id)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		web.Success(writer, respBody)
	}
}

// GetServices returns all the services from user
func (h *User) GetServices() []web.Service {
	var services []web.Service
	v := ViewCredential{}
	s := StoreCredential{}
	g := GetCredential{}
	i := GetCredentialById{}
	services = append(services, h)
	services = append(services, &v)
	services = append(services, &s)
	services = append(services, &g)
	services = append(services, &i)
	return services
}
