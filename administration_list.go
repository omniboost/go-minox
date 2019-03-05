package minox

import (
	"net/http"
	"net/url"
)

func (c *Client) NewAdministrationListGetRequest() AdministrationListGetRequest {
	return AdministrationListGetRequest{
		client:      c,
		queryParams: c.NewAdministrationListGetQueryParams(),
		pathParams:  c.NewAdministrationListGetPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewAdministrationListGetRequestBody(),
	}
}

type AdministrationListGetRequest struct {
	client      *Client
	queryParams *AdministrationListGetQueryParams
	pathParams  *AdministrationListGetPathParams
	method      string
	headers     http.Header
	requestBody AdministrationListGetRequestBody
}

func (c *Client) NewAdministrationListGetQueryParams() *AdministrationListGetQueryParams {
	return &AdministrationListGetQueryParams{}
}

type AdministrationListGetQueryParams struct {
	CurrentPage int    `schema:"currentpage,omitempty"`
	Status      string `schema:"status,omitempty"`
}

func (p AdministrationListGetQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AdministrationListGetRequest) QueryParams() *AdministrationListGetQueryParams {
	return r.queryParams
}

func (c *Client) NewAdministrationListGetPathParams() *AdministrationListGetPathParams {
	return &AdministrationListGetPathParams{}
}

type AdministrationListGetPathParams struct {
}

func (p *AdministrationListGetPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AdministrationListGetRequest) PathParams() *AdministrationListGetPathParams {
	return r.pathParams
}

func (r *AdministrationListGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AdministrationListGetRequest) Method() string {
	return r.method
}

func (s *Client) NewAdministrationListGetRequestBody() AdministrationListGetRequestBody {
	return AdministrationListGetRequestBody{}
}

type AdministrationListGetRequestBody struct {
}

func (r *AdministrationListGetRequest) RequestBody() *AdministrationListGetRequestBody {
	return &r.requestBody
}

func (r *AdministrationListGetRequest) SetRequestBody(body AdministrationListGetRequestBody) {
	r.requestBody = body
}

func (r *AdministrationListGetRequest) NewResponseBody() *AdministrationListGetResponseBody {
	return &AdministrationListGetResponseBody{}
}

type AdministrationListGetResponseBody struct {
	Links      CollectionLinks `json:"links"`
	Collection Administrations `json:"collection"`
}

func (r *AdministrationListGetRequest) URL() url.URL {
	return r.client.GetEndpointURL("administration", r.PathParams())
}

func (r *AdministrationListGetRequest) Do() (AdministrationListGetResponseBody, error) {
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
