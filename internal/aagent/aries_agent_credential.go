package aagent

import (
	"affirmatios/hospital/internal/credential"
	"affirmatios/hospital/internal/ledger"
	"fmt"
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
		Name:  "sex",
		Value: c.Sex,
	})
	ariesAttributes = append(ariesAttributes, AriesAttribute{
		Name:  "age",
		Value: fmt.Sprintf("%d", c.Age),
	})
	ariesAttributes = append(ariesAttributes, AriesAttribute{
		Name:  "address",
		Value: c.Address,
	})
	ariesAttributes = append(ariesAttributes, AriesAttribute{
		Name:  "place",
		Value: c.Place,
	})
	ariesAttributes = append(ariesAttributes, AriesAttribute{
		Name:  "date",
		Value: c.Date,
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
		Comment:            "Hospital Issued The Credential",
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
