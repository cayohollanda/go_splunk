package go_splunk

import (
	"crypto/tls"
	"net/http"
)

// HTTPClient : this function is used to return a http client with ssl check skip
func HTTPClient() *http.Client {
	transp := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transp}

	return client
}

// HTTPGetRequest : this function returns response to a GET request
// receiving a URL parameter on function
func (conn SplunkConnection) HTTPGetRequest(URL string) (response *http.Response, err error) {
	// setting client to do a request
	c := HTTPClient()

	// creating request and checking if have error
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return &http.Response{}, err
	}

	// setting authentication of request
	req.SetBasicAuth(conn.Username, conn.Password)

	// sending request and checking if have error
	resp, err := c.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	// returning response
	return resp, nil
}
