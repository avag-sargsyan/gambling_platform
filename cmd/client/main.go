package main

import (
	"context"
	"github.com/avag-sargsyan/gambling_platform/internal/conf"
	"github.com/avag-sargsyan/gambling_platform/internal/logger"
	"github.com/avag-sargsyan/gambling_platform/internal/util"
	"github.com/avag-sargsyan/gambling_platform/proto/walletpb"
	"github.com/caarlos0/env/v9"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	logger.Setup()
	configApp := conf.App{}
	err := env.Parse(&configApp)
	util.FatalIfError(err)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(configApp.GRPcAddressClient, opts...)
	if err != nil {
		util.FatalIfError(err)
	}
	defer conn.Close()

	// Initialize the WalletService client
	client := walletpb.NewWalletServiceClient(conn)

	// Timeout after 5 seconds if the server doesn't respond
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Call the GetBalance method and receive a response
	response, err := client.GetBalance(ctx, &walletpb.BalanceRequest{UserId: "user123"})
	util.FatalIfError(err)

	log.Info().Msgf("Received response: %v", response)
}
