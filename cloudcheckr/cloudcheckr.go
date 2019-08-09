package cloudcheckr

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"

	"github.com/google/go-querystring/query"
	"github.com/moul/http2curl"
)

const (
	baseURL      = "https://api.cloudcheckr.com/api/"
	inventoryURL = "https://api.cloudcheckr.com/api/inventory.json/"
)

// Parameters for Requests
type Parameters struct {
	Date            string `url:"date,omitempty"`
	AwsAccountIDs   string `url:"aws_account_ids,omitempty"`
	InstanceIDS     string `url:"instance_ids,omitempty"`
	ResourceTags    string `url:"resource_tags,omitempty"`
	UseAccount      string `url:"use_account,omitempty"`
	UseCCAccountID  string `url:"use_cc_account_id,omitempty"`
	UseAWSAccountID string `url:"use_aws_account_id,omitempty"`
	NextToken       string `url:"next_token,omitempty"`
}

func (p *Parameters) SetDate(date string) {
	p.Date = date
}

func (p *Parameters) SetAwsAccountIDs(awsaccountids string) {
	p.AwsAccountIDs = awsaccountids
}

func (p *Parameters) SetInstanceIDS(instanceids string) {
	p.InstanceIDS = instanceids
}

func (p *Parameters) SetResourceTags(resourcetags string) {
	p.ResourceTags = resourcetags
}

func (p *Parameters) SetUseAccount(useaccount string) {
	p.UseAccount = useaccount
}

func (p *Parameters) SetUseCCAccountID(useccaccountid string) {
	p.UseCCAccountID = useccaccountid
}

func (p *Parameters) SetUseAWSAccountID(useawsaccountid string) {
	p.UseAWSAccountID = useawsaccountid
}

// Client represents CloudCheckr HTTP client
type Client struct {
	m *sync.Mutex

	Token        string
	BaseURL      string
	InventoryURL string
	Debug        bool
	Client       *http.Client
}

// NewClient creates new CloudCheckr client
func NewClient(client *http.Client, token string) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	return &Client{
		Token:        token,
		BaseURL:      baseURL,
		InventoryURL: inventoryURL,
		Client:       client,
		m:            &sync.Mutex{},
	}
}

// NewEnvClient creates new CloudCheckr client using environment variable
// CC_API as the token.
func NewEnvClient(client *http.Client, envname string) *Client {
	if envname == "" {
		envname = "CC_API"
	}
	return NewClient(client, os.Getenv(envname))
}

func (c *Client) PrintDebug(message interface{}) {
	if c.Debug == true {
		fmt.Printf("[ DEBUG ] %v\n", message)
	}
}

// SetDebug toggles the debug mode.
func (c *Client) SetDebug(debug bool) {
	c.m.Lock()
	defer c.m.Unlock()

	c.Debug = debug
}

// NewRequest prepares new request to common CloudCheckr api.
func (c *Client) NewRequest(method string, path string, params interface{}, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}

	return c.baseRequest(method, u, params, body)
}

// NewRequest prepares new request to common CloudCheckr api.
func (c *Client) NewInventoryRequest(method string, path string, params interface{}, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(c.InventoryURL + path)
	if err != nil {
		return nil, err
	}

	return c.baseRequest(method, u, params, body)
}

func (c *Client) baseRequest(method string, u *url.URL, params interface{}, body io.Reader) (*http.Request, error) {
	qs, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	qs.Add("access_key", c.Token)

	u.RawQuery = qs.Encode()

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	return req, nil
}

// Do executes common (non-streaming) request.
func (c *Client) Do(ctx context.Context, req *http.Request, destination interface{}) error {
	// resp, err := c.do(ctx, req)
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if c.Debug {
		if command, err := http2curl.GetCurlCommand(req); err == nil {
			log.Printf("CloudCheckr client request: %s\n", command)
		}
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil
	}

	if destination == nil {
		return nil
	}

	return c.parseResponse(destination, resp.Body)
}

func (c *Client) parseResponse(destination interface{}, body io.Reader) error {
	var err error

	if w, ok := destination.(io.Writer); ok {
		_, err = io.Copy(w, body)
	} else {
		decoder := json.NewDecoder(body)
		err = decoder.Decode(destination)
	}

	return err
}
