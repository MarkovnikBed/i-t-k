package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"testing"
)

func TestGetTypes(t *testing.T) {
	equal(t, getTypes(1), "1 is int type")
	equal(t, getTypes(2.3), "2.3 is float64 type")
	equal(t, getTypes(true), "true is bool type")
	equal(t, getTypes("'some test string'"), "'some test string' is string type")
	equal(t, getTypes(2+2i), "(2+2i) is complex128 type")
	equal(t, getTypes(1, 2.3, true, "'some test string'", 2+2i), "1 is int type\n2.3 is float64 type\ntrue is bool type\n'some test string' is string type\n(2+2i) is complex128 type")
}

func equal(t *testing.T, str string, compareStr string) {
	str = strings.Trim(str, "\n")

	if str != compareStr {
		t.Logf("Expected %s but get '%s'", compareStr, str)
		t.Fail()
	}
}

func TestConcatinate(t *testing.T) {
	equal(t, concatinate(1, 2, 3, 4, 5, 6, 7), "1234567")
	equal(t, concatinate(1, 2.3, 1, 2, 1000, 100.5), "12.3121000100.5")
	equal(t, concatinate(0), "0")
	equal(t, concatinate(1, 2, "some string", false), "12some stringfalse")
	equal(t, concatinate(1, 10.5, "string", true, 25+1i, 2.567, 233), "110.5stringtrue(25+1i)2.567233")
}

func TestGetHash(t *testing.T) {
	validHash(t, getHash("hello", "salt"), "hello", "salt")
	validHash(t, getHash("hsdfgsdfgdsfgsdfg", "2342345o45uji245joi2j43pjori43jf"), "hsdfgsdfgdsfgsdfg", "2342345o45uji245joi2j43pjori43jf")
}

func validHash(t *testing.T, hash string, word string, salt string) {
	hashTmplt := sha256.Sum256(append(append([]byte(word[:len(word)/2]), []byte(salt)...), []byte(word[len(word)/2:])...))
	hashTmpltStr := hex.EncodeToString(hashTmplt[:])
	if hashTmpltStr != hash {
		t.Logf("hashs not equal")
		t.Fail()
	}
}
