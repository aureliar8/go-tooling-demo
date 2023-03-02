package ex3

import (
	"encoding/hex"
	"fmt"

	"github.com/cespare/xxhash/v2"
)

func nameHash(name string) string {
	hasher := xxhash.New()
	hash := hasher.Sum([]byte(name))
	return hex.EncodeToString(hash)
}

func nameScore(name string) int {
	score := 0
	for _, c := range name {
		score += int(c - 'a')
	}
	score /= len(name)
	return score
}

func formatReply(name string, hash string, score int) string {
	return fmt.Sprintf("Hello %s. Your hash is %s and your score is %d", name, hash, score)
}

func foo()     {}
func setup()   {}
func cleanup() {}
