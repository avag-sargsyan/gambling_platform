package main

import (
	"github.com/avag-sargsyan/gambling_platform/internal/adapter/grpcapi"
	"github.com/avag-sargsyan/gambling_platform/internal/adapter/httpapi"
	"github.com/avag-sargsyan/gambling_platform/internal/adapter/wsapi"
	"github.com/avag-sargsyan/gambling_platform/internal/conf"
	"github.com/avag-sargsyan/gambling_platform/internal/logger"
	"github.com/avag-sargsyan/gambling_platform/internal/util"
	"github.com/avag-sargsyan/gambling_platform/proto/walletpb"
	"github.com/caarlos0/env/v9"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func main() {
	logger.Setup()

	configApp := conf.App{}
	err := env.Parse(&configApp)
	util.FatalIfError(err)

	initHTTPServer(&configApp)
	initGRPcServer(&configApp)
}

// These functions should be moved to internal/adapter/server.go
func initGRPcServer(configApp *conf.App) {
	grpcServer := grpc.NewServer()
	walletpb.RegisterWalletServiceServer(grpcServer, grpcapi.NewWalletServiceServer(configApp)) // Ideally should be interface
	go func() {
		listen, err := net.Listen("tcp", configApp.GRPcAddress)
		util.FatalIfError(err)
		util.FatalIfError(grpcServer.Serve(listen))
	}()
}

func initHTTPServer(configApp *conf.App) {
	setUpHTTPRoutes(configApp)
	util.FatalIfError(http.ListenAndServe(configApp.HTTPAddress, nil))
}

func setUpHTTPRoutes(configApp *conf.App) {
	walletHandler := httpapi.NewWalletHandler(configApp)
	http.HandleFunc("/api/wallet/deposit", walletHandler.Deposit)
	http.HandleFunc("/api/wallet/withdraw", walletHandler.Withdraw)
	http.HandleFunc("/api/wallet/balance", walletHandler.Balance)
	http.HandleFunc("/websocket", wsapi.WebsocketHandler) // Should be used interface for loose coupling
}
