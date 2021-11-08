package rest_fiber

import (
	"encoding/json"
	"io"
	"io/ioutil"

	response "idaman.id/storage/internal/rest-response"
)

func StringifyResponse(r io.Reader) string {
	body, _ := ioutil.ReadAll(r)
	resBody := string(body)

	return resBody
}

func UnmarshallResponseBody(r io.Reader) *response.ResponseEntity {
	resBody := StringifyResponse(r)

	var resEntity *response.ResponseEntity
	json.Unmarshal([]byte(resBody), &resEntity)

	return resEntity
}
