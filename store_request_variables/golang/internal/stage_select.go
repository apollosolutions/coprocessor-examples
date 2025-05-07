package coprocessor

import (
	"fmt"
	"log"
	"net/http"
)

func HandleRequest(w http.ResponseWriter, req *http.Request) (interface{}, error) {
	// TODO: Instead of unmarshalling twice, unmarshal once for the whole request
	httpRequestBody, stage, err := NewRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling httpRequestBody: %w", err)
	}

	log.Printf("Stating stage: %v", stage)

	switch stage {
	case "RouterRequest":
		return handleRouterRequest(httpRequestBody)
	case "RouterResponse":
		return handleRouterResponse(httpRequestBody)
	case "SupergraphRequest":
		return handleSupergraphRequest(httpRequestBody)
	case "SupergraphResponse":
		return handleSupergraphResponse(httpRequestBody)
	case "SubgraphRequest":
		return handleSubgraphRequest(httpRequestBody)
	case "SubgraphResponse":
		return handleSubgraphResponse(httpRequestBody)
	case "":
		// This shouldn't happen, everything should have a Stage
		return nil, fmt.Errorf("no stage for request %+v", httpRequestBody)
	default:
		return nil, fmt.Errorf("unhandled coprocessor request stage of type: %T", stage)
	}
}
