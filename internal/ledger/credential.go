package ledger

import "os"

// GetCredentialDefID returns the def id from the ledger
func GetCredentialDefID() string {
	return getOrDefaultEnv("CRED_DEF_ID", "Th7MpTaRZVRYnPiabds81Y:3:CL:23:default")
}

// GetCredentialDefID returns the def id from the ledger
func GetDegreeCredentialDefID() string {
	return getOrDefaultEnv("DEGREE_CRED_DEF_ID", "Th7MpTaRZVRYnPiabds81Y:3:CL:23:default")
}

// GetCredentialDefID returns the def id from the ledger
func GetHealthCredentialDefID() string {
	return getOrDefaultEnv("HEALTH_CRED_DEF_ID", "Th7MpTaRZVRYnPiabds81Y:3:CL:23:default")
}

// GetSchemaID returns the schema id that is generated earlier
func GetSchemaID() string {
	return getOrDefaultEnv("SCHEMA_ID", "Th7MpTaRZVRYnPiabds81Y:2:employer_record:1.0")
}

// GetSchemaIssuerID returns the DID of the issuer
func GetSchemaIssuerID() string {
	return getOrDefaultEnv("SCHEMA_ISSUER_ID", "Th7MpTaRZVRYnPiabds81Y")
}

// GetSchemaName returns the name of the schema that was added initially
func GetSchemaName() string {
	return getOrDefaultEnv("SCHEMA_NAME", "employer_record")
}

func getOrDefaultEnv(key string, defaultValue string) string {
	if _, exists := os.LookupEnv(key); !exists {
		return defaultValue
	}
	return os.Getenv(key)
}
