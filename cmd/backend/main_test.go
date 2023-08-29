package main

import (
	"context"
	"github.com/avag-sargsyan/gambling_platform/internal/adapter/grpcapi"
	"github.com/avag-sargsyan/gambling_platform/proto/walletpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

// Not for test coverage, just for testing GRPC calls
func TestGetBalance(t *testing.T) {
	const bufSize = 1024 * 1024
	lis := bufconn.Listen(bufSize)
	s := grpc.NewServer()
	walletpb.RegisterWalletServiceServer(s, &grpcapi.WalletServiceServer{})

	go func() {
		if err := s.Serve(lis); err != nil {
			t.Fatalf("Server exited with error: %v", err)
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := walletpb.NewWalletServiceClient(conn)

	resp, err := client.GetBalance(context.Background(), &walletpb.BalanceRequest{UserId: "1"})
	if err != nil {
		t.Fatalf("GetBalance failed: %v", err)
	}

	if resp.Balance != 123.0 {
		t.Fatalf("Expected balance 100.0, got %v", resp.Balance)
	}
}
