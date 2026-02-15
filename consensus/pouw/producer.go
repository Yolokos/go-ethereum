package pouw

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

func (e *Engine) StartBlockProducer(chain *core.BlockChain, interval time.Duration, stop <-chan struct{}) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Info("PoUW block producer started new")

	for {
		select {
		case <-ticker.C:
			parent := chain.CurrentBlock()
			if parent == nil {
				continue
			}

			header := &types.Header{
				ParentHash: parent.Hash(),
				Number:     new(big.Int).Add(parent.Number, big.NewInt(1)),
				Time:       uint64(time.Now().Unix()),
				Difficulty: big.NewInt(1),
				Coinbase:   common.Address{},
			}

			if err := e.Prepare(chain, header); err != nil {
				log.Error("Prepare failed", "err", err)
				continue
			}

			statedb, err := chain.StateAt(parent.Root)
			if err != nil {
				log.Error("Failed to create state", "err", err)
				continue
			}

			// 3️⃣ Финализация
			block, err := e.FinalizeAndAssemble(
				chain,
				header,
				statedb,
				&types.Body{},   // пустое тело
				nil,             // нет receipts
			)
			if err != nil {
				log.Error("Finalize failed", "err", err)
				continue
			}

			// 4️⃣ Seal
			results := make(chan *types.Block, 1)
			if err := e.Seal(chain, block, results, stop); err != nil {
				log.Error("Seal failed", "err", err)
				continue
			}
			sealed := <-results

			// 5️⃣ Вставка
			if _, err := chain.InsertChain([]*types.Block{sealed}); err != nil {
				log.Error("InsertChain failed", "err", err)
			} else {
				log.Info("New block inserted", "number", sealed.Number())
			}

		case <-stop:
			return
		}
	}
}