package commands

import (
	"encoding/base64"
	"net/url"
)

var (
	EncodersDecoders = func() []Command {
		var converters []Command
		converters = append(converters, base64enc{
			base: NewBase("Base64 Encode", "Encode input to base64"),
		})
		converters = append(converters, base64dec{
			base: NewBase("Base64 Decode", "Decode input from base64"),
		})
		converters = append(converters, urlenc{
			base: NewBase("URL Encode", "Encode input to url"),
		})
		converters = append(converters, urldec{
			base: NewBase("URL Decode", "Decode input from url encoded"),
		})
		return converters
	}
)

type base64enc struct {
	base
}

func (e base64enc) Exec(raw string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(raw)), nil
}

type base64dec struct {
	base
}

func (e base64dec) Exec(raw string) (string, error) {
	out, err := base64.StdEncoding.DecodeString(raw)
	return string(out), err
}

type urlenc struct {
	base
}

func (e urlenc) Exec(raw string) (string, error) {
	return url.PathEscape(raw), nil
}

type urldec struct {
	base
}

func (e urldec) Exec(raw string) (string, error) {
	return url.PathUnescape(raw)
}
