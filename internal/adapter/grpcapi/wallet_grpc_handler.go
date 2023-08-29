package grpcapi

import (
	"context"
	"github.com/avag-sargsyan/gambling_platform/internal/conf"
	"github.com/avag-sargsyan/gambling_platform/internal/service"
	"github.com/avag-sargsyan/gambling_platform/proto/walletpb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WalletServiceServer struct {
	walletpb.UnimplementedWalletServiceServer
	walletService service.WalletService
}

func NewWalletServiceServer(configApp *conf.App) *WalletServiceServer {
	return &WalletServiceServer{
		walletService: service.GetWalletServiceInstance(configApp),
	}
}

func (s *WalletServiceServer) GetBalance(ctx context.Context, in *walletpb.BalanceRequest) (*walletpb.BalanceResponse, error) {
	userID := in.GetUserId()

	if userID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "User ID must not be empty")
	}

	balance, err := s.walletService.Balance(userID)
	if err != nil {
		return nil, err
	}
	log.Info().Any("User: ", userID).Send()
	log.Info().Any("Balance: ", balance).Send()

	return &walletpb.BalanceResponse{Balance: balance}, nil
}
