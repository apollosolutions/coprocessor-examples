package com.example.demo.models;

import java.util.LinkedHashMap;

public class RouterRequestBody {
    private String query;
    private String operationName;
    private LinkedHashMap<String, Object> variables;

    public String getQuery() {
        return this.query;
    }

    public void setQuery(String query) {
        this.query = query;
    }

    public String getOperationName() {
        return this.operationName;
    }

    public void setOperationName(String operationName) {
        this.operationName = operationName;
    }

    public LinkedHashMap<String, Object> getVariables() {
        return this.variables;
    }

    public void setVariables(LinkedHashMap<String, Object> variables) {
        this.variables = variables;
    }
}
