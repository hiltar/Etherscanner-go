# Etherscanner-go
This CLI program will fetch etherscan V2 API every 120 seconds and returns balance of Ethereum in the address.

# Usage
`./etherscanner <address> <API-key> [chainid]`  
`chainid` argument is optional, default is `1`

## Output
```
Balance: 0.001172726557641221 ETH
```

# Obtain API key from etherscan
Register into etherscan.io and create an API-key from ![API Dashboard](https://etherscan.io/apidashboard).

# Build
```
go mod init etherscanner
go mod tidy
CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -o etherscanner
upx --best etherscanner
```
