package main

import (
	"encoding/hex"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"
)

type EthCmd struct {
	Pattern string `long:"pattern" default:"42" description:"match pattern"`
}

func (ec *EthCmd) Execute(args []string) error {
	beginTime := time.Now()
	prefix := strings.ToLower(ec.Pattern)

	var numAttempts int64 = 0
	addrStr := ""
	keyStr := ""

	for {
		numAttempts++

		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		addrStr = hex.EncodeToString(addr[:])
		if matchPrefix(addrStr, prefix) {
			keyStr = hex.EncodeToString(crypto.FromECDSA(key))
			// fmt.Println("pub:", hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey)))
			break
		}
	}

	log.Infof("\nElapsed: %s\naddr: 0x%s\npvt: 0x%s\nattempts: %d.\n",
		time.Since(beginTime), addrStr, keyStr, numAttempts)

	return nil
}

var ethCmd EthCmd

func init() {
	parser.AddCommand("eth", "get a ETH vanity address", "The ticker command get a ETH vanity address", &ethCmd)
}
