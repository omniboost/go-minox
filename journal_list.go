package minox

import (
	"net/http"
	"net/url"
)

func (c *Client) NewJournalListGetRequest() JournalListGetRequest {
	return JournalListGetRequest{
		client:      c,
		queryParams: c.NewJournalListGetQueryParams(),
		pathParams:  c.NewJournalListGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewJournalListGetRequestBody(),
	}
}

type JournalListGetRequest struct {
	client      *Client
	queryParams *JournalListGetQueryParams
	pathParams  *JournalListGetPathParams
	method      string
	headers     http.Header
	requestBody JournalListGetRequestBody
}

func (c *Client) NewJournalListGetQueryParams() *JournalListGetQueryParams {
	return &JournalListGetQueryParams{}
}

type JournalListGetQueryParams struct {
	CurrentPage int    `schema:"currentpage,omitempty"`
	Status      string `schema:"status,omitempty"`
}

func (p JournalListGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *JournalListGetRequest) QueryParams() *JournalListGetQueryParams {
	return r.queryParams
}

func (c *Client) NewJournalListGetPathParams() *JournalListGetPathParams {
	return &JournalListGetPathParams{}
}

type JournalListGetPathParams struct {
	AdministrationID string
}

func (p *JournalListGetPathParams) Params() map[string]string {
	return map[string]string{
		"administration_id": p.AdministrationID,
	}
}

func (r *JournalListGetRequest) PathParams() *JournalListGetPathParams {
	return r.pathParams
}

func (r *JournalListGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalListGetRequest) Method() string {
	return r.method
}

func (s *Client) NewJournalListGetRequestBody() JournalListGetRequestBody {
	return JournalListGetRequestBody{}
}

type JournalListGetRequestBody struct {
}

func (r *JournalListGetRequest) RequestBody() *JournalListGetRequestBody {
	return &r.requestBody
}

func (r *JournalListGetRequest) SetRequestBody(body JournalListGetRequestBody) {
	r.requestBody = body
}

func (r *JournalListGetRequest) NewResponseBody() *JournalListGetResponseBody {
	return &JournalListGetResponseBody{}
}

type JournalListGetResponseBody struct {
	Links      CollectionLinks `json:"links"`
	Collection Journals        `json:"collection"`
}

func (r *JournalListGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("administration/{{.administration_id}}/journal", r.PathParams())
}

func (r *JournalListGetRequest) Do() (JournalListGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r.Method(), r.URL(), r.RequestBody())
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
