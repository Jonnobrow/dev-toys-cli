package commands

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
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
