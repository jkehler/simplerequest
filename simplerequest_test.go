package simplerequest

import (
    "testing"
)

func Test200(t *testing.T) {
    resp := Get("http://httpbin.org/get", nil)
    if resp.StatusCode != 200 {
        t.Error("Expected 200, got ", resp.StatusCode)
    }
}

func Test404(t *testing.T) {
    resp := Get("http://httpbin.org/status/404", nil)
    if resp.StatusCode != 404 {
        t.Error("Expected 404, got ", resp.StatusCode)
    }
}

func Test301(t *testing.T) {
    resp := Get("http://httpbin.org/status/301", nil)
    if resp.StatusCode != 200 {
        t.Error("Expected 200, got ", resp.StatusCode)
    }
    if resp.FinalUrl != "http://httpbin.org/get" {
        t.Error("Expected redirect to httpbin.org/get, got ", resp.FinalUrl)
    }
    if resp.Url != "http://httpbin.org/status/301" {
        t.Error("Expected original url to be http://httpbin.org/status/301, got ",
            resp.Url)
    }
}

func TestMultipleRedirect(t *testing.T) {
    resp := Get("http://httpbin.org/redirect/6", nil)
    if resp.StatusCode != 200 {
        t.Error("Expected 200, got ", resp.StatusCode)
    }
    if resp.FinalUrl != "http://httpbin.org/get" {
        t.Error("Expected redirect to httpbin.org/get, got ", resp.FinalUrl)
    }
    if resp.Url != "http://httpbin.org/redirect/6" {
        t.Error("Expected original url to be http://httpbin.org/redirect/6, got ",
            resp.Url)
    }
}

func TestBadRequest(t *testing.T) {
    resp := Get("failed", nil)
    if resp.StatusCode != 0 {
        t.Error("Expected 0 StatusCode, got ", resp.StatusCode)
    }
    if resp.Error == "" {
        t.Error("Expected an Error string")
    }
}

func TestHeaders(t *testing.T) {
    // Test default header
    resp := Get("http://httpbin.org/headers", nil)
    json_resp, err := resp.Json()
    if err != nil {
        t.Error("Error extracting JSON response ", resp.Error)
    }
    if headers, ok := json_resp["headers"]; ok {
        if agent, ok := headers.(map[string]interface{})["User-Agent"]; ok {
            if agent != "go-simplerequest/"+Version {
                t.Errorf("Expected go-simplerequest/%s, got %s", Version, agent)
            }
        }
    } else {

        t.Error("user-agent response not found")
    }

    // Test custom header
    header := make(map[string]string)
    header["User-Agent"] = "test-agent"
    header["Cache-Control"] = "no-cache, no-store, must-revalidate"
    header["Pragma"] = "no-cache"
    header["Expires"] = "0"

    resp = Get("http://httpbin.org/headers", header)
    json_resp, err = resp.Json()
    if err != nil {
        t.Error("Error extracting JSON response ", resp.Error)
    }
    if headers, ok := json_resp["headers"]; ok {
        if agent, ok := headers.(map[string]interface{})["User-Agent"]; ok {
            if agent != "test-agent" {
                t.Error("Expected test-agent, got ", agent)
            }
        }
    } else {
        t.Error("user-agent response not found")
    }
}
