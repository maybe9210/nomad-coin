package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/maybe9210/nomad-coin/utils"
)

const (
	privateKey    string = "3077020101042049319f52e92f6ba86a384a882f1e9cb3837518cab7441ac5542efcd6b0962bcba00a06082a8648ce3d030107a144034200045522f84857f582bd3ac4ef609487ac0a692baf56de5d70ab7b028d1a709b2dabb3079c95ca23d675a682bfac76dc04ed7ce792dc9950d2f2d243638a24d404c6"
	signature     string = "7e193ff6bbe34ba6249dc2fe2c320fcd8ea1e1522b08c9dec43198a49547bb0728d903021bb70b83462612328de63ae135f8f2c9419868a5353ce9d5f0cbe532"
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

func Start() {
	privBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	private, err := x509.ParseECPrivateKey(privBytes)
	fmt.Printf("%d\n\n", private)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)
	utils.HandleErr(err)
	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)
}
