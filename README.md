# Gambling Services Platform Backend

## Overview

This project is a simple backend system for a gambling services platform. It is designed to utilize multiple communication protocols, including HTTP, WebSockets, and gRPC, for various features. The system primarily manages user wallets, implements real-time updates for game outcomes and leaderboards, and demonstrates the use of multiple communication transports.

## Requirements
### User Wallet API

The user wallet is managed via HTTP APIs. These endpoints allow for deposit, withdrawal, and balance check operations. Here are the API endpoints and their functionalities:

- POST /api/wallet/deposit: Accepts a JSON payload containing the user ID and the amount to be deposited. It updates the user's wallet balance accordingly.

```json
{
  "userId": "user123",
  "amount": 100.0
}
```

- POST /api/wallet/withdraw: Accepts a JSON payload containing the user ID and the amount to be withdrawn. It updates the user's wallet balance if sufficient funds are available.

```json
{
  "userId": "user123",
  "amount": 50.0
}
```

- GET /api/wallet/balance/:user_id: Retrieves the current balance of a user's wallet. Replace :user_id with the actual user ID.

User wallets are managed using an in-memory data store for demonstration purposes.

### Real-time Updates with WebSockets

The backend system also includes a WebSocket server for providing real-time updates to the users. The WebSocket server notifies users about:

Game outcomes
- Leaderboard changes
- Clients can subscribe to specific events like game outcomes and leaderboard changes using WebSocket connections.

### gRPC Endpoint

The system also features a gRPC endpoint that allows users to retrieve their wallet balances. The gRPC service has a single RPC method for retrieving the balance.

## How to Run
```bash
make run
```