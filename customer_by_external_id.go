package minox

import (
	"net/http"
	"net/url"
)

func (c *Client) NewCustomerByExternalIDRequest() CustomerByExternalIDRequest {
	return CustomerByExternalIDRequest{
		client:      c,
		queryParams: c.NewCustomerByExternalIDQueryParams(),
		pathParams:  c.NewCustomerByExternalIDPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewCustomerByExternalIDRequestBody(),
	}
}

type CustomerByExternalIDRequest struct {
	client      *Client
	queryParams *CustomerByExternalIDQueryParams
	pathParams  *CustomerByExternalIDPathParams
	method      string
	headers     http.Header
	requestBody CustomerByExternalIDRequestBody
}

func (c *Client) NewCustomerByExternalIDQueryParams() *CustomerByExternalIDQueryParams {
	return &CustomerByExternalIDQueryParams{}
}

type CustomerByExternalIDQueryParams struct {
}

func (p CustomerByExternalIDQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerByExternalIDRequest) QueryParams() *CustomerByExternalIDQueryParams {
	return r.queryParams
}

func (c *Client) NewCustomerByExternalIDPathParams() *CustomerByExternalIDPathParams {
	return &CustomerByExternalIDPathParams{}
}

type CustomerByExternalIDPathParams struct {
	AdministrationID string
	ExternalID string
}

func (p *CustomerByExternalIDPathParams) Params() map[string]string {
	return map[string]string{
		"administration_id": p.AdministrationID,
		"externalid": p.ExternalID,
	}
}

func (r *CustomerByExternalIDRequest) PathParams() *CustomerByExternalIDPathParams {
	return r.pathParams
}

func (r *CustomerByExternalIDRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerByExternalIDRequest) Method() string {
	return r.method
}

func (s *Client) NewCustomerByExternalIDRequestBody() CustomerByExternalIDRequestBody {
	return CustomerByExternalIDRequestBody{}
}

type CustomerByExternalIDRequestBody struct {
}

func (r *CustomerByExternalIDRequest) RequestBody() *CustomerByExternalIDRequestBody {
	return &r.requestBody
}

func (r *CustomerByExternalIDRequest) SetRequestBody(body CustomerByExternalIDRequestBody) {
	r.requestBody = body
}

func (r *CustomerByExternalIDRequest) NewResponseBody() *CustomerByExternalIDResponseBody {
	return &CustomerByExternalIDResponseBody{}
}

type CustomerByExternalIDResponseBody struct {
	Links      CollectionLinks `json:"links"`
	Collection []CustomerPost  `json:"collection"`
}

func (r *CustomerByExternalIDRequest) URL() url.URL {
	return r.client.GetEndpointURL("administration/{{.administration_id}}/customer", r.PathParams())
}

func (r *CustomerByExternalIDRequest) Do() (CustomerByExternalIDResponseBody, error) {
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

