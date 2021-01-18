package aagent

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func (a *AriesAgent) post(message []byte, path string) ([]byte, error) {
	res, err := http.Post("http://"+a.Host+":"+a.Port+path, "application/json", bytes.NewReader(message))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}
