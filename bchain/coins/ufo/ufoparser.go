package ufo

import (
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

// magic numbers
const (
	MainnetMagic wire.BitcoinNet = 0xddb7d9fc
	TestnetMagic wire.BitcoinNet = 0xdbb8c0fb
)

// chain parameters
var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{27}
	MainNetParams.ScriptHashAddrID = []byte{68}
	MainNetParams.Bech32HRPSegwit = "uf"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{111}
	TestNetParams.ScriptHashAddrID = []byte{130}
	TestNetParams.Bech32HRPSegwit = "ut"
}

// UfoParser handle
type UfoParser struct {
	*btc.BitcoinLikeParser
}

// NewUfoParser returns new UfoParser instance
func NewUfoParser(params *chaincfg.Params, c *btc.Configuration) *UfoParser {
	return &UfoParser{BitcoinLikeParser: btc.NewBitcoinLikeParser(params, c)}
}

// GetChainParams contains network parameters for the main UFO network,
// and the test UFO network
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}
