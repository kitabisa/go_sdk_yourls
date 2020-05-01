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

type OutputFormat string
type Action string
type Filter string

const (
	FormKeyTagName string       = "form_key"
	Json           OutputFormat = "json"
	ShortURL       Action       = "shorturl"
	Expand         Action       = "expand"
	Stats          Action       = "stats"
	URLStats       Action       = "url-stats"
	DBStats        Action       = "db-stats"
	Top            Filter       = "top"
	Bottom         Filter       = "bottom"
	Rand           Filter       = "rand"
	Last           Filter       = "last"
)

// ActionShortURLRequest ...
type ActionShortURLRequest struct {
	URL       string       `form_key:"url"`
	Keyword   string       `form_key:"keyword"`
	Title     string       `form_key:"title"`
	Action    Action       `form_key:"action"`
	Format    OutputFormat `form_key:"format"`
	Signature string       `form_key:"signature"`
	UserName  string       `form_key:"username"`
	Password  string       `form_key:"password"`
}

// ActionExpandRequest ...
type ActionExpandRequest struct {
	ShortURL  string       `form_key:"shorturl"`
	Action    Action       `form_key:"action"`
	Format    OutputFormat `form_key:"format"`
	Signature string       `form_key:"signature"`
	UserName  string       `form_key:"username"`
	Password  string       `form_key:"password"`
}

// ActionURLStatsRequest ...
type ActionURLStatsRequest struct {
	ShortURL  string       `form_key:"shorturl"`
	Action    Action       `form_key:"action"`
	Format    OutputFormat `form_key:"format"`
	Signature string       `form_key:"signature"`
	UserName  string       `form_key:"username"`
	Password  string       `form_key:"password"`
}

// ActionDBStatsRequest ...
type ActionDBStatsRequest struct {
	Action    Action       `form_key:"action"`
	Format    OutputFormat `form_key:"format"`
	Signature string       `form_key:"signature"`
	UserName  string       `form_key:"username"`
	Password  string       `form_key:"password"`
}

// ActionStatsRequest ...
type ActionStatsRequest struct {
	Filter    Filter       `form_key:"filter"`
	Limit     int          `form_key:"limit"`
	Action    Action       `form_key:"action"`
	Format    OutputFormat `form_key:"format"`
	Signature string       `form_key:"signature"`
	UserName  string       `form_key:"username"`
	Password  string       `form_key:"password"`
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

// ActionShortURLResponse ...
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

// ActionStatsSimpleResponse ...
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
