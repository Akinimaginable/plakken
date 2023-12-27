package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	mathrand "math/rand"
	"strconv"
	"strings"
)

func GenerateUrl() string {
	listChars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, currentConfig.urlLength)
	for i := range b {
		b[i] = listChars[mathrand.Intn(len(listChars))]
	}

	return string(b)
}

func GenerateSecret() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Printf("Failed to generate secret")
	}

	return hex.EncodeToString(key)
}

func UrlExist(url string) bool {
	return db.Exists(ctx, url).Val() == 1
}

func VerifySecret(url string, secret string) bool {
	return secret == db.HGet(ctx, url, "secret").Val()
}

func parseIntBeforeSeparator(source *string, sep string) int { // return -1 if error, only accept positive number
	var value int
	var err error
	if strings.Contains(*source, sep) {
		value, err = strconv.Atoi(strings.Split(*source, sep)[0])
		if err != nil {
			log.Println(err)
			return -1
		}
		if value < 0 { // Only positive value is correct
			return -1
		}
		*source = strings.Join(strings.Split(*source, sep)[1:], "")
	}
	return value
}

func parseExpiration(source string) int { // return -1 if error
	var expiration int
	if source == "0" {
		return 0
	}
	expiration = 86400 * parseIntBeforeSeparator(&source, "d")
	if expiration < 0 {
		return -1
	}
	expiration += 3600 * parseIntBeforeSeparator(&source, "h")
	if expiration < 0 {
		return -1
	}
	expiration += 60 * parseIntBeforeSeparator(&source, "m")
	if expiration < 0 {
		return -1
	}
	expiration += parseIntBeforeSeparator(&source, "s")
	if expiration < 0 {
		return -1
	}

	return expiration
}
