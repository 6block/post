package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"log"
	"os"
	"path/filepath"
)

func load_pubkey() []byte {
	keyPath := filepath.Join(opts.DataDir, edKeyFileName)
	data, err := os.ReadFile(keyPath)
	if err != nil {
		return nil
	}
	dst := make([]byte, ed25519.PrivateKeySize)
	n, err := hex.Decode(dst, data)
	if err != nil {
		log.Fatalf("failed to decode private key from %s: %s\n", keyPath, err)
		return nil
	}
	if n != ed25519.PrivateKeySize {
		log.Fatalf("size of key (%d) not expected size %d\n", n, ed25519.PrivateKeySize)
		return nil
	}
	pub := ed25519.NewKeyFromSeed(dst[:ed25519.SeedSize]).Public().(ed25519.PublicKey)
	return pub
}
