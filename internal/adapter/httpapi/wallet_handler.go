package httpapi

import (
	"encoding/json"
	"errors"
	"github.com/avag-sargsyan/gambling_platform/internal/conf"
	"github.com/avag-sargsyan/gambling_platform/internal/dto"
	"github.com/avag-sargsyan/gambling_platform/internal/service"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"net/http"
)

type WalletHandler interface {
	Deposit(w http.ResponseWriter, r *http.Request)
	Withdraw(w http.ResponseWriter, r *http.Request)
	Balance(w http.ResponseWriter, r *http.Request)
}

type walletHandler struct {
	walletService service.WalletService
}

// Should have context also
func NewWalletHandler(configApp *conf.App) WalletHandler {
	return &walletHandler{
		walletService: service.GetWalletServiceInstance(configApp),
	}
}

func (h *walletHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	var wallet dto.Wallet
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&wallet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate input
	if wallet.UserID == "" || wallet.Balance <= 0 {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}
	log.Info().Any("Deposit Request: ", wallet).Send()

	// Ideally should be better tuned errors, maybe define custom errors in service and check here with error.is
	err := h.walletService.Deposit(&wallet)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *walletHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	var wallet dto.Wallet
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&wallet); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	// Validate input
	if wallet.UserID == "" || wallet.Balance <= 0 {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}
	log.Info().Any("Withdraw Request: ", wallet).Send()

	err := h.walletService.Withdraw(&wallet)
	if err != nil {
		if errors.Is(err, service.WalletServiceInsufficientFunds) {
			http.Error(w, "Insufficient funds", http.StatusBadRequest)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *walletHandler) Balance(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID := r.URL.Query().Get("user_id")

	// Validate input
	if userID == "" {
		http.Error(w, "Invalid or missing user ID", http.StatusBadRequest)
		return
	}

	val, err := h.walletService.Balance(userID)
	log.Info().Any("User: ", userID).Send()
	log.Info().Any("Balance: ", val).Send()

	if err != nil && err != redis.Nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	wallet := dto.Wallet{
		UserID:  userID,
		Balance: val,
	}

	json.NewEncoder(w).Encode(wallet)
}
