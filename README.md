# Polymarket Leaders

A Go application that monitors Polymarket blockchain activity on Polygon and tracks the top 10 most active addresses based on trading and position events.

## Overview

This service continuously monitors the Polygon blockchain for Polymarket-related events, extracts user addresses from trading activities, and maintains a leaderboard of the most active participants. It provides both a web dashboard and a JSON API to view the current leaders.

## Features

- **Real-time Monitoring**: Continuously tracks Polymarket events on Polygon
- **Event Tracking**: Monitors multiple contract types:
    - CTFExchange (OrderFilled, OrdersMatched)
    - NegRiskCtfExchange (OrderFilled, OrdersMatched)
    - ConditionalTokens (PositionSplit, PositionsMerge)
    - NegRiskAdapter (PositionSplit, PositionsMerge, PositionsConverted)
- **Leaderboard**: Tracks top 10 most active addresses
- **Web Dashboard**: HTML dashboard showing current leaders with statistics
- **JSON API**: RESTful API endpoint for programmatic access
- **Docker Support**: Containerized deployment with Docker and Docker Compose

## Prerequisites

- Go 1.24.0 or later
- Access to a Polygon RPC endpoint (public or private)
- Docker and Docker Compose (optional, for containerized deployment)

## Installation

### From Source

1. Clone the repository:
```bash
git clone https://github.com/gladcow/polymarket_leaders.git
cd polymarket_leaders
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o polymarket_leaders ./main.go
```

### Using Docker

1. Build the Docker image:
```bash
docker build -t polymarket_leaders:latest .
```

2. Or use Docker Compose (see Configuration section below)

## Configuration

The application can be configured via environment variables or a `.env` file. Configuration options:

| Variable | Default | Description |
|----------|---------|-------------|
| `RPC_URL` | `https://polygon-rpc.com` | Polygon RPC endpoint URL |
| `CHAIN_ID` | `137` | Polygon chain ID |
| `REQUEST_INTERVAL` | `2s` | Interval between block checks (e.g., "2s", "5s") |
| `DASHBOARD_ADDRESS` | `:8080` | HTTP server address for dashboard |

### Environment File

Create a `.env` file in the project root:

```env
RPC_URL=https://polygon-rpc.com
CHAIN_ID=137
REQUEST_INTERVAL=2s
DASHBOARD_ADDRESS=:8080
```

### Docker Compose

For Docker Compose deployment, create a `.env.production` file:

```env
RPC_URL=https://your-polygon-rpc-endpoint.com
CHAIN_ID=137
REQUEST_INTERVAL=2s
DASHBOARD_ADDRESS=:8080
```

Then start the service:
```bash
docker-compose up -d
```

## Usage

### Running Locally

```bash
./polymarket_leaders
```

Or with a custom config path:
```bash
./polymarket_leaders -config /path/to/config.yaml
```

### Running with Docker

```bash
docker run -d \
  -p 8080:8080 \
  -e RPC_URL=https://polygon-rpc.com \
  -e CHAIN_ID=137 \
  -e REQUEST_INTERVAL=2s \
  -e DASHBOARD_ADDRESS=:8080 \
  polymarket_leaders:latest
```

### Accessing the Dashboard

Once running, access the dashboard at:
- **Web UI**: http://localhost:8080
- **JSON API**: http://localhost:8080/api/leaders

## API Endpoints

### GET `/api/leaders`

Returns the current leaderboard as JSON.

**Response:**
```json
{
  "service": "polymarket-activity-leaders",
  "startBlock": 12345678,
  "lastBlock": 12345690,
  "addresses": [
    {
      "address": "0x...",
      "count": 42
    }
  ]
}
```

## Architecture

The application consists of three main components:

1. **Polymarket Service** (`internal/polymarket/`): Monitors blockchain events and tracks address activity
2. **Dashboard Service** (`internal/dashboard/`): Serves the web UI and JSON API
3. **Config** (`internal/config/`): Handles configuration loading from environment variables

### Event Processing

The service processes blocks sequentially, extracting events from Polymarket contracts and counting occurrences per address. The top 10 addresses are recalculated every 10 blocks.

### Contract Bindings

The application includes Go bindings for various Polymarket contracts:
- ConditionalTokens
- CTFExchange
- NegRiskCtfExchange
- NegRiskAdapter
- UmaCtfAdapter
- NegRiskFeeModule
- ProxyWalletFactory
- SafeProxyFactory
- UChildERC20Proxy
- UmaCtfAdapter2


## License

This project is licensed under the GNU General Public License v3.0 (GPL-3.0).

See the [LICENSE](LICENSE) file for details.

## Support

For issues and questions, please open an issue on the GitHub repository.
