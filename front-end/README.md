# Deploy front-end
```
cd front-end
```

## Project setup
```
yarn install
```

### Compiles and hot-reloads for development
```
yarn serve
```

### Compiles and minifies for production
```
yarn build
```

### Lints and fixes files
```
yarn lint
```

# Run Ganache
https://www.trufflesuite.com/docs/ganache/quickstart

# Update files with correct private key for your Ganache instance
server.go lines 52, 162
/util/util.go line 35

# Install go
https://golang.org/doc/install

# Install go ethereum
https://geth.ethereum.org/docs/install-and-build/installing-geth

# Install the Solidity Compiler
https://docs.soliditylang.org/en/v0.5.3/installing-solidity.html

# Deploy go-vote-api
cd go-vote-api
go run server.go


The app will then be available at
http://localhost:8080
