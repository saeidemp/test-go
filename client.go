package ghasedakapi

import(
	"net/http"
)

const(
	baseUrl="http://api.parsasms.com/ping"
)
type Client struct {
	ApiKey string
	BaseUrl    string
	HTTPClient *http.Client
}

// Create a new ghasedakapi struct.
func NewClient(apiKey string) *ghasedakapi {
	return NewHttpClient(apiKey, nil)
}

// Create a new ghasedakapi client, using a http.Client
func NewHttpClient(apiKey string, HTTPClient *http.Client) *Client {
	if HTTPClient == nil {
		HTTPClient = http.DefaultClient
	}
	return &ghasedakapi{apiKey,baseUrl, HTTPClient}
}


func (c *Client) Execute(urlStr string, b url.Values, v interface{}) error {
	body := strings.NewReader(b.Encode())
	ul, _ := url.Parse(urlStr)
	u := c.BaseURL.ResolveReference(ul)
	req, _ := http.NewRequest("POST", u.String(), body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Charset", "utf-8")
	resp, err := c.BaseClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_ = json.NewDecoder(resp.Body).Decode(&v)
	return nil
}
