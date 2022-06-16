package hmac

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"
)

func CheckMAC(message, messageMAC, key []byte, sha func() hash.Hash) bool {
	mac := hmac.New(sha, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)

	return hmac.Equal(messageMAC, expectedMAC)
}

func Sign(message, key []byte, sha func() hash.Hash) []byte {
	mac := hmac.New(sha, key)
	mac.Write(message)
	signed := mac.Sum(nil)

	return signed
}

func Validate(bytesIn []byte, encodeHash string, secretKey string) error {
	var validated error

	var hashFn func() hash.Hash
	var payload string

	if strings.HasPrefix(encodeHash, "sha1=") {

		payload = strings.TrimPrefix(encodeHash, "sha1=")

		hashFn = sha1.New

	} else if strings.HasPrefix(encodeHash, "sha256=") {

		payload = strings.TrimPrefix(encodeHash, "sha256=")

		hashFn = sha256.New

	} else {
		return fmt.Errorf("valid hash prefixes: [sha1=, sha256=], got: %s", encodeHash)
	}

	messageMAC := payload
	messageMACBuf, _ := hex.DecodeString(messageMAC)

	res := CheckMAC(bytesIn, []byte(messageMACBuf), []byte(secretKey), hashFn)

	if !res {
		validated = fmt.Errorf("invalid message digest or secret")
	}

	return validated
}
