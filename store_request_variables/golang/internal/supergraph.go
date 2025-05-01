package coprocessor

import (
	"encoding/json"
	"fmt"
)

type SupergraphRequest struct {
	// Control properties
	// minus "id" according to the docs
	*CommonProperties
	Control any `json:"control,omitempty"`

	// Data properties
	*Headers
	Body        *Body    `json:"body,omitempty"`
	Context     *Context `json:"context,omitempty"`
	ServiceName string   `json:"serviceName,omitempty"`
	URI         string   `json:"uri,omitempty"`
}

type SupergraphResponse struct {
	// Control properties
	*CommonProperties
	Control any `json:"control,omitempty"`

	// Data properties
	Body    *Body    `json:"body,omitempty"`
	Context *Context `json:"context,omitempty"`
	*Headers
}

func handleSupergraphRequest(httpRequestBody *[]byte) (*SupergraphRequest, error) {
	cr, err := NewCoprocessorSupergraphRequest(httpRequestBody)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling httpRequestBody: %w", err)
	}
	return cr, nil
}

func handleSupergraphResponse(httpRequestBody *[]byte) (*SupergraphResponse, error) {
	cr, err := NewCoprocessorSupergraphResponse(httpRequestBody)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling httpRequestBody: %w", err)
	}
	return cr, nil
}

func NewCoprocessorSupergraphRequest(httpRequestBody *[]byte) (*SupergraphRequest, error) {
	var err error
	var cr *SupergraphRequest
	err = json.Unmarshal(*httpRequestBody, &cr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return cr, nil
}

func NewCoprocessorSupergraphResponse(httpRequestBody *[]byte) (*SupergraphResponse, error) {
	var err error
	var cr *SupergraphResponse
	err = json.Unmarshal(*httpRequestBody, &cr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return cr, nil
}
