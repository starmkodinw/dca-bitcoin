package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func mustMarshal(v interface{}) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}

func genSign(apiSecret, payloadString string) string {
	mac := hmac.New(sha256.New, []byte(apiSecret))
	mac.Write([]byte(payloadString))
	return hex.EncodeToString(mac.Sum(nil))
}
