package util

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/cooperstrahan/distributed_voting/go-vote-api/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Address string

//stringToKeccak256 converts a string to a keccak256 hash of type [32]byte
func StringToKeccak256(s string) [32]byte {
	var toReturn [32]byte
	copy(toReturn[:], crypto.Keccak256([]byte(s))[:])
	return toReturn
}

func KeccakToString(b []byte) string {
	h := crypto.Keccak256Hash(b)
	return h.String()
}

func DeployContract() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA("4df038bcc399d525b1dd61ae39cb29909b7c45fbeb9a38744fd3f2894a7f0c90")
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("invalid key")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = big.NewInt(1000000)

	address, tx, instance, err := contracts.DeployContracts(auth, client)
	if err != nil {
		panic(err)
	}

	fmt.Println(address.Hex())

	_, _ = instance, tx

	Address = address.Hex()
	// return address.Hex()

}

func GetAddress() string {
	return Address
}
