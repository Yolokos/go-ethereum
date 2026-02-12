package pouw

import (
	"errors"
	"math/big"
	"github.com/ethereum/go-ethereum/log"
    "github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
    "github.com/ethereum/go-ethereum/rpc"
    "github.com/ethereum/go-ethereum/trie"
    "github.com/ethereum/go-ethereum/core/vm"
)


type Engine struct {
	config *params.ChainConfig
	db     ethdb.Database
	// Validators []common.Address
	// EpochCache map[uint64][]common.Address
	// EpochLength uint64
	// Admin       common.Address
}

func New(config *params.ChainConfig, db ethdb.Database/*, epochLength uint64 */) *Engine {
	return &Engine{
		config: config,
		db:     db,
		// Validators: []common.Address{},
		// EpochCache: make(map[uint64][]common.Address),
		// EpochLength: epochLength,
	}
}

// func (e *Engine) AddValidator(caller, addr common.Address) error {
// 	if caller != e.Admin {
// 		return errors.New("unauthorized: only admin can add validators")
// 	}

// 	for _, v := range e.Validators {
// 		if v == addr {
// 			return nil
// 		}
// 	}
// 	e.Validators = append(e.Validators, addr)

// 	// очищаем кэш будущих эпох
// 	for epoch := range e.EpochCache {
// 		if epoch >= e.CurrentEpoch() {
// 			delete(e.EpochCache, epoch)
// 		}
// 	}

// 	log.Info("Validator added", "address", addr.Hex())
// 	return nil
// }

// func (e *Engine) RandomValidators(epoch uint64, count int) []common.Address {
// 	if e.EpochCache == nil {
// 		e.EpochCache = make(map[uint64][]common.Address)
// 	}
// 	if val, ok := e.EpochCache[epoch]; ok {
// 		return val
// 	}

// 	selected := make([]common.Address, 0, count)
// 	total := len(e.Validators)
// 	if total == 0 {
// 		return selected
// 	}

// 	seed := big.NewInt(0).SetBytes(sha256.New().Sum([]byte(string(epoch))))
// 	for len(selected) < count {
// 		idx := int(seed.Int64() % int64(total))
// 		selected = append(selected, e.Validators[idx])
// 		seed.Add(seed, big.NewInt(1))
// 	}

// 	e.EpochCache[epoch] = selected
// 	return selected
// }

// func (e *Engine) EpochNumber(blockNumber uint64) uint64 {
// 	return blockNumber / EpochLength
// }

func (e *Engine) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func (e *Engine) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header) error {
	log.Info("PoUW VerifyHeader called", "number", header.Number)

	if header.Number.Uint64() == 0 {
		return nil
	}

	if header.PoUWLoss == nil {
		return errors.New("missing PoUWLoss")
	}
	if header.PoUWCommitment == nil {
		return errors.New("missing PoUWCommitment")
	}

	// epoch := e.EpochNumber(header.Number.Uint64())
	// validators := e.RandomValidators(epoch, 3)

	// isValidator := false
	// for _, v := range validators {
	// 	if v == header.Coinbase {
	// 		isValidator = true
	// 		break
	// 	}
	// }

	// if !isValidator {
	// 	return errors.New("author is not selected validator for this epoch")
	// }

	return nil
}

func (e *Engine) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
    return nil
}

func (e *Engine) VerifyHeaders(chain consensus.ChainHeaderReader, headers []*types.Header) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for _, h := range headers {
			select {
			case <-abort:
				return
			default:
				results <- e.VerifyHeader(chain, h)
			}
		}
	}()
	return abort, results
}

func (e *Engine) VerifySeal(chain consensus.ChainHeaderReader, header *types.Header) error {

	return nil
}

func (e *Engine) SealHash(header *types.Header) common.Hash {
    return header.Hash()
}

func (e *Engine) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	header.Difficulty = big.NewInt(1)
	return nil
}

func (e *Engine) Finalize(
    chain consensus.ChainHeaderReader,
    header *types.Header,
    state vm.StateDB,
    body *types.Body,
) {
	log.Info("PoUW Finalize called", "number", header.Number)
}


func (e *Engine) FinalizeAndAssemble(
    chain consensus.ChainHeaderReader,
    header *types.Header,
    state *state.StateDB,
    body *types.Body,
    receipts []*types.Receipt,
) (*types.Block, error) {
    header.Root = state.IntermediateRoot(false)

    return types.NewBlock(
        header,
        body,
        receipts,
        trie.NewStackTrie(nil),
    ), nil
}

func (e *Engine) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	log.Info("PoUW Seal called", "number", block.Number())
	select {
	case results <- block:
		return nil
	case <-stop:
		return errors.New("sealing stopped")
	}
}

func (e *Engine) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent *types.Header) *big.Int {
	return big.NewInt(0)
}

func (e *Engine) APIs(chain consensus.ChainHeaderReader) []rpc.API {
	return nil
}

func (e *Engine) Close() error {
	return nil
}