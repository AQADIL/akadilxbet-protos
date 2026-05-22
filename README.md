# AkadilXbet Protobuf Definitions

This repository contains the Protocol Buffer (`.proto`) definitions and generated Go code for the AkadilXbet microservices architecture.

## Purpose

This repository serves as the **single source of truth** for all gRPC service contracts across the AkadilXbet microservices. All services import these proto definitions to ensure type safety and API compatibility.

## Services

- **Auth Service** (`proto/auth/auth.proto`) - User authentication and profile management
- **Wallet Service** (`proto/wallet/wallet.proto`) - Wallet balance and transaction management
- **Game Engine Service** (`proto/game_engine/game_engine.proto`) - Game outcome calculation

## Usage

### Adding as a dependency

In your Go service's `go.mod`:

```bash
go get github.com/AQADIL/akadilxbet-protos
```

### Importing in code

```go
import (
    "github.com/AQADIL/akadilxbet-protos/pb/auth"
    "github.com/AQADIL/akadilxbet-protos/pb/wallet"
    "github.com/AQADIL/akadilxbet-protos/pb/game_engine"
)
```

### Example: Auth Service

```go
req := &auth.RegisterRequest{
    Email:    "user@example.com",
    Password: "securepassword",
}
```

## Project Structure

```
.
├── proto/
│   ├── auth/
│   │   └── auth.proto
│   ├── wallet/
│   │   └── wallet.proto
│   └── game_engine/
│       └── game_engine.proto
├── pb/
│   ├── auth/
│   │   └── auth.pb.go
│   ├── wallet/
│   │   └── wallet.pb.go
│   └── game_engine/
│       └── game_engine.pb.go
├── go.mod
└── README.md
```

## Adding New Proto Files

1. Create or modify `.proto` file in the appropriate `proto/` subdirectory
2. Update the `go_package` option to point to this repository:
   ```protobuf
   option go_package = "github.com/AQADIL/akadilxbet-protos/pb/{service}";
   ```
3. Generate the corresponding `.pb.go` file in the `pb/` directory
4. Commit and push changes

## Versioning

This repository follows semantic versioning. When making breaking changes to proto definitions:
1. Update the version tag (e.g., `v1.0.0` → `v2.0.0`)
2. Update dependent services to use the new version

## License

Internal use only for AkadilXbet project.
