package com.example.demo.handlers;

import com.example.demo.models.RouterPayload;
import com.example.demo.models.RouterRequestBody;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class RouterRequestHandler {

    @Autowired
    ObjectMapper mapper;

    public RouterPayload handle(RouterPayload request) {

        // Need to convert to a serialized JSON object since it's coming in as an escaped string, which doesn't seem to be deserializable into a POJO
        String stringifiedNode = request.getBody().asText();

        try {
            // Since the request/response body changes depending on where in the transaction we are I've stored it in a generic JsonNode
            // Since it would need to be processed differently at the RouterRequest stage vs the SupergraphRequest stage for example
            JsonNode node = mapper.readTree(stringifiedNode);
            RouterRequestBody body = mapper.treeToValue(node, RouterRequestBody.class);

            // Adding the variables as stringified JSON to context allows you to pass them to any other part of the request
            // For example you could pass this context variable to an OTel span to log out what is being passed by the client
            // I do want to note though that when passing variable values as in this sample, to exercise extreme caution as
            // variables could contain PII (personally identifiable information)
            String serializedVariables = mapper.writeValueAsString(body.getVariables());
            request.getContext().addEntry("variableValues", serializedVariables);

            System.out.println(request.getContext().getEntries());
        } catch (JsonProcessingException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }

        return request;
    }
}
