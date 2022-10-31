package services

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("ABCDEFGHJKLMNPRSTUVWXYZ")
var numericRunes = []rune("23456789")
var runes = []rune("23456789ABCDEFGHJKLMNPRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RandAlphaNumericRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func RandNumericRunes(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = numericRunes[rand.Intn(len(numericRunes))]
	}
	return string(b)
}

func RandomPrefixUniqID() (string, int) {
	rand.Seed(time.Now().UnixNano())

	prePrefix := ""
	prefix := RandStringRunes(1)
	suffix := RandStringRunes(1)
	rand.Seed(time.Now().UnixNano())
	prePrefixNum := rand.Intn(6-1) + 1
	if prePrefixNum > 0 {
		prePrefix = RandNumericRunes(prePrefixNum)
	}

	midleNum := RandNumericRunes(rand.Intn((6-prePrefixNum)-1) + 1)
	finalPrefix := prePrefix + prefix + midleNum + suffix
	return finalPrefix, len(finalPrefix)
}
