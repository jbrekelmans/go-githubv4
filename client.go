package githubv4

import (
	"context"
	"net/http"

	"github.com/jbrekelmans/go-graphql"
)

type Input interface {
	isInput()
}

// Client is a GitHub GraphQL v4 client.
type Client struct {
	c *graphql.Client
}

// NewClient constructs a client for https://api.github.com/graphql.
// The *http.Client should add credentials/tokens to requests.
func NewClient(httpClient *http.Client) *Client {
	return &Client{
		c: graphql.NewClient("https://api.github.com/graphql", httpClient),
	}
}

// NewEnterpriseClient constructs a client for the specified GitHub GraphQL v4 endpoint.
// The *http.Client should add credentials/tokens to requests.
func NewEnterpriseClient(url string, httpClient *http.Client) *Client {
	return &Client{
		c: graphql.NewClient(url, httpClient),
	}
}

// Query does a query operation.
// q is a pointer to a struct that defines the GraphQL query, and also receives the response data.
//
// If the HTTP response status and headers were received successfully then returns a non-nil *http.Response that reflects the status and
// headers. The body of the returned HTTP response is always closed.
//
// If the GraphQL response was completely received and parsed, and contains GraphQL-level errors, an error of type *Error is returned
// that reflects the GraphQL-level errors.
//
// Users should interpret any of the following as a transient error condition that may go away (after some time) by retrying:
// 1. The returned error, say x, implements interface{ Temporary() bool } and x.Temporary() is true.
// 2. The returned error wraps an error, say x, that implements interface{ Temporary() bool } and x.Temporary() is true.
// 3. The returned error, say x, implements interface{ Timeout() bool } and x.Timeout() is true.
// 4. The returned error wraps an error, say x, that implements interface{ Timeout() bool } and x.Timeout() is true.
// 5. The returned response has status code between 500 and 599.
// 6. The returned response has status code 429.
//
// 1, 2, 3 and 4 are known to be produced in the event of transient errors and timeouts by
//   - the *http.Client;
//   - the http.RoundTripper of the *http.Client when dialing, sending the HTTP request and reading
//     the HTTP response; -and
//   - the underlying connnection when reading the HTTP response body.
func (c *Client) Query(ctx context.Context, q any, variables map[string]any) (*http.Response, error) {
	resp, err := c.c.Query(ctx, q, variables)
	err = enhanceError(err)
	return resp, err
}

// Mutate does a mutation operation.
// m is a pointer to a struct that defines the GraphQL mutation, and also receives the response data.
//
// See Query for more information.
func (c *Client) Mutate(ctx context.Context, m any, input Input, variables map[string]any) (*http.Response, error) {
	if input != nil {
		if variables == nil {
			variables = map[string]any{}
		}
		variables["input"] = input
	}
	resp, err := c.c.Mutate(ctx, m, variables)
	err = enhanceError(err)
	return resp, err
}
