package com.example.demo.models;

import java.util.LinkedHashMap;

public class Context {
    private LinkedHashMap<String, Object> entries;

    public Context() {}

    public LinkedHashMap<String, Object> getEntries() {
        return this.entries;
    }

    protected void setEntries(LinkedHashMap<String, Object> entries) {
        this.entries = entries;
    }

    public void addEntry(String key, Object value) {
        this.entries.put(key, value);
    }
}
