package service

import (
	"context"
	"errors"
	"github.com/avag-sargsyan/gambling_platform/internal/conf"
	"github.com/avag-sargsyan/gambling_platform/internal/database"
	"github.com/avag-sargsyan/gambling_platform/internal/dto"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"sync"
)

var (
	once                           sync.Once
	singleton                      WalletService
	WalletServiceInsufficientFunds = errors.New("insufficient funds")
)

type WalletService interface {
	Deposit(wallet *dto.Wallet) error
	Withdraw(wallet *dto.Wallet) error
	Balance(userID string) (float64, error)
}

type walletServiceImpl struct {
	rdb *redis.Client
	ctx context.Context
}

func newWalletService(configApp *conf.App) WalletService {
	ctx := context.Background()
	rdb := database.GetConnection(configApp)
	return &walletServiceImpl{rdb: rdb, ctx: ctx}
}

func GetWalletServiceInstance(configApp *conf.App) WalletService {
	once.Do(func() {
		singleton = newWalletService(configApp)
	})
	return singleton
}

func (s *walletServiceImpl) Deposit(wallet *dto.Wallet) error {
	// Ideally should be handled in repository, skipping to save time
	val, err := s.rdb.Get(s.ctx, wallet.UserID).Float64()
	if err != nil && err != redis.Nil {
		log.Error().Err(err).Send()
		return err
	}

	newBalance := val + wallet.Balance
	log.Info().Any("New Balance should be: ", newBalance).Send()
	err = s.rdb.Set(s.ctx, wallet.UserID, newBalance, 0).Err()
	return err
}

func (s *walletServiceImpl) Withdraw(wallet *dto.Wallet) error {
	// Ideally should be handled in repository, skipping to save time
	val, err := s.rdb.Get(s.ctx, wallet.UserID).Float64()
	if err != nil && err != redis.Nil {
		log.Error().Err(err).Send()
		return err
	}

	if val < wallet.Balance {
		return WalletServiceInsufficientFunds
	}

	newBalance := val - wallet.Balance
	log.Info().Any("New Balance should be: ", newBalance).Send()
	err = s.rdb.Set(s.ctx, wallet.UserID, newBalance, 0).Err()
	return err
}

func (s *walletServiceImpl) Balance(userID string) (float64, error) {
	// Ideally should be handled in repository, skipping to save time
	balance, err := s.rdb.Get(s.ctx, userID).Float64()
	if err != nil && err != redis.Nil {
		log.Error().Err(err).Send()
		return 0, err
	}
	return balance, nil
}
