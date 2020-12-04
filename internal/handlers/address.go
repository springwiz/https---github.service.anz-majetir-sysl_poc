package handlers

import (
	"context"

	"github.com/anz-bank/addressservice/internal/gen/pkg/servers/addressservice"
	"github.com/anz-bank/addressservice/internal/gen/pkg/servers/addressservice/addressdirectory"
	"github.com/anz-bank/sysl-go/common"
)

// GetAddressCheckList
func GetAddressCheckList(ctx context.Context, req *addressservice.GetAddressCheckListRequest, client addressservice.GetAddressCheckListClient) (*addressservice.Address, error) {

	// Set response content type to JSON
	headers := common.RequestHeaderFromContext(ctx)
	headers["Content-Type"] = []string{"application/json; charset=utf-8"}

	reqAddress := addressdirectory.GetAddressListRequest{}
	address, err := client.AddressdirectoryGetAddressList(ctx, &reqAddress)
	if err != nil {
		return nil, err
	}

	return &addressservice.Address{
		Country: address.Country,
		PinCode: address.PinCode,
		Street:  address.Street,
	}, nil
}
