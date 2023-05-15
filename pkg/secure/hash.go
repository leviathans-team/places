package secure

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
)

func CalcSignature(secret string, message string) string {
	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
