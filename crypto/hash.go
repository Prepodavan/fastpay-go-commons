package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(payload string) string {
	hashAsBytes := sha256.Sum256([]byte(payload))
	return hex.EncodeToString(hashAsBytes[:])
}
