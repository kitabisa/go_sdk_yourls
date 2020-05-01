package yourls

import (
	"net/http"
	"net/url"
)

// Builder ...
type Builder struct {
	option  Options
	builder BuildYourls
}

// SetBuilder ...
func (c *Builder) SetBuilder(b BuildYourls) {
	c.builder = b
}

// SetHTTPClient ...
func (c *Builder) SetHTTPClient(httpClient *http.Client) {
	c.option.httpClient = httpClient
}

// SetBaseURL ...
func (c *Builder) SetBaseURL(baseURL *url.URL) {
	c.option.baseURL = baseURL
}

// SetToken ...
func (c *Builder) SetToken(token string) {
	c.option.Token = token
}

// SetUsername ...
func (c *Builder) SetUsername(username string) {
	c.option.Username = username
}

// SetPassword ...
func (c *Builder) SetPassword(password string) {
	c.option.Password = password
}

// Build ...
func (c *Builder) Build() BuildYourls {
	c.builder.setOptions(c.option)
	return c.builder
}
