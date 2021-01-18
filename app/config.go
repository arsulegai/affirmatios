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

// GetAriesHost returns the aries host configuration
func (c *Config) GetAriesHost() string {
	return getOrDefaultEnv("ARIES_HOST", "hospital-agent")
}

// GetAriesPort returns the port
func (c *Config) GetAriesPort() string {
	return getOrDefaultEnv("ARIES_PORT", "8081")
}

// GetHost is where the server is bound
func (c *Config) GetHost() string {
	return getOrDefaultEnv("HTTP_HOST", "0.0.0.0")
}

// GetPort is where the server is binding the port
func (c *Config) GetPort() string {
	return getOrDefaultEnv("HTTP_PORT", "8080")
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

// GetMethod for config is GET
func (c *Config) GetMethod() string {
	return http.MethodGet
}

// GetServices returns all the services associated with config
func (c *Config) GetServices() []web.Service {
	var services []web.Service
	services = append(services, c)
	return services
}

func getOrDefaultEnv(key string, defaultValue string) string {
	if _, exists := os.LookupEnv(key); !exists {
		return defaultValue
	}
	return os.Getenv(key)
}
