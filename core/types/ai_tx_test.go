package types

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestAITxHashing(t *testing.T) {
	key, _ := crypto.GenerateKey()

	tx := createAITx(key)
	hash := tx.Hash()

	t.Log("AI tx hash:", hash)

	enc, err := tx.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	var decodedTx Transaction
	if err := decodedTx.UnmarshalBinary(enc); err != nil {
		t.Fatal(err)
	}

	hash2 := decodedTx.Hash()

	if hash != hash2 {
		t.Fatal("hash mismatch after encode/decode", hash, hash2)
	}
}

func TestAITxSize(t *testing.T) {
	key, _ := crypto.GenerateKey()

	tx := createAITx(key)

	size := tx.Size()
	t.Log("AI tx size:", size)

	enc, err := tx.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}

	if size != uint64(len(enc)) {
		t.Error("wrong size:", size, "encoded:", len(enc))
	}
}

func TestAITxSigning(t *testing.T) {
	key, _ := crypto.GenerateKey()

	tx := createAITx(key)

	signer := NewCancunSigner(big.NewInt(1))

	signed := MustSignNewTx(key, signer, tx.inner)

	sender, err := Sender(signer, signed)
	if err != nil {
		t.Fatal(err)
	}

	expected := crypto.PubkeyToAddress(key.PublicKey)

	t.Log("sender:", sender)

	if sender != expected {
		t.Fatal("wrong sender recovered", sender, expected)
	}
}

func createAITx(key *ecdsa.PrivateKey) *Transaction {
	share := ShareObject{
		MinerPubKey: common.HexToAddress("0x1234"),
		TargetID:    12345,
		DataHash:    common.HexToHash("0x01"),
		DataLink:    "ipfs://QmHash",
		LossClaimed: 5,
		Fingerprint: []byte{1, 2, 3},
		EpochID:     10,
		Nonce:       999,
		Signature:   []byte{0xaa, 0xbb},
	}

	tx := NewAITx()

	tx.ChainID.SetFromBig(big.NewInt(1))
	tx.Nonce = 1
	tx.GasTipCap.SetFromBig(big.NewInt(2))
	tx.GasFeeCap.SetFromBig(big.NewInt(3))
	tx.Gas = 21000
	tx.To = &common.Address{0x01, 0x02}
	tx.Value.SetFromBig(big.NewInt(100))
	tx.Data = []byte("ai transaction test")
	tx.Share = share

	signer := NewCancunSigner(tx.ChainID.ToBig())

	return MustSignNewTx(key, signer, tx)
}