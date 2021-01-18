package aagent

import (
	"io/ioutil"
	"net/http"
)

func (a *AriesAgent) get(path string) ([]byte, error) {
	res, err := http.Get("http://" + a.Host + ":" + a.Port + path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}
