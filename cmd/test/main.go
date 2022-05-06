package test

import (
	"fmt"

	"github.com/xpaymentsorg/go-xpayments/crypto"
)

var (
	key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr   = crypto.PubkeyToAddress(key.PublicKey)
)

func main() {
	fmt.Print("KEY: ", key)
	fmt.Print("ADDR: ", addr)
}
