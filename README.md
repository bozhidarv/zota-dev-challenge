# Zota dev challenge

# Run the project

after you clone the project tidy the go modules

```bash
go mod tidy
```

then you can run the project with the following command

```bash
ZOTAPAY_ENDPOINT_ID="<YOUR_ENDPOINT_ID>" ZOTAPAY_API_KEY="<YOUR_API_KEY>" ZOTAPAY_CURR="<YOUR_CURRENCY>" ZOTAPAY_MERCHANT_ID="<YOUR_MERCHANT_ID>" ZOTAPAY_BASE_URL="<YOUR_BASE_URL>" go run ./cmd/zota-dev-challenge/main.go
```

### Author: Bozhidar Videnov
