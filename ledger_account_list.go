package minox

import (
	"net/http"
	"net/url"
)

func (c *Client) NewLedgerAccountListGetRequest() LedgerAccountListGetRequest {
	return LedgerAccountListGetRequest{
		client:      c,
		queryParams: c.NewLedgerAccountListGetQueryParams(),
		pathParams:  c.NewLedgerAccountListGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewLedgerAccountListGetRequestBody(),
	}
}

type LedgerAccountListGetRequest struct {
	client      *Client
	queryParams *LedgerAccountListGetQueryParams
	pathParams  *LedgerAccountListGetPathParams
	method      string
	headers     http.Header
	requestBody LedgerAccountListGetRequestBody
}

func (c *Client) NewLedgerAccountListGetQueryParams() *LedgerAccountListGetQueryParams {
	return &LedgerAccountListGetQueryParams{}
}

type LedgerAccountListGetQueryParams struct {
	CurrentPage int    `schema:"currentpage,omitempty"`
	Status      string `schema:"status,omitempty"`
}

func (p LedgerAccountListGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *LedgerAccountListGetRequest) QueryParams() *LedgerAccountListGetQueryParams {
	return r.queryParams
}

func (c *Client) NewLedgerAccountListGetPathParams() *LedgerAccountListGetPathParams {
	return &LedgerAccountListGetPathParams{}
}

type LedgerAccountListGetPathParams struct {
	AdministrationID string
}

func (p *LedgerAccountListGetPathParams) Params() map[string]string {
	return map[string]string{
		"administration_id": p.AdministrationID,
	}
}

func (r *LedgerAccountListGetRequest) PathParams() *LedgerAccountListGetPathParams {
	return r.pathParams
}

func (r *LedgerAccountListGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *LedgerAccountListGetRequest) Method() string {
	return r.method
}

func (s *Client) NewLedgerAccountListGetRequestBody() LedgerAccountListGetRequestBody {
	return LedgerAccountListGetRequestBody{}
}

type LedgerAccountListGetRequestBody struct {
}

func (r *LedgerAccountListGetRequest) RequestBody() *LedgerAccountListGetRequestBody {
	return &r.requestBody
}

func (r *LedgerAccountListGetRequest) SetRequestBody(body LedgerAccountListGetRequestBody) {
	r.requestBody = body
}

func (r *LedgerAccountListGetRequest) NewResponseBody() *LedgerAccountListGetResponseBody {
	return &LedgerAccountListGetResponseBody{}
}

type LedgerAccountListGetResponseBody struct {
	Links      CollectionLinks `json:"links"`
	Collection LedgerAccounts  `json:"collection"`
}

func (r *LedgerAccountListGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("administration/{{.administration_id}}/ledger_account", r.PathParams())
}

func (r *LedgerAccountListGetRequest) Do() (LedgerAccountListGetResponseBody, error) {
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
