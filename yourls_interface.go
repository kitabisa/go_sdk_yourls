package yourls

import "net/http"

// BuildYourls ...
type BuildYourls interface {
	setOptions(yourlsOptions Options) BuildYourls
	SendAction(request interface{}) (returnData interface{}, errorResponse interface{}, resp *http.Response, err error)
}
