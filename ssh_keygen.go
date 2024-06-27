package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func GetPubKeyFingerprintLegacyMD5(key string) (string, error) {
	pk, _, _, _, err := ssh.ParseAuthorizedKey([]byte(key)) //nolint:dogsled
	if err != nil {
		return "", fmt.Errorf("parse authorized key: %w", err)
	}

	fingerprint := ssh.FingerprintLegacyMD5(pk)

	return fingerprint, nil
}
