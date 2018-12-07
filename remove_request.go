package authy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RemoveRequest is the approval request response.
type RemoveRequest struct {
	HTTPResponse *http.Response
}

// NewRemoveRequest returns an instance of RemoveRequest.
func NewRemoveRequest(response *http.Response) (*RemoveRequest, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	jsonResponse := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return nil, err
	}

	if !jsonResponse.Success {
		return nil, fmt.Errorf("invalid remove request response: %s", jsonResponse.Message)
	}
	var removeRequest RemoveRequest
	removeRequest.HTTPResponse = response

	return &removeRequest, nil
}

// Valid returns true if the remove request was valid.
func (request *RemoveRequest) Valid() bool {
	return request.HTTPResponse.StatusCode == 200
}
