package minox

import (
	"encoding/json"
	"net/http"
	"net/url"

	uuid "github.com/satori/go.uuid"
)

func (c *Client) NewTransactionLineBatchPutRequest() TransactionLineBatchPutRequest {
	return TransactionLineBatchPutRequest{
		client:      c,
		queryParams: c.NewTransactionLineBatchPutQueryParams(),
		pathParams:  c.NewTransactionLineBatchPutPathParams(),
		method:      http.MethodPut,
		headers:     http.Header{},
		requestBody: c.NewTransactionLineBatchPutRequestBody(),
	}
}

type TransactionLineBatchPutRequest struct {
	client      *Client
	queryParams *TransactionLineBatchPutQueryParams
	pathParams  *TransactionLineBatchPutPathParams
	method      string
	headers     http.Header
	requestBody TransactionLineBatchPutRequestBody
}

func (c *Client) NewTransactionLineBatchPutQueryParams() *TransactionLineBatchPutQueryParams {
	return &TransactionLineBatchPutQueryParams{}
}

type TransactionLineBatchPutQueryParams struct {
	CurrentPage int    `schema:"currentpage,omitempty"`
	Status      string `schema:"status,omitempty"`
}

func (r *TransactionLineBatchPutRequest) Headers() http.Header {
	return r.headers
}

func (p TransactionLineBatchPutQueryParams) ToURLValues() (url.Values, error) {
	encoder := newSchemaEncoder()
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *TransactionLineBatchPutRequest) QueryParams() *TransactionLineBatchPutQueryParams {
	return r.queryParams
}

func (c *Client) NewTransactionLineBatchPutPathParams() *TransactionLineBatchPutPathParams {
	return &TransactionLineBatchPutPathParams{}
}

type TransactionLineBatchPutPathParams struct {
	AdministrationID string
	BatchID          uuid.UUID
}

func (p *TransactionLineBatchPutPathParams) Params() map[string]string {
	return map[string]string{
		"administration_id": p.AdministrationID,
		"batch_id":          p.BatchID.String(),
	}
}

func (r *TransactionLineBatchPutRequest) PathParams() *TransactionLineBatchPutPathParams {
	return r.pathParams
}

func (r *TransactionLineBatchPutRequest) SetMethod(method string) {
	r.method = method
}

func (r *TransactionLineBatchPutRequest) Method() string {
	return r.method
}

func (s *Client) NewTransactionLineBatchPutRequestBody() TransactionLineBatchPutRequestBody {
	return TransactionLineBatchPutRequestBody{}
}

type TransactionLineBatchPutRequestBody struct {
	TransactionLineBatchPut
}

func (r TransactionLineBatchPutRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.TransactionLineBatchPut)
}

func (r *TransactionLineBatchPutRequest) RequestBody() *TransactionLineBatchPutRequestBody {
	return &r.requestBody
}

func (r *TransactionLineBatchPutRequest) SetRequestBody(body TransactionLineBatchPutRequestBody) {
	r.requestBody = body
}

func (r *TransactionLineBatchPutRequest) NewResponseBody() *TransactionLineBatchPutResponseBody {
	return &TransactionLineBatchPutResponseBody{}
}

type TransactionLineBatchPutResponseBody struct {
}

func (r *TransactionLineBatchPutRequest) URL() url.URL {
	return r.client.GetEndpointURL("administration/{{.administration_id}}/transaction_line_batch/{{.batch_id}}", r.PathParams())
}

func (r *TransactionLineBatchPutRequest) Do() (TransactionLineBatchPutResponseBody, error) {
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
