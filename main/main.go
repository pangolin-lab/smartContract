package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/proton-lab/smartContract"
)

var RopstenKey = `{"address":"12ab78538d47121f4c0af38c602e35f703e2363a","crypto":{"cipher":"aes-128-ctr","ciphertext":"b96c5da525d4cd044b60bb6ee0ea7b7896da260f0833833fa5093c4a551dd916","cipherparams":{"iv":"a879d558254511a063293ed867449356"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"cecc744962c2c6af044ba412b4f14de18ae92ef1561dead2c8a68f70d9970796"},"mac":"8a345ac4b9b11c3e6fc6f47ab2da509f9a74c28d0df4d2c77704cc1a4f79f541"},"id":"e22d11c2-96df-45e2-b132-905f9d8e821f","version":3}`

func main() {
	conn, err := ethclient.Dial("https://ropsten.infura.io/v3/8b8db3cca50a4fcf97173b7619b1c4c3")
	if err != nil {
		panic(err)
	}
	//auth, err := bind.NewTransactor(strings.NewReader(RopstenKey), os.Args[1])
	//if err != nil {
	//	panic(err)
	//}
	manager, err := service.NewSimpleProtonManager(common.HexToAddress("0x7fBA95e059D574A6B7001B1CA9bF03FD08E47F37"), conn)

	//bub := [32]byte{0x01, 0x02, 0x03}
	//bub[0] = 0x0a
	//bub[1] = 0x0b
	//
	addr, err := manager.Owner(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(addr.Hex())
}
