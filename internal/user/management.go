package user

import (
	"affirmatios/hospital/web"
	"net/http"

	"github.com/gorilla/sessions"
)

// Store is for sessions
var Store = sessions.NewCookieStore([]byte("user-logged-in"))

// GetStore returns the store
func GetStore() *sessions.CookieStore {
	return Store
}

// Management is used for login operation
type Management struct {
	Name     string
	Password string
}

// GetAPI for all the connections
func (m *Management) GetAPI() string {
	return "/login"
}

// GetHandler for the returning all connections
func (m *Management) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Call the agent and create the information
		session, err := Store.Get(request, "user-logged-in")
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		// Set some session values.
		session.Values["user-logged-in"] = "1"
		// Save it before we write to the response/return from the handler.
		err = session.Save(request, writer)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		web.Success(writer, []byte("Success"))
	}
}

// GetMethod for EstablishedConnections is GET
func (m *Management) GetMethod() string {
	return http.MethodPost
}

// GetServices returns the services for user management
func (m *Management) GetServices() []web.Service {
	var services []web.Service
	services = append(services, m)
	return services
}
