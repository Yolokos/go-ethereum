package score

import (
	"strings"

	_ "embed"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

//go:embed ScoreContract.abi
var scoreABIJson string

var parsedABI abi.ABI

func init() {
	var err error
	parsedABI, err = abi.JSON(strings.NewReader(scoreABIJson))
	if err != nil {
		panic(err)
	}
}

func GetABI() abi.ABI {
	return parsedABI
}
