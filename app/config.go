package app

import (
	"affirmatios/hospital/web"
	"net/http"
	"os"
)

// Config is the information about the application
type Config struct {
	Name    string
	Version string
}

// GetConfig gets the application configuration
func GetConfig(appName, appVersion string) *Config {
	config := Config{
		Name:    appName,
		Version: appVersion,
	}
	return &config
}

// GetHost is where the server is bound
func (c *Config) GetHost() string {
	return os.Getenv("HTTP_HOST")
}

// GetPort is where the server is binding the port
func (c *Config) GetPort() string {
	return os.Getenv("HTTP_PORT")
}

// GetAPI returns the API to be added
func (c *Config) GetAPI() string {
	return "/info"
}

// GetHandler returns the handler for the API
func (c *Config) GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		respBody, err := web.StructToBytes(c)
		if err != nil {
			web.BadRequest(writer, err)
			return
		}
		web.Success(writer, respBody)
	}
}

func (c *Config) getServices() []web.Service {
	var services []web.Service
	services = append(services, c)
	return services
}
