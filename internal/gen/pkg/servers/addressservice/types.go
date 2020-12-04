// Code generated by sysl DO NOT EDIT.
package addressservice

import (
	"github.com/anz-bank/sysl-go/validator"
)

// Address ...
type Address struct {
	Country string `json:"country"`
	PinCode string `json:"pinCode"`
	Street  string `json:"street"`
}

// GetAddressCheckListRequest ...
type GetAddressCheckListRequest struct {
}

// *Address validator
func (s *Address) Validate() error {
	return validator.Validate(s)
}