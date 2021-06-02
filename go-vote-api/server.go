package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/cooperstrahan/distributed_voting/go-vote-api/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/cooperstrahan/distributed_voting/go-vote-api/contracts"
)

type voter struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Ssn        string `json:"ssn"`
}

type vote struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Ssn        string `json:"ssn"`
	Candidate  string `json:"candidate"`
}

var Address string

func registerVoter(w http.ResponseWriter, r *http.Request) {

	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
	conn, err := contracts.NewContracts(common.HexToAddress(Address), client)
	if err != nil {
		panic(err)
	}

	// connector, err := contracts.NewContractsCaller(common.HexToAddress(Address), client)
	// if err != nil {
	// 	panic(err)
	// }

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
	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	panic(err)
	// }

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	if r.Method == "OPTIONS" {
		return
	}

	voter := &voter{}

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&voter)
	if err != nil {
		log.Println("err decoding")
		w.WriteHeader(http.StatusInternalServerError)
	}

	v, err := json.MarshalIndent(voter, "", "  ")
	if err != nil {
		log.Println("err marshalling")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("voter: ", string(v))

	str := voter.First_name + voter.Last_name + voter.Ssn

	hash := util.StringToKeccak256(str)

	str = util.KeccakToString(hash[:])

	fmt.Println("voter id string: ", str)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, `{"voter_id":%q`, str)

	var user_id [32]byte
	copy(user_id[:], str)

	output, err := conn.RegisterVoter(&bind.TransactOpts{
		From:   fromAddress, // Ethereum account to send the transaction from
		Nonce:  nil,         // Nonce to use for the transaction execution (nil = use pending state)
		Signer: auth.Signer, // Method to use for signing the transaction (mandatory)

		Value:    big.NewInt(0),       // *big.Int // Funds to transfer along the transaction (nil = 0 = no funds)
		GasPrice: big.NewInt(1000000), // *big.Int // Gas price to use for the transaction execution (nil = gas price oracle)
		GasLimit: uint64(3000000),     //uint64   // Gas limit to set for the transaction execution (0 = estimate)
	}, user_id)

	// o, err := connector.RegisterVoter(&bind.CallOpts{}, user_id)

	// log.Println(o)
	// if(output){
	// 	log.Println("voter is registered")
	// } else {
	// 	log.Println("voter is NOT registered")
	// }
	log.Println(output)
}

func recordVote(w http.ResponseWriter, request *http.Request) {
	if origin := request.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
	conn, err := contracts.NewContracts(common.HexToAddress(Address), client)
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
	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	// if err != nil {
	// 	panic(err)
	// }

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}

	if request.Method == "OPTIONS" {
		return
	}

	vote := &vote{}

	decoder := json.NewDecoder(request.Body)

	err = decoder.Decode(&vote)
	if err != nil {
		log.Println("err decoding")
		w.WriteHeader(http.StatusInternalServerError)
	}

	v, err := json.MarshalIndent(vote, "", "  ")
	if err != nil {
		log.Println("err marshalling")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("voter: ", string(v))

	str := vote.First_name + vote.Last_name + vote.Ssn

	hash := util.StringToKeccak256(str)

	str = util.KeccakToString(hash[:])

	fmt.Println("voter id string: ", str)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, `{"voter_id":%q`, str)

	var user_id [32]byte
	copy(user_id[:], str)

	output, err := conn.RecordVote(&bind.TransactOpts{
		From:   fromAddress, // Ethereum account to send the transaction from
		Nonce:  nil,         // Nonce to use for the transaction execution (nil = use pending state)
		Signer: auth.Signer, // Method to use for signing the transaction (mandatory)

		Value:    big.NewInt(0),       // *big.Int // Funds to transfer along the transaction (nil = 0 = no funds)
		GasPrice: big.NewInt(1000000), // *big.Int // Gas price to use for the transaction execution (nil = gas price oracle)
		GasLimit: uint64(3000000),     //uint64   // Gas limit to set for the transaction execution (0 = estimate)
	}, user_id)

	log.Println(output)
}

func confirmRegister(w http.ResponseWriter, request *http.Request) {
	if origin := request.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
	conn, err := contracts.NewContracts(common.HexToAddress(Address), client)
	if err != nil {
		panic(err)
	}

	first_name := request.URL.Query().Get("first_name")
	last_name := request.URL.Query().Get("last_name")
	ssn := request.URL.Query().Get("ssn")

	str := first_name + last_name + ssn

	hash := util.StringToKeccak256(str)

	str = util.KeccakToString(hash[:])

	fmt.Println("voter id string: ", str)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, `{"voter_id":%q`, str)

	var user_id [32]byte
	copy(user_id[:], str)

	output, err := conn.CheckVoterRegistered(&bind.CallOpts{}, user_id)

	if output {
		log.Println("voter is registered")
	} else {
		log.Println("voter is NOT registered")
	}

}

func confirmVote(w http.ResponseWriter, request *http.Request) {
	if origin := request.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		panic(err)
	}
	conn, err := contracts.NewContracts(common.HexToAddress(Address), client)
	if err != nil {
		panic(err)
	}

	first_name := request.URL.Query().Get("first_name")
	last_name := request.URL.Query().Get("last_name")
	ssn := request.URL.Query().Get("ssn")

	str := first_name + last_name + ssn

	hash := util.StringToKeccak256(str)

	str = util.KeccakToString(hash[:])

	fmt.Println("voter id string: ", str)

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, `{"voter_id":%q`, str)

	var user_id [32]byte
	copy(user_id[:], str)

	output, err := conn.CheckCastVote(&bind.CallOpts{}, user_id)

	if output {
		log.Println("voter has voted")
	} else {
		log.Println("voter has NOT voted")
	}
}

func main() {

	util.DeployContract()

	Address = util.GetAddress()

	http.HandleFunc("/register", registerVoter)
	http.HandleFunc("/vote", recordVote)
	http.HandleFunc("/confirm-register", confirmRegister)
	http.HandleFunc("/confirm-vote", confirmVote)
	http.ListenAndServe(":8040", nil)

}
