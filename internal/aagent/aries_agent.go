package aagent

import (
	"encoding/json"
	"log"
	"math/rand"
)

// AriesAgent is used for connections with the agent
type AriesAgent struct {
	Host string
	Port string
}

var agent *AriesAgent = nil

// GetAgent returns the agent and creates one if not present already
func GetAgent() AriesAgent {
	return *agent
}

// InitAgent creates an agent connector
func InitAgent(host string, port string) {
	if agent == nil {
		agent = &AriesAgent{
			Host: host,
			Port: port,
		}
	}
}

func (agent AriesAgent) sendCredential(ariesCredential AriesCredential) ([]byte, error) {
	// http operation to send to the agent
	message, err := json.Marshal(ariesCredential)
	if err != nil {
		return nil, err
	}
	log.Printf("Sending %v", string(message))
	return agent.post(message, "/issue-credential/send")
}

func (agent AriesAgent) viewIssuedCredential() ([]byte, error) {
	// http operation to send to the agent
	return agent.get("/issue-credential/records")
}

func (agent AriesAgent) getConnections() ([]byte, error) {
	// http operation to send to the agent
	return agent.get("/connections")
}

func (agent AriesAgent) acceptConnection(message []byte) ([]byte, error) {
	return agent.post(message, "/connections/receive-invitation")
}

func (agent AriesAgent) createConnection() ([]byte, error) {
	alias := RandStringBytes(6)
	return agent.post(nil, "/connections/create-invitation?alias="+alias)
}

func (agent AriesAgent) requestProof(message []byte) ([]byte, error) {
	return agent.post(message, "/present-proof/send-request")
}

func (agent AriesAgent) verifyCredential(message string) ([]byte, error) {
	return agent.post(nil, "/present-proof/records/"+message+"/verify-presentation")
}

func (agent AriesAgent) ppList() ([]byte, error) {
	return agent.get("/present-proof/records")
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandStringBytes generates a random string
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
