package commands

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"math/big"
)

var (
	Generators = func() []Command {
		var commands []Command

		commands = append(commands, hashGenerator{
			base:   NewBase("MD5 Hash Generator", ""),
			hasher: md5.New(),
		})

		commands = append(commands, hashGenerator{
			base:   NewBase("SHA1 Hash Generator", ""),
			hasher: sha1.New(),
		})

		commands = append(commands, hashGenerator{
			base:   NewBase("SHA256 Hash Generator", ""),
			hasher: sha256.New(),
		})

		commands = append(commands, hashGenerator{
			base:   NewBase("SHA512 Hash Generator", ""),
			hasher: sha512.New(),
		})

		commands = append(commands, secretGenerator{
			base:   NewBase("16 Character Secret", "").withoutInputDisplay(),
			length: 16,
		})

		commands = append(commands, secretGenerator{
			base:   NewBase("32 Character Secret", "").withoutInputDisplay(),
			length: 32,
		})

		commands = append(commands, secretGenerator{
			base:   NewBase("64 Character Secret", "").withoutInputDisplay(),
			length: 64,
		})

		return commands
	}
)

type hashGenerator struct {
	base
	hasher hash.Hash
}

func (g hashGenerator) Exec(input string) (string, error) {
	g.hasher.Reset()
	g.hasher.Write([]byte(input))
	return hex.EncodeToString(g.hasher.Sum(nil)), nil
}

type secretGenerator struct {
	base
	length int
}

func (g secretGenerator) Exec(string) (string, error) {
	// Shamelessly copied from: https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, g.length)
	for i := 0; i < g.length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret), nil
}
