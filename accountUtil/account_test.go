package accountUtil

import (
	"regexp"
	"testing"

	cryptoutil "github.com/lily0749labs/goutils/crypto"
)

func TestGenerateNickname(t *testing.T) {
	pattern := regexp.MustCompile(`^PLAY__[A-Za-z]{9}$`)
	for range 100 {
		nickname := GenerateNickname()
		if !pattern.MatchString(nickname) {
			t.Fatalf("GenerateNickname() = %q, want PLAY__ followed by 9 ASCII letters", nickname)
		}
	}
}

func TestGenerateSMSCode(t *testing.T) {
	for range 1000 {
		code := GenerateSMSCode()
		if code < 100000 || code > 999998 {
			t.Fatalf("GenerateSMSCode() = %d, want value in [100000, 999998]", code)
		}
	}
}

func TestGeneratePasswordAt(t *testing.T) {
	const (
		userID      = uint64(42)
		unixSeconds = int64(1_700_000_000)
	)

	hash := GeneratePasswordAt(userID, unixSeconds)
	if hash == "" {
		t.Fatal("GeneratePasswordAt() returned an empty hash")
	}
	if !cryptoutil.Crypto.BcryptCheck("42_generate_1700000000", hash) {
		t.Fatal("GeneratePasswordAt() hash does not match the legacy plaintext format")
	}
}
