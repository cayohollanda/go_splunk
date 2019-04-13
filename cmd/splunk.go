package go_splunk

// SplunkConnection : guard data of connection with Splunk API
type SplunkConnection struct {

	// Username to connect on Splunk (usually is your login username)
	Username string

	// Password to connect on Splunk (usually is your logi password of your username)
	Password string

	// URL to connect with Splunk API. Ex.: https://localhost:8089/
	APIURL string
}

// SearchResult : have result of a search that has been sended
// to Splunk, having EventsCounter that is a counter to how many
// events has returned on request and Raw that is a []string
// of returned raws of events
type SearchResult struct {
	EventsCounter int64
	Raw           []string
}

// GetSearchResults : is a function to return a SearchResult of one search
// sended to Splunk API. This function receives search param that is a string of
// search
func (conn SplunkConnection) GetSearchResults(search string) (result *SearchResult, err error) {
	// needs to be implemented

	return &SearchResult{}, nil
}
