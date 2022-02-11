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
	"strings"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
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

		commands = append(commands, uuidGenerator{
			base:   NewBase("UUID v1", "date-time and mac address").withoutInputDisplay(),
			random: false,
		})

		commands = append(commands, uuidGenerator{
			base:   NewBase("UUID v4", "randoms").withoutInputDisplay(),
			random: false,
		})

		commands = append(commands, uuidGenerator{
			base: NewBase("Nil UUID", "zero").withoutInputDisplay(),
			zero: true,
		})

		commands = append(commands, lipsumGenerator{
			base:       NewBase("Lipsum - 1 Paragraph", "").withoutInputDisplay(),
			paragraphs: 1,
		})
		commands = append(commands, lipsumGenerator{
			base:       NewBase("Lipsum - 2 Paragraphs", "").withoutInputDisplay(),
			paragraphs: 2,
		})
		commands = append(commands, lipsumGenerator{
			base:       NewBase("Lipsum - 3 Paragraphs", "").withoutInputDisplay(),
			paragraphs: 3,
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

type uuidGenerator struct {
	base
	// When set to true UUID4, otherwise UUID1
	random bool
	// When set to true will be all Zeros
	zero bool
}

func (g uuidGenerator) Exec(string) (string, error) {
	if g.zero {
		return "00000000-0000-0000-0000-000000000000", nil
	}
	if g.random {
		res, err := uuid.NewRandom()
		return res.String(), err
	} else {
		res, err := uuid.NewUUID()
		return res.String(), err
	}
}

type lipsumGenerator struct {
	base
	paragraphs int
}

func (g lipsumGenerator) Exec(string) (string, error) {
	var lipsum []string

	for p := 0; p < g.paragraphs; p++ {
		lipsum = append(lipsum, faker.Paragraph())
	}

	return strings.Join(lipsum, "\n\n"), nil
}
