// Copyright (c) 2015 The Decred developers
// Copyright (c) 2016-2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chainhash

import (
	"crypto"
	"crypto/sha256"

	// Blake2s is required due to Hash.New(blake2s)
	_ "golang.org/x/crypto/blake2s"

	"github.com/Groestlcoin/go-groestl-hash/groestl"
	"github.com/bitgoin/lyra2rev2"
	"github.com/marpme/go-x17"
	"golang.org/x/crypto/scrypt"
)

// HashB calculates hash(b) and returns the resulting bytes.
func HashB(b []byte) []byte {
	hash := sha256.Sum256(b)
	return hash[:]
}

// HashH calculates hash(b) and returns the resulting bytes as a Hash.
func HashH(b []byte) Hash {
	return Hash(sha256.Sum256(b))
}

// DoubleHashB calculates hash(hash(b)) and returns the resulting bytes.
func DoubleHashB(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}

// DoubleHashH calculates hash(hash(b)) and returns the resulting bytes as a
// Hash.
func DoubleHashH(b []byte) Hash {
	first := sha256.Sum256(b)
	return Hash(sha256.Sum256(first[:]))
}

// ScryptHash calculates scryptHash(b) and returns the resulting bytes as a Hash
func ScryptHash(b []byte) Hash {
	scryptHash, err := scrypt.Key(b, b, 1024, 1, 1, 32)
	if err != nil {
		panic(err)
	}

	var hash [32]byte
	for i := 0; i < len(hash); i++ {
		hash[i] = scryptHash[i]
	}
	return Hash(hash)
}

// GroestlHash calculates GroestlHash(b) and returns the resulting bytes as a Hash
func GroestlHash(b []byte) Hash {
	groestl, out := groestl.New(), [32]byte{}

	_, err := groestl.Write(b)
	if err != nil {
		panic(err)
	}

	groestl.Close(out[:], 0, 0)

	return Hash(out)
}

// BlakeHash calculates GroestlHash(b) and returns the resulting bytes as a Hash
func BlakeHash(b []byte) Hash {
	blake2s, out := crypto.BLAKE2s_256.New(), []byte{}

	_, err := blake2s.Write(b)
	if err != nil {
		panic(err)
	}

	out = blake2s.Sum(out)
	var hash [32]byte
	for i := 0; i < len(hash); i++ {
		hash[i] = out[i]
	}

	return Hash(hash)
}

// Lyra2Rev2Hash calculates Lyra2Rev2Hash(b) and returns the resulting bytes as a Hash
func Lyra2Rev2Hash(b []byte) Hash {
	lyraHash, err := lyra2rev2.Sum(b)

	if err != nil {
		panic(err)
	}

	var hash [32]byte
	for i := 0; i < len(hash); i++ {
		hash[i] = lyraHash[i]
	}

	return Hash(hash)
}

func X17Hash(b []byte) Hash {
	var dst [32]byte
	x17.New().Hash(b, dst[:])

	return Hash(dst)
}
