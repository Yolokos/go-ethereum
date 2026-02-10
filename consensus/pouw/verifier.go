package pouw


import (
	"github.com/ethereum/go-ethereum/common"
)

type PoUWPayload struct {
	WorkCommitment common.Hash
	Loss           uint64
}

func Verify(payload *PoUWPayload, referenceHash common.Hash, referenceLoss uint64) bool {
	if payload.WorkCommitment != referenceHash {
		return false
	}
	if payload.Loss != referenceLoss {
		return false
	}
	return true
}