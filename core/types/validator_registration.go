package types

type ValidatorRegistration struct {
	Pubkey []byte `json:"pubkey"` // BLS public key of the validator
}

type ValidatorRegistrations []*ValidatorRegistration
