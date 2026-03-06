package types

import (
	"math/big"
	"bytes"

	"github.com/holiman/uint256"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type AITx struct {
	ChainID   *uint256.Int
	Nonce     uint64
	GasTipCap *uint256.Int
	GasFeeCap *uint256.Int
	Gas       uint64
	To        *common.Address `rlp:"nil"`
	Value     *uint256.Int
	Data      []byte

	Share ShareObject

	V         *uint256.Int
	R         *uint256.Int
	S         *uint256.Int
}

type ShareObject struct {
	MinerPubKey common.Address `json:"miner_pubkey"`

	TargetID uint64 `json:"target_id"`

	DataHash common.Hash `json:"data_hash"`
	DataLink string `json:"data_link"`

	LossClaimed uint64 `json:"loss_claimed"`

	Fingerprint []byte `json:"fingerprint"`

	EpochID uint64 `json:"epoch_id"`
	Nonce   uint64 `json:"nonce"`

	Signature []byte `json:"signature"`
}

func NewAITx() *AITx {
    return &AITx{
        ChainID:   uint256.NewInt(0),
        GasTipCap: uint256.NewInt(0),
        GasFeeCap: uint256.NewInt(0),
        Value:     uint256.NewInt(0),
        V:         uint256.NewInt(0),
        R:         uint256.NewInt(0),
        S:         uint256.NewInt(0),
    }
}

func (tx *AITx) txType() byte { 
	return AITxType 
}

func (tx *AITx) copy() TxData {
	cpy := *tx
	return &cpy
}

func (tx *AITx) chainID() *big.Int { return tx.ChainID.ToBig() }

func (tx *AITx) accessList() AccessList { return nil }

func (tx *AITx) data() []byte { return tx.Data }

func (tx *AITx) gas() uint64 { return tx.Gas }

func (tx *AITx) gasFeeCap() *big.Int { return tx.GasFeeCap.ToBig() }

func (tx *AITx) gasTipCap() *big.Int { return tx.GasTipCap.ToBig() }

func (tx *AITx) gasPrice() *big.Int { return tx.GasFeeCap.ToBig() }

func (tx *AITx) value() *big.Int { return tx.Value.ToBig() }

func (tx *AITx) nonce() uint64 { return tx.Nonce }

func (tx *AITx) to() *common.Address { return tx.To }

func (tx *AITx) rawSignatureValues() (v, r, s *big.Int) {
    return tx.V.ToBig(), tx.R.ToBig(), tx.S.ToBig()
}

func (tx *AITx) setSignatureValues(chainID, v, r, s *big.Int) {
    if tx.ChainID == nil {
        tx.ChainID = uint256.NewInt(0)
    }
    if tx.V == nil {
        tx.V = uint256.NewInt(0)
    }
    if tx.R == nil {
        tx.R = uint256.NewInt(0)
    }
    if tx.S == nil {
        tx.S = uint256.NewInt(0)
    }

    tx.ChainID.SetFromBig(chainID)
    tx.V.SetFromBig(v)
    tx.R.SetFromBig(r)
    tx.S.SetFromBig(s)
}

func (tx *AITx) encode(b *bytes.Buffer) error {
	return rlp.Encode(b, []interface{}{
		tx.ChainID,
		tx.Nonce,
		tx.GasTipCap,
		tx.GasFeeCap,
		tx.Gas,
		tx.To,
		tx.Value,
		tx.Data,
		tx.Share,
		tx.V,
		tx.R,
		tx.S,
	})
}

func (tx *AITx) decode(input []byte) error {
	var dec struct {
		ChainID   *uint256.Int
		Nonce     uint64
		GasTipCap *uint256.Int
		GasFeeCap *uint256.Int
		Gas       uint64
		To        *common.Address
		Value     *uint256.Int
		Data      []byte
		Share     ShareObject
		V         *uint256.Int
		R         *uint256.Int
		S         *uint256.Int
	}

	if err := rlp.DecodeBytes(input, &dec); err != nil {
		return err
	}

	tx.ChainID = dec.ChainID
	tx.Nonce = dec.Nonce
	tx.GasTipCap = dec.GasTipCap
	tx.GasFeeCap = dec.GasFeeCap
	tx.Gas = dec.Gas
	tx.To = dec.To
	tx.Value = dec.Value
	tx.Data = dec.Data
	tx.Share = dec.Share
	tx.V = dec.V
	tx.R = dec.R
	tx.S = dec.S

	return nil
}

func (tx *AITx) sigHash(chainID *big.Int) common.Hash {
	return prefixedRlpHash(
		AITxType,
		[]interface{}{
			chainID,
			tx.Nonce,
			tx.GasTipCap,
			tx.GasFeeCap,
			tx.Gas,
			tx.To,
			tx.Value,
			tx.Data,
			tx.Share,
		},
	)
}

func (tx *AITx) effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int {
	if baseFee == nil {
		return dst.Set(tx.GasFeeCap.ToBig())
	}

	tip := new(big.Int).Sub(tx.GasFeeCap.ToBig(), baseFee)
	if tip.Cmp(tx.GasTipCap.ToBig()) > 0 {
		tip = tx.GasTipCap.ToBig()
	}

	return dst.Add(tip, baseFee)
}