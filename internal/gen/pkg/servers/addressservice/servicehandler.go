// Code generated by sysl DO NOT EDIT.
package addressservice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/config"
	"github.com/anz-bank/sysl-go/core"
	"github.com/anz-bank/sysl-go/core/authrules"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"

	"github.com/anz-bank/addressservice/internal/gen/pkg/servers/addressservice/addressdirectory"
)

// Handler interface for addressservice
type Handler interface {
	GetAddressCheckListHandler(w http.ResponseWriter, r *http.Request)
}

// ServiceHandler for addressservice API
type ServiceHandler struct {
	genCallback             core.RestGenCallback
	serviceInterface        *ServiceInterface
	authorizationRules      map[string]authrules.Rule
	addressdirectoryService addressdirectory.Service
}

// NewServiceHandler for addressservice
func NewServiceHandler(
	ctx context.Context,
	cfg *config.DefaultConfig,
	hooks *core.Hooks,
	genCallback core.RestGenCallback,
	serviceInterface *ServiceInterface,
	addressdirectoryAddressdirectoryService addressdirectory.Service,
) (*ServiceHandler, error) {

	authorizationRules := make(map[string]authrules.Rule)

	return &ServiceHandler{
		genCallback,
		serviceInterface,
		authorizationRules,
		addressdirectoryAddressdirectoryService,
	}, nil
}

// GetAddressCheckListHandler ...
func (s *ServiceHandler) GetAddressCheckListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetAddressCheckList == nil {
		common.HandleError(r.Context(), w, common.InternalError, "not implemented", nil, s.genCallback.MapError)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetAddressCheckListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		common.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr, s.genCallback.MapError)
		return
	}

	client := GetAddressCheckListClient{
		AddressdirectoryGetAddressList: s.addressdirectoryService.GetAddressList,
	}

	defer func() {
		if rec := recover(); rec != nil {
			var err error
			switch rec := rec.(type) {
			case error:
				err = rec
			default:
				err = fmt.Errorf("Unknown error: %v", rec)
			}
			common.HandleError(ctx, w, common.InternalError, "Unexpected panic", err, s.genCallback.MapError)
		}
	}()
	address, err := s.serviceInterface.GetAddressCheckList(ctx, &req, client)
	if err != nil {
		common.HandleError(ctx, w, common.InternalError, "Handler error", err, s.genCallback.MapError)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	if headermap.Get("Content-Type") == "" {
		headermap.Set("Content-Type", "application/json")
	}
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, address)
}