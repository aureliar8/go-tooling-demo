package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func nameHash(name string) string {
	hasher := md5.New()
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
