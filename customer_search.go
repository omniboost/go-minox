package minox

import (
	"net/http"
	"net/url"
)

func (c *Client) NewCustomerSearchRequest() CustomerSearchRequest {
	return CustomerSearchRequest{
		client:      c,
		queryParams: c.NewCustomerSearchQueryParams(),
		pathParams:  c.NewCustomerSearchPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewCustomerSearchRequestBody(),
	}
}

type CustomerSearchRequest struct {
	client      *Client
	queryParams *CustomerSearchQueryParams
	pathParams  *CustomerSearchPathParams
	method      string
	headers     http.Header
	requestBody CustomerSearchRequestBody
}

func (c *Client) NewCustomerSearchQueryParams() *CustomerSearchQueryParams {
	return &CustomerSearchQueryParams{}
}

type CustomerSearchQueryParams struct {
	Next   string `schema:"next,omitempty"`
	Search string `schema:"search,omitempty"`
}

func (p CustomerSearchQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerSearchRequest) QueryParams() *CustomerSearchQueryParams {
	return r.queryParams
}

func (c *Client) NewCustomerSearchPathParams() *CustomerSearchPathParams {
	return &CustomerSearchPathParams{}
}

type CustomerSearchPathParams struct {
	AdministrationID string
}

func (p *CustomerSearchPathParams) Params() map[string]string {
	return map[string]string{
		"administration_id": p.AdministrationID,
	}
}

func (r *CustomerSearchRequest) PathParams() *CustomerSearchPathParams {
	return r.pathParams
}

func (r *CustomerSearchRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerSearchRequest) Method() string {
	return r.method
}

func (s *Client) NewCustomerSearchRequestBody() CustomerSearchRequestBody {
	return CustomerSearchRequestBody{}
}

type CustomerSearchRequestBody struct {
}

func (r *CustomerSearchRequest) RequestBody() *CustomerSearchRequestBody {
	return &r.requestBody
}

func (r *CustomerSearchRequest) SetRequestBody(body CustomerSearchRequestBody) {
	r.requestBody = body
}

func (r *CustomerSearchRequest) NewResponseBody() *CustomerSearchResponseBody {
	return &CustomerSearchResponseBody{}
}

type CustomerSearchResponseBody struct {
	Links      CollectionLinks `json:"links"`
	Collection Administrations `json:"collection"`
}

func (r *CustomerSearchRequest) URL() url.URL {
	return r.client.GetEndpointURL("administration/{{.administration_id}}/customer", r.PathParams())
}

func (r *CustomerSearchRequest) Do() (CustomerSearchResponseBody, error) {
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
