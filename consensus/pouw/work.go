package pouw

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/common"
)

func ComputePoUW(data []byte) (common.Hash, uint64) {
	hash := sha256.Sum256(data)
	
	var loss uint64
	for _, b := range hash[:] {
		loss += uint64(b)
	}
	return common.BytesToHash(hash[:]), loss
}