package commands

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	"html"
	"github.com/golang-jwt/jwt"
	"fmt"
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
		converters = append(converters, htmlenc{
			base: NewBase("HTML Escape", "Encode input to HTML escape characters"),
		})
		converters = append(converters, htmldec{
			base: NewBase("HTML Unescape", "Decode input from HTML escape characters"),
		})
		converters = append(converters, jwtdec{
			base: NewBase("JWT Decode", ""),
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

type htmlenc struct { 
	base 
}

func (e htmlenc) Exec(raw string) (string, error) {
	return html.EscapeString(raw), nil
}

type htmldec struct { 
	base 
}

func (e htmldec) Exec(raw string) (string, error) {
	return html.UnescapeString(raw), nil
}

type jwtdec struct {
	base
}

func (e jwtdec) Exec(raw string) (string, error) {
	token, _ := jwt.Parse(raw, nil)
	if token == nil {
		return "", fmt.Errorf("Failed parsing token")
	}
	headers, err := json.Marshal(token.Header)
	if err != nil {
		return "", err
	}
	claims, err := json.Marshal(token.Claims)
	if err != nil {
		return "", err
	}
	prettyHeaders, err := prettyJson(string(headers))
	if err != nil {
		return "", err
	}
	prettyClaims, err := prettyJson(string(claims))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Header:\n%s\n\nClaims:\n%s",prettyHeaders,prettyClaims), nil
}
