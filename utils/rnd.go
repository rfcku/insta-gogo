package utils

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// string from timestamp sha256
func NowFileName() string {
	now := time.Now().Format("20060102150405")

	hash := sha256.Sum256([]byte(now))
	return fmt.Sprintf("%x", hash)

}
