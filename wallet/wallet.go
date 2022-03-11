package wallet

import (
	"crypto/ecdsa"
	"os"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat("nomadcoin.wallet")
	return !os.IsNotExist(err)
}

func Wallet() *wallet {
	if w == nil {
		if hasWalletFile() {

		}
	}
	return w
}

func Start() {

}
