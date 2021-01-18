package aagent

import (
	"affirmatios/hospital/internal/credential"
	"affirmatios/hospital/internal/ledger"
)

// AriesCredential is used to send the request to the Aries
type AriesCredential struct {
	AutoRemove         bool                    `json: "auto_remove"`
	IssuerDID          string                  `json: "issuer_did"`
	RevoRegID          string                  `json: "revoc_reg_id"`
	SchemaVersion      string                  `json: "schema_version"`
	CredentialProposal AriesCredentialProposal `json: "credential_proposal"`
	Comment            string                  `json: "comment"`
	CredDefID          string                  `json: "cred_def_id"`
	SchemaIssuerID     string                  `json: "schema_issuer_id"`
	SchemaName         string                  `json: "schema_name"`
	ConnectionID       string                  `json: "connection_id"`
	SchemaID           string                  `json: "schema_id"`
}

// AriesCredentialProposal has information to be sent
type AriesCredentialProposal struct {
	Type       string           `json: "@Type"`
	Attributes []AriesAttribute `json: "attributes"`
}

// AriesAttribute is what is issued
type AriesAttribute struct {
	Name  string `json: "name"`
	Value string `json: "value"`
}

// IssueCredential sends the request to the Aries agent
func IssueCredential(cID string, c credential.Credential) ([]byte, error) {
	// Get the credential structure
	// Fill up in the format required by Aries Agent
	ariesCred := getAriesCredential(cID, c)
	// Send to the Aries Agent
	agent := GetAgent()
	return agent.sendCredential(ariesCred)
}

func getAriesCredential(cID string, c credential.Credential) AriesCredential {
	ariesCredentialProposal := AriesCredentialProposal{}
	ariesCredential := AriesCredential{
		AutoRemove:         true,
		IssuerDID:          "Th7MpTaRZVRYnPiabds81Y",
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
