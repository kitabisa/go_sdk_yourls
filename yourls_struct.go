package yourls

import (
	"net/http"
	"net/url"
)

// Options ...
type Options struct {
	httpClient *http.Client
	baseURL    *url.URL
	Token      string
	Username   string
	Password   string
}

// OutputFormat ...
type OutputFormat string

// Action ...
type Action string

// Filter ...
type Filter string

const (

	// FormKeyTagName ...
	FormKeyTagName string = "form_key"

	// Json ...
	Json OutputFormat = "json"

	// ShortURL ...
	ShortURL Action = "shorturl"

	// Expand ...
	Expand Action = "expand"

	// Stats ...
	Stats Action = "stats"

	// URLStats ...
	URLStats Action = "url-stats"

	// DBStats ...
	DBStats Action = "db-stats"

	// Top ...
	Top Filter = "top"

	// Bottom ...
	Bottom Filter = "bottom"

	// Rand ...
	Rand Filter = "rand"

	// Last ...
	Last Filter = "last"
)

// FormatAndSignature ...
type FormatAndSignature struct {
	Format    OutputFormat `form_key:"format"`
	Signature string       `form_key:"signature"`
	UserName  string       `form_key:"username"`
	Password  string       `form_key:"password"`
}

// ActionShortURLRequest ...
type ActionShortURLRequest struct {
	URL     string `form_key:"url"`
	Keyword string `form_key:"keyword"`
	Title   string `form_key:"title"`
	Action  Action `form_key:"action"`
}

// ActionExpandRequest ...
type ActionExpandRequest struct {
	ShortURL string `form_key:"shorturl"`
	Action   Action `form_key:"action"`
}

// ActionURLStatsRequest ...
type ActionURLStatsRequest struct {
	ShortURL string `form_key:"shorturl"`
	Action   Action `form_key:"action"`
}

// ActionDBStatsRequest ...
type ActionDBStatsRequest struct {
	Action Action `form_key:"action"`
}

// ActionStatsRequest ...
type ActionStatsRequest struct {
	Filter Filter `form_key:"filter"`
	Limit  int    `form_key:"limit"`
	Action Action `form_key:"action"`
}

// ActionShortURL ...
type ActionShortURL struct {
	Keyword   string `json:"keyword,omitempty"`
	URL       string `json:"url,omitempty"`
	ShortURL  string `json:"shorturl,omitempty"`
	Title     string `json:"title,omitempty"`
	Date      string `json:"date,omitempty"`
	TimeStamp string `json:"timestamp,omitempty"`
	IP        string `json:"ip,omitempty"`
	Clicks    string `json:"clicks,omitempty"`
}

// GeneralErrorResponse ...
type GeneralErrorResponse struct {
	ErrorCode int    `json:"errorCode,omitempty"`
	Message   string `json:"message,omitempty"`
	Keyword   bool   `json:"keyword,omitempty"`
}

// ActionUnknownResponse ...
type ActionUnknownResponse struct {
	Status     string `json:"status,omitempty"`
	Code       string `json:"code,omitempty"`
	ErrorCode  string `json:"errorCode,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
}

// ActionShortURLResponse ...
type ActionShortURLResponse struct {
	URL        ActionShortURL `json:"url,omitempty"`
	Status     string         `json:"status,omitempty"`
	Code       string         `json:"code,omitempty"`
	ErrorCode  string         `json:"errorCode,omitempty"`
	Message    string         `json:"message,omitempty"`
	Title      string         `json:"title,omitempty"`
	ShortURL   string         `json:"shorturl,omitempty"`
	StatusCode int            `json:"statusCode,omitempty"`
}

// ActionExpandResponse ...
type ActionExpandResponse struct {
	Keyword    string `json:"keyword,omitempty"`
	ShortURL   string `json:"shorturl,omitempty"`
	LongURL    string `json:"longurl,omitempty"`
	Title      string `json:"title,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"statusCode,omitempty"`
}

// ActionURLStatsResponse ...
type ActionURLStatsResponse struct {
	StatusCode int            `json:"statusCode,omitempty"`
	Message    string         `json:"message,omitempty"`
	Link       ActionShortURL `json:"link,omitempty"`
}

// ActionStats ...
type ActionStats struct {
	TotalLinks  string `json:"total_links,omitempty"`
	TotalClicks string `json:"total_clicks,omitempty"`
}

// ActionStatsSimpleResponse ...
type ActionStatsSimpleResponse struct {
	Stats      ActionStats `json:"stats,omitempty"`
	Message    string      `json:"message,omitempty"`
	StatusCode int         `json:"statusCode,omitempty"`
}

// ActionStatsFullResponse ...
type ActionStatsFullResponse struct {
	Links      map[string]ActionShortURL `json:"links,omitempty"`
	Stats      ActionStats               `json:"stats,omitempty"`
	Message    string                    `json:"message,omitempty"`
	StatusCode int                       `json:"statusCode,omitempty"`
}

// ActionDBStatsSimpleResponse ...
type ActionDBStatsSimpleResponse struct {
	Stats      ActionStats `json:"db-stats,omitempty"`
	Message    string      `json:"message,omitempty"`
	StatusCode int         `json:"statusCode,omitempty"`
}
