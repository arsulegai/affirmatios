package web

import "encoding/json"

// StructToBytes is an utility that converts a structure into byte array
func StructToBytes(structure interface{}) ([]byte, error) {
	bytes, err := json.Marshal(structure)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
