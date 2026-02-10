package pouw

import (
	"errors"
	"math/big"
    "github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
    "github.com/ethereum/go-ethereum/rpc"
    "github.com/ethereum/go-ethereum/trie"
)

type Engine struct {
	config *params.ChainConfig
	db     ethdb.Database
}

func New(config *params.ChainConfig, db ethdb.Database) *Engine {
	return &Engine{
		config: config,
		db:     db,
	}
}

func (e *Engine) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

func (e *Engine) VerifyHeader(chain consensus.ChainHeaderReader, header *types.Header) error {
	// Пропускаем genesis
	if header.Number.Uint64() == 0 {
		return nil
	}

	// TODO: включить позже
    if header.PoUWLoss == 0 {
        return errors.New("missing PoUWLoss")
    }
	if header.PoUWCommitment == (common.Hash{}) {
		return errors.New("missing PoUWCommitment")
	}
	return nil
}

func (e *Engine) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
    return nil // пока uncles не используем
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
	// PoUW проверяется в VerifyHeader
	return nil
}

func (e *Engine) SealHash(header *types.Header) common.Hash {
    return header.Hash() // минимально
}

func (e *Engine) Prepare(chain consensus.ChainHeaderReader, header *types.Header) error {
	header.Difficulty = big.NewInt(1)
	return nil
}

func (e *Engine) Finalize(
    chain consensus.ChainHeaderReader,
    header *types.Header,
    state *state.StateDB,
    txs []*types.Transaction,
    uncles []*types.Header,
    receipts []*types.Receipt,
) (*types.Block, error) {

    if e.config != nil && header.Number != nil {
        header.Root = state.IntermediateRoot(e.config.IsEIP158(header.Number))
    } else {
        header.Root = state.IntermediateRoot(false)
    }

    body := &types.Body{
        Transactions: txs,
        Uncles:       uncles,
    }

    return types.NewBlock(header, body, receipts, trie.NewStackTrie(nil)), nil
}

func (e *Engine) Seal(chain consensus.ChainHeaderReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
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