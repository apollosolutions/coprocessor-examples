package com.example.demo;

import com.example.demo.models.RouterPayload;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.example.demo.handlers.RouterRequestHandler;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class CoprocessorController {
    private final RouterRequestHandler routerRequestHandler;

    public CoprocessorController(RouterRequestHandler routerRequestHandler) {
        this.routerRequestHandler = routerRequestHandler;
    }

    @PostMapping("/")
    public RouterPayload StageSelect(@RequestBody RouterPayload request) {
        

        if (request.getStage().equals("RouterRequest")) {
            return routerRequestHandler.handle(request);
        }
        
        return request;
    }
}
