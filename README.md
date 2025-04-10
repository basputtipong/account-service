# Account-Service

## Build & Run

### Build Docker Container
```bash
make docker-build
```

### Start Docker Container
```bash
make docker-up
```

---

## Local Development

- Run the service locally:
```bash
make run
```

- The service will be available at:
  http://localhost:1400

---

## API Details

### GET /account – Get Account Data

**Request Header:**
```
Authorization: Bearer {authToken}
```

**Request Body:**
```json
{}
```

**Response:**
```json
{
  "accounts": [
    {
      "accountId": "000020ece1a211ef95a30242ac180002",
      "type": "saving-account",
      "currency": "THB",
      "accountNumber": "568-2-62729",
      "issuer": "TestLab",
      "amount": 86048.06,
      "color": "#24c875",
      "isMainAccount": true,
      "progress": 42,
      "flags": [
        {
          "flagId": 3810097,
          "flagType": "system",
          "flagValue": "Flag4"
        },
        {
          "flagId": 3810098,
          "flagType": "system",
          "flagValue": "Flag5"
        }
      ]
    },
    {
      "accountId": "000024eae1a211ef95a30242ac180002",
      "type": "saving-account",
      "currency": "THB",
      "accountNumber": "568-2-94760",
      "issuer": "TestLab",
      "amount": 93311.34,
      "color": "#f186c4",
      "isMainAccount": false,
      "progress": 82,
      "flags": [
        {
          "flagId": 3810099,
          "flagType": "system",
          "flagValue": "Flag3"
        },
        {
          "flagId": 3810100,
          "flagType": "system",
          "flagValue": "Disbursement"
        }
      ]
    },
    {
      "accountId": "000028aae1a211ef95a30242ac180002",
      "type": "saving-account",
      "currency": "THB",
      "accountNumber": "568-2-45295",
      "issuer": "TestLab",
      "amount": 68884.58,
      "color": "#5523bb",
      "isMainAccount": false,
      "progress": 8,
      "flags": [
        {
          "flagId": 3810101,
          "flagType": "system",
          "flagValue": "Flag5"
        },
        {
          "flagId": 3810102,
          "flagType": "system",
          "flagValue": "Overdue"
        }
      ]
    }
  ],
  "totalBalance": 248243.98
}
```

---

### PUT /update-account – Update Account Data

#### Update Main Account

**Request Header:**
```
Authorization: Bearer {authToken}
```

**Request Body:**
```json
{
  "accountId": "000020ece1a211ef95a30242ac180002",
  "isMainAccount": true,
  "color": ""
}
```

**Response:**
```json
{}
```

#### Update Color

**Request Header:**
```
Authorization: Bearer {authToken}
```

**Request Body:**
```json
{
  "accountId": "000020ece1a211ef95a30242ac180002",
  "isMainAccount": false,
  "color": "#24c875"
}
```

**Response:**
```json
{}
```

---

### GET /transactions – Get Transaction Data

**Request Header:**
```
Authorization: Bearer {authToken}
```

**Request Body:**
```json
{}
```

**Response:**
```json
{
  "transactions": [
    {
      "transactionId": "000018c1e1a211ef95a30242ac180002",
      "name": "Transaction_135017",
      "image": "https://dummyimage.com/54x54/999/fff",
      "isBank": false
    },
    {
      "transactionId": "000018c1e1a211ef95a30242ac180003",
      "name": "Transaction_1350170",
      "image": "https://dummyimage.com/54x54/999/fff",
      "isBank": false
    }
  ]
}
```

---

### GET /health – Health Check

**Response:**
```json
{
  "status": "healthy"
}
```

---

## Unit Testing

### Run Unit Tests
```bash
make test-service
```

### Generate Mocks with Mockery

1. Install mockery:
```bash
make mock-install
```

2. Generate mock files:
```bash
mockery --all --dir=internal/core/port --output=internal/core/port/mocks --outpkg=mocks
```

---

See `Makefile` for other useful commands.
