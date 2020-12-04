// Code generated by sysl DO NOT EDIT.
package addressdirectory

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
)

// Service interface for addressdirectory
type Service interface {
	GetAddressList(ctx context.Context, req *GetAddressListRequest) (*Address, error)
}

// Client for addressdirectory API
type Client struct {
	Client  *http.Client
	URL     string
	Headers map[string][]string
}

// NewClient for addressdirectory
func NewClient(client *http.Client, serviceURL string) *Client {
	return &Client{client, serviceURL, nil}
}

// GetAddressList ...
func (s *Client) GetAddressList(ctx context.Context, req *GetAddressListRequest) (*Address, error) {
	required := []string{}
	var okResponse Address
	u, err := url.Parse(fmt.Sprintf("%s/address", s.URL))
	if err != nil {
		return nil, common.CreateError(ctx, common.InternalError, "failed to parse url", err)
	}

	result, err := restlib.DoHTTPRequest2(ctx, &restlib.HTTPRequest{
		Client:        s.Client,
		Method:        "GET",
		URLString:     u.String(),
		Body:          nil,
		Required:      required,
		OKResponse:    &okResponse,
		ErrorResponse: nil,
		ExtraHeaders:  nil,
	})
	restlib.OnRestResultHTTPResult(ctx, result, err)
	if err != nil {
		return nil, common.CreateError(ctx, common.DownstreamUnavailableError, "call failed: addressdirectory <- GET "+u.String(), err)
	}

	if result.HTTPResponse.StatusCode == http.StatusUnauthorized {
		return nil, common.CreateDownstreamError(ctx, common.DownstreamUnauthorizedError, result.HTTPResponse, result.Body, nil)
	}
	OkAddressResponse, ok := result.Response.(*Address)
	if ok {
		valErr := validator.Validate(OkAddressResponse)
		if valErr != nil {
			return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, valErr)
		}

		return OkAddressResponse, nil
	}
	return nil, common.CreateDownstreamError(ctx, common.DownstreamUnexpectedResponseError, result.HTTPResponse, result.Body, nil)
}
