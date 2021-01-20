package employer

const employerRecord = `
{
	"connection_id": "$$CONNECTIONID$$",
	"comment": "request for proof",
	"proof_request": {
	  "name": "Proof of Education",
	  "version": "1.0",
	  "requested_attributes": {
		"0_name_uuid": {
		  "name": "name",
		  "restrictions": [
			{
			  "cred_def_id": "$$CREDID$$"
			}
		  ]
		},
		"0_joining_date_uuid": {
		  "name": "joining_date",
		  "restrictions": [
			{
			  "cred_def_id": "$$CREDID$$"
			}
		  ]
		},
		"0_relieving_date_uuid": {
		  "name": "relieving_date",
		  "restrictions": [
			{
			  "cred_def_id": "$$CREDID$$"
			}
		  ]
		},
		"0_role_uuid": {
			"name": "role",
			"restrictions": [
			  {
				"cred_def_id": "$$CREDID$$"
			  }
			]
		  }
	  }
	}
}
`
const healthRecord = `
{
	"connection_id": "$$CONNECTIONID$$",
	"comment": "request for proof",
	"proof_request": {
	  "name": "Proof of Education",
	  "version": "1.0",
	  "requested_attributes": {
		"0_name_uuid": {
		  "name": "name",
		  "restrictions": [
			{
			  "cred_def_id": "$$CREDID$$"
			}
		  ]
		},
		"0_place_uuid": {
		  "name": "place",
		  "restrictions": [
			{
			  "cred_def_id": "$$CREDID$$"
			}
		  ]
		}
	}
}
`

const degreeRecord = `
{
	"connection_id": "$$CONNECTIONID$$",
	"comment": "request for proof",
	"proof_request": {
	  "name": "Proof of Education",
	  "version": "1.0",
	  "requested_attributes": {
		"0_name_uuid": {
		  "name": "name",
		  "restrictions": [
			{
			  "cred_def_id": "$$CREDID$$"
			}
		  ]
		},
		"0_completed_date_uuid": {
		  "name": "completed_date",
		  "restrictions": [
			{
			  "cred_def_id": "$$CREDID$$"
			}
		  ]
		}
	}
}
`
