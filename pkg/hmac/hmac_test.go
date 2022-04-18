// Copyright 2022 Wang Lei
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hmac

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func Test_GenerateInvalidGivesError(t *testing.T) {
	input := []byte("test")
	signature := "ab"
	secretKey := "key"

	err := Validate(input, signature, secretKey)
	if err == nil {
		t.Errorf("expected error when signature didn't have at least 5 characters in length")
		t.Fail()
		return
	}

	wantErr := "valid hash prefixes: [sha1=, sha256=], got: ab"
	if err.Error() != wantErr {
		t.Errorf("want: %s, got: %s", wantErr, err.Error())
		t.Fail()
	}
}

func Test_ValidateWithoutSha1PrefixFails(t *testing.T) {
	digest := "sign this message"
	key := "my key"

	encodeHash := "6791a762f7568f945c2e1e396cea243e944100a6"

	err := Validate([]byte(digest), encodeHash, key)

	if err == nil {
		t.Errorf("Expected error due to missing prefix")
		t.Fail()
	}
}

func Test_ValidateWithSha1Prefix(t *testing.T) {
	digest := "sign this message"
	key := "my key"

	encodeHash := "sha1=" + "6791a762f7568f945c2e1e396cea243e944100a6"

	err := Validate([]byte(digest), encodeHash, key)

	if err != nil {
		t.Errorf("Expected no error, but got: %s", err.Error())
		t.Fail()
	}
}

func Test_SignWithKey(t *testing.T) {
	digest := "sign this message"
	key := []byte("my key")

	wantHash := "6791a762f7568f945c2e1e396cea243e944100a6"

	hash := Sign([]byte(digest), key, sha1.New)
	encodeHash := hex.EncodeToString(hash)

	if encodeHash != wantHash {
		t.Errorf("Sign want hash: %s, got: %s", wantHash, encodeHash)
		t.Fail()
	}
}

func Test_SignWithKey_SHA256(t *testing.T) {
	digest := "sign this message"
	key := []byte("my key")

	wantHash := "41f8b7712c58dc25be8d30cf25e57739a65f5f2f449b59a42e04da1f191512e7"

	hash := Sign([]byte(digest), key, sha256.New)
	encodeHash := hex.EncodeToString(hash)

	if encodeHash != wantHash {
		t.Errorf("Sign want hash: %s, got: %s", wantHash, encodeHash)
		t.Fail()
	}
}

func Test_ValidateWithSha256Prefix(t *testing.T) {
	digest := "sign this message"
	key := "my key"

	encodeHash := "sha256=" + "41f8b7712c58dc25be8d30cf25e57739a65f5f2f449b59a42e04da1f191512e7"

	err := Validate([]byte(digest), encodeHash, key)

	if err != nil {
		t.Errorf("Expectecd no error, but got: %s", err.Error())
		t.Fail()
	}
}
