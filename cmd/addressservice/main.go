package main

import (
	"context"

	"github.com/anz-bank/addressservice/internal/gen/pkg/servers/addressservice"
	"github.com/anz-bank/addressservice/internal/handlers"
	"github.com/anz-bank/sysl-go/core"
)

type AppConfig struct {
	// Define app-level config fields here.
}

func main() {
	addressservice.Serve(context.Background(),
		func(ctx context.Context, config AppConfig) (*addressservice.ServiceInterface, *core.Hooks, error) {
			// Perform one-time setup based on config here.
			return &addressservice.ServiceInterface{
				// Add handlers here.
				GetAddressCheckList: handlers.GetAddressCheckList,
			}, nil, nil
		},
	)
}
