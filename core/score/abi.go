package score

import (
	"strings"

	_ "embed"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var parsedABI abi.ABI

func init() {
	var err error
	parsedABI, err = abi.JSON(strings.NewReader(ScoreMetaData.ABI))
	if err != nil {
		panic(err)
	}
}

func GetABI() abi.ABI {
	return parsedABI
}
