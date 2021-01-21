package aagent

import (
	"affirmatios/employer/internal/credential"
	"affirmatios/employer/internal/ledger"
	"log"
)

// AriesCredential is used to send the request to the Aries
type AriesCredential struct {
	AutoRemove         bool                    `json:"auto_remove"`
	IssuerDID          string                  `json:"issuer_did"`
	SchemaVersion      string                  `json:"schema_version"`
	CredentialProposal AriesCredentialProposal `json:"credential_proposal"`
	Comment            string                  `json:"comment"`
	CredDefID          string                  `json:"cred_def_id"`
	SchemaIssuerID     string                  `json:"schema_issuer_id"`
	SchemaName         string                  `json:"schema_name"`
	ConnectionID       string                  `json:"connection_id"`
	SchemaID           string                  `json:"schema_id"`
	RevoRegID          string                  `json:"revoc_reg_id"`
}

// AriesCredentialProposal has information to be sent
type AriesCredentialProposal struct {
	Type       string           `json:"@type"`
	Attributes []AriesAttribute `json:"attributes"`
}

// AriesAttribute is what is issued
type AriesAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// IssueCredential sends the request to the Aries agent
func IssueCredential(cID string, c credential.Credential) ([]byte, error) {
	// Get the credential structure
	// Fill up in the format required by Aries Agent
	ariesCred := getAriesCredential(cID, c)
	log.Printf("%v", ariesCred)
	// Send to the Aries Agent
	agent := GetAgent()
	return agent.sendCredential(ariesCred)
}

func getAriesCredential(cID string, c credential.Credential) AriesCredential {
	ariesAttributes := []AriesAttribute{}
	ariesAttributes = append(ariesAttributes, AriesAttribute{
		Name:  "name",
		Value: c.Name,
	})
	ariesAttributes = append(ariesAttributes, AriesAttribute{
		Name:  "employee_id",
		Value: c.EmployeeId,
	})
	ariesAttributes = append(ariesAttributes, AriesAttribute{
		Name:  "joining_date",
		Value: c.JoiningDate,
	})
	ariesAttributes = append(ariesAttributes, AriesAttribute{
		Name:  "department",
		Value: c.Department,
	})
	ariesAttributes = append(ariesAttributes, AriesAttribute{
		Name:  "role",
		Value: c.Role,
	})
	ariesAttributes = append(ariesAttributes, AriesAttribute{
		Name:  "issued_date",
		Value: c.IssuedDate,
	})
	ariesAttributes = append(ariesAttributes, AriesAttribute{
		Name:  "relieving_date",
		Value: c.RelievingDate,
	})
	ariesCredentialProposal := AriesCredentialProposal{
		Type:       "issue-credential/1.0/credential-preview",
		Attributes: ariesAttributes,
	}
	ariesCredential := AriesCredential{
		AutoRemove:         true,
		IssuerDID:          ledger.GetSchemaIssuerID(),
		SchemaVersion:      credential.SchemaVersion,
		CredentialProposal: ariesCredentialProposal,
		Comment:            "Employer Issued The Credential",
		CredDefID:          ledger.GetCredentialDefID(),
		SchemaIssuerID:     ledger.GetSchemaIssuerID(),
		SchemaName:         ledger.GetSchemaName(),
		ConnectionID:       cID,
		SchemaID:           ledger.GetSchemaID(),
	}
	return ariesCredential
}

// ViewCredential proxies the request to agent
func ViewCredential() ([]byte, error) {
	// Call agent to know what's there
	agent := GetAgent()
	return agent.viewIssuedCredential()
}

func RequestProof(message []byte) ([]byte, error) {
	agent := GetAgent()
	return agent.requestProof(message)
}

// VerifyCredential will verify the credential id
func VerifyCredential(credId string) ([]byte, error) {
	agent := GetAgent()
	return agent.verifyCredential(credId)
}

// PPList will respond back all the presented proof
func PPList() ([]byte, error) {
	agent := GetAgent()
	return agent.ppList()
}
