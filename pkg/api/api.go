package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	contractbindv1 "github.com/ic-n/ethcontractbind/pkg/api/gen/contractbind/v1"
)

type Server struct {
	log    *slog.Logger
	auth   *bind.TransactOpts
	client bind.ContractBackend
	contractbindv1.UnimplementedContractbindServiceServer
}

func NewServer(
	log *slog.Logger,
	auth *bind.TransactOpts,
	client bind.ContractBackend,
) *Server {
	return &Server{
		log:    log,
		auth:   auth,
		client: client,
	}
}

func NewHandler(ctx context.Context, s *Server) (http.Handler, error) {
	mux := runtime.NewServeMux()
	if err := contractbindv1.RegisterContractbindServiceHandlerServer(ctx, mux, s); err != nil {
		return nil, fmt.Errorf("failed to register service: %w", err)
	}

	return mux, nil
}

func (s Server) Health(ctx context.Context, _ *contractbindv1.HealthRequest) (*contractbindv1.HealthResponse, error) {
	if _, err := s.client.HeaderByNumber(ctx, nil); err != nil {
		err = fmt.Errorf("blockchain is not online: %w", err)
		s.log.Warn(err.Error())
		return nil, err
	}

	return &contractbindv1.HealthResponse{
		Ok: true,
	}, nil
}

// TODO: implement API
