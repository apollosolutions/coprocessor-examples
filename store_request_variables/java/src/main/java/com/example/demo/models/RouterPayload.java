package com.example.demo.models;

import java.util.ArrayList;
import java.util.LinkedHashMap;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonInclude(JsonInclude.Include.NON_ABSENT)
public class RouterPayload {
    private JsonNode control;
    private LinkedHashMap<String, ArrayList<String>> headers;
    private String id;
    private String method;
    private String sdl;
    private String stage;
    private Integer version;
    private JsonNode body;
    private Context context;
    private String path;
    private String uri;
    private String serviceName;
    private String subgraphRequestId;
    private Integer statusCode;


    public JsonNode getControl() {
        return this.control;
    }

    public void setControl(JsonNode control) {
        this.control = control;
    }

    public LinkedHashMap<String, ArrayList<String>> getHeaders() {
        return this.headers;
    }

    public void setHeaders(LinkedHashMap<String, ArrayList<String>> headers) {
        this.headers = headers;
    }

    public String getId() {
        return this.id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getMethod() {
        return this.method;
    }

    public void setMethod(String method) {
        this.method = method;
    }

    public String getSDL() {
        return this.sdl;
    }

    public void setSDL(String sdl) {
        this.sdl = sdl;
    }

    public String getStage() {
        return this.stage;
    }

    public void setStage(String stage) {
        this.stage = stage;
    }

    public Integer getVersion() {
        return this.version;
    }

    public void setVersion(Integer version) {
        this.version = version;
    }

    public JsonNode getBody() {
        return this.body;
    }

    public void setBody(JsonNode body) {
        this.body = body;
    }

    public Context getContext() {
        return this.context;
    }

    public String getPath() {
        return this.path;
    }

    public void setPath(String path) {
        this.path = path;
    }

    // TODO: Think about if I need to expose a setter for context
    protected void setContext(Context context) {
        this.context = context;
    }


    public String getUri() {
        return this.uri;
    }

    public void setUri(String uri) {
        this.uri = uri;
    }


    public String getServiceName() {
        return this.serviceName;
    }

    public void setServiceName(String serviceName) {
        this.serviceName = serviceName;
    }


    public String getSubgraphRequestId() {
        return this.subgraphRequestId;
    }

    public void setSubgraphRequestId(String subgraphRequestId) {
        this.subgraphRequestId = subgraphRequestId;
    }

    public Integer getStatusCode() {
        return this.statusCode;
    }

    public void setStatusCode(Integer statusCode) {
        this.statusCode = statusCode;
    }
}
