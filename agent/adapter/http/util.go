package http

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"jonatann-silva/network/dsn/structs"
)

const (
	paginationPageQueryKey    = "page"
	paginationPerPageQueryKey = "per_page"
	defaultPaginationPage     = 1
	defaultPaginationPerPage  = 10
)

func parsePaginationQueryParams(query url.Values) (int, int) {

	var err error
	var page, perPage int

	if page, err = strconv.Atoi(query.Get(paginationPageQueryKey)); err != nil {
		page = defaultPaginationPage
	}

	if perPage, err = strconv.Atoi(query.Get(paginationPerPageQueryKey)); err != nil {
		perPage = defaultPaginationPerPage
	}

	return page, perPage
}

func parseBody(body io.ReadCloser, out interface{}) error {
	defer body.Close()

	encoded, err := ioutil.ReadAll(body)
	if err != nil {
		return fmt.Errorf("error reading request body")
	}

	if err := json.Unmarshal(encoded, out); err != nil {
		return fmt.Errorf("error decoding request body")
	}

	return nil
}

func parsePathParams(req *http.Request) []string {
	path := req.URL.Path
	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")
	params := strings.Split(path, "/")
	return params
}

func parseAuthToken(req *http.Request) string {
	return req.Header.Get("X-DSN-Token")
}

func trimPathPrefix(req *http.Request, prefix string) *http.Request {
	s := strings.TrimSuffix(strings.TrimPrefix(req.URL.Path, "/"), "/")
	s = strings.TrimPrefix(s, prefix)
	req.URL.Path = s
	return req
}

func parseQueryOptions(req *http.Request) structs.QueryOptions {
	return structs.QueryOptions{
		AuthToken: parseAuthToken(req),
		Filters:   parseFilters(req),
	}
}

func parseFilters(req *http.Request) structs.Filters {
	return structs.Filters(req.URL.Query())
}

func parseWriteRequestOptions(req *http.Request) structs.WriteRequest {
	return structs.WriteRequest{
		AuthToken: parseAuthToken(req),
	}
}

func parseError(err error) error {
	switch err.Error() {
	case structs.ErrPermissionDenied.Error():
		return NewCodedError(403, err.Error())
	case structs.ErrNotFound.Error():
		return NewCodedError(404, err.Error())
	default:
		return NewCodedError(500, err.Error())
	}
}
