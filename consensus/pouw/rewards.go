package pouw

import "fmt"

func RewardMiner(minerAddress string, amount uint64) {
	fmt.Printf("Rewarding miner %s with %d tokens\n", minerAddress, amount)
}