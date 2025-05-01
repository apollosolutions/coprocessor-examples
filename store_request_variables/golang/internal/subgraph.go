package coprocessor

import (
	"encoding/json"
	"fmt"
)

type SubgraphBody struct {
	*Body
	Variables map[string]any `json:"variables,omitempty"`
}

type SubgraphRequest struct {
	// Control properties
	*CommonProperties
	Control any `json:"control,omitempty"`

	// Data properties
	*Headers
	Body        *SubgraphBody `json:"body,omitempty"`
	Context     *Context      `json:"context,omitempty"`
	URI         string        `json:"uri,omitempty"`
	Method      string        `json:"method,omitempty"`
	ServiceName string        `json:"serviceName,omitempty"`
}

type SubgraphResponse struct {
	// Control properties
	*CommonProperties
	Control any `json:"control,omitempty"`

	// Data properties
	*Headers
	Body        *Body    `json:"body,omitempty"`
	Context     *Context `json:"context,omitempty"`
	URI         string   `json:"uri,omitempty"`
	ServiceName string   `json:"serviceName,omitempty"`
	StatusCode  float64  `json:"statusCode,omitempty"`
}

func handleSubgraphRequest(httpRequestBody *[]byte) (*SubgraphRequest, error) {
	cr, err := NewCoprocessorSubgraphRequest(httpRequestBody)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling httpRequestBody: %w", err)
	}

	return cr, nil
}

func handleSubgraphResponse(httpRequestBody *[]byte) (*SubgraphResponse, error) {
	cr, err := NewCoprocessorSubgraphResponse(httpRequestBody)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling httpRequestBody: %w", err)
	}
	return cr, nil
}

func NewCoprocessorSubgraphRequest(httpRequestBody *[]byte) (*SubgraphRequest, error) {
	var err error
	var cr *SubgraphRequest
	err = json.Unmarshal(*httpRequestBody, &cr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return cr, nil
}

func NewCoprocessorSubgraphResponse(httpRequestBody *[]byte) (*SubgraphResponse, error) {
	var err error
	var cr *SubgraphResponse
	err = json.Unmarshal(*httpRequestBody, &cr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return cr, nil
}
