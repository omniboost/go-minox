package minox

import (
	"net/http"
	"net/url"

	uuid "github.com/satori/go.uuid"
)

func (c *Client) NewCustomerPostRequest() CustomerPostRequest {
	return CustomerPostRequest{
		client:      c,
		queryParams: c.NewCustomerPostQueryParams(),
		pathParams:  c.NewCustomerPostPathParams(),
		method:      http.MethodPost,
		headers:     http.Header{},
		requestBody: c.NewCustomerPostRequestBody(),
	}
}

type CustomerPostRequest struct {
	client      *Client
	queryParams *CustomerPostQueryParams
	pathParams  *CustomerPostPathParams
	method      string
	headers     http.Header
	requestBody CustomerPostRequestBody
}

func (c *Client) NewCustomerPostQueryParams() *CustomerPostQueryParams {
	return &CustomerPostQueryParams{}
}

type CustomerPostQueryParams struct {
	JournalID string `schema:"journal_id,omitempty"`
}

func (r *CustomerPostRequest) Headers() http.Header {
	return r.headers
}

func (p CustomerPostQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CustomerPostRequest) QueryParams() *CustomerPostQueryParams {
	return r.queryParams
}

func (c *Client) NewCustomerPostPathParams() *CustomerPostPathParams {
	return &CustomerPostPathParams{}
}

type CustomerPostPathParams struct {
	AdministrationID string
	BatchID          uuid.UUID
}

func (p *CustomerPostPathParams) Params() map[string]string {
	return map[string]string{
		"administration_id": p.AdministrationID,
		"batch_id":          p.BatchID.String(),
	}
}

func (r *CustomerPostRequest) PathParams() *CustomerPostPathParams {
	return r.pathParams
}

func (r *CustomerPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerPostRequest) Method() string {
	return r.method
}

func (s *Client) NewCustomerPostRequestBody() CustomerPostRequestBody {
	return CustomerPostRequestBody{}
}

type CustomerPostRequestBody CustomerPost

func (r *CustomerPostRequest) RequestBody() *CustomerPostRequestBody {
	return &r.requestBody
}

func (r *CustomerPostRequest) SetRequestBody(body CustomerPostRequestBody) {
	r.requestBody = body
}

func (r *CustomerPostRequest) NewResponseBody() *CustomerPostResponseBody {
	return &CustomerPostResponseBody{}
}

type CustomerPostResponseBody struct {
}

func (r *CustomerPostRequest) URL() url.URL {
	return r.client.GetEndpointURL("administration/{{.administration_id}}/customer", r.PathParams())
}

func (r *CustomerPostRequest) Do() (CustomerPostResponseBody, error) {
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
