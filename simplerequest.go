// Package that opens urls and returns a Response struct
package simplerequest

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "strings"
)

var Version string = "0.1.0"
var UserAgent string = "go-simplerequest/" + Version
var contentTypeRx = regexp.MustCompile(`/.*`)
var jsonTypeRx = regexp.MustCompile(`(?i)json`)

// Base HTTP response struct
type Response struct {
    Url           string // Original URL that was requested
    FinalUrl      string // Final url after any redirects
    ContentType   string // Content-type
    ContentLength int64  // Content-Length
    StatusCode    int    // Response status code
    Error         string // Any errors that occurred
    Body          string // Body of response
}

// Returns a JSON map of the response body
func (r Response) Json() (map[string]interface{}, error) {
    var j map[string]interface{}

    err := json.Unmarshal([]byte(r.Body), &j)
    if err != nil {
        return j, err
    }

    // Convert keys to lower-case
    json_map := make(map[string]interface{}, len(j))
    for k, v := range j {
        json_map[strings.ToLower(k)] = v
    }

    return json_map, nil
}

// Fetches a url and returns a Response struct
func Get(url string, header map[string]string) (retr Response) {

    // No header was passed in. Create a default header
    if header == nil {
        header = make(map[string]string)
        header["User-Agent"] = UserAgent
    }

    client := &http.Client{}

    retr.Url = url
    retr.FinalUrl = url
    retr.StatusCode = 0

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Sprintf("Error fetching: %s", err)
        retr.Error = fmt.Sprintf("%s", err)
        return retr
    }

    //req.Header.Set("User-Agent", UserAgent)
    for k, v := range header {
        req.Header.Set(k, v)
    }

    resp, err := client.Do(req)
    if err != nil {
        fmt.Sprintf("Error fetching: %s", err)
        retr.Error = fmt.Sprintf("%s", err)
        return retr
    }

    if _, ok := resp.Header["Content-Type"]; ok {

        switch contentTypeRx.ReplaceAllString(resp.Header["Content-Type"][0], "") {
        case "text":
            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                retr.Error = fmt.Sprintf("Error reading body: %s", err)
            } else {
                retr.Body = fmt.Sprintf("%s", body)
            }
        case "application":
            if jsonTypeRx.MatchString(resp.Header["Content-Type"][0]) {
                body, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                    retr.Error = fmt.Sprintf("Error reading body: %s", err)
                } else {
                    retr.Body = fmt.Sprintf("%s", body)
                }
            }
        }
        retr.ContentType = resp.Header["Content-Type"][0]
    }

    retr.ContentLength = resp.ContentLength
    retr.StatusCode = resp.StatusCode
    retr.FinalUrl = resp.Request.URL.String()

    return retr
}
