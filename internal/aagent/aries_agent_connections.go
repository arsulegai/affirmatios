package aagent

// GetConnections returns the agent connections
func GetConnections() ([]byte, error) {
	agent := GetAgent()
	return agent.getConnections()
}

// AcceptConnection accepts a connection rquest via the ID
func AcceptConnection(message []byte) ([]byte, error) {
	agent := GetAgent()
	return agent.acceptConnection(message)
}

// CreateConnection will create a new connection
func CreateConnection() ([]byte, error) {
	agent := GetAgent()
	return agent.createConnection()
}
