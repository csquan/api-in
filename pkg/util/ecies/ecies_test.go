// Copyright (c) 2013 Kyle Isom <kyle@tyrfingr.is>
// Copyright (c) 2012 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package ecies

import (
	"bytes"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
	"user/pkg/util"
)

func TestKDF(t *testing.T) {
	tests := []struct {
		length int
		output []byte
	}{
		{6, decode("858b192fa2ed")},
		{32, decode("858b192fa2ed4395e2bf88dd8d5770d67dc284ee539f12da8bceaa45d06ebae0")},
		{48, decode("858b192fa2ed4395e2bf88dd8d5770d67dc284ee539f12da8bceaa45d06ebae0700f1ab918a5f0413b8140f9940d6955")},
		{64, decode("858b192fa2ed4395e2bf88dd8d5770d67dc284ee539f12da8bceaa45d06ebae0700f1ab918a5f0413b8140f9940d6955f3467fd6672cce1024c5b1effccc0f61")},
	}

	for _, test := range tests {
		h := sha256.New()
		k := concatKDF(h, []byte("input"), nil, test.length)
		if !bytes.Equal(k, test.output) {
			t.Fatalf("KDF: generated key %x does not match expected output %x", k, test.output)
		}
	}
}

var ErrBadSharedKeys = fmt.Errorf("ecies: shared keys don't match")

// cmpParams compares a set of ECIES parameters. We assume, as per the
// docs, that AES is the only supported symmetric encryption algorithm.
func cmpParams(p1, p2 *Params) bool {
	return p1.hashAlgo == p2.hashAlgo &&
		p1.KeyLen == p2.KeyLen &&
		p1.BlockSize == p2.BlockSize
}

// Validate the ECDH component.
func TestSharedKey(t *testing.T) {
	prv1, err := GenerateKey(rand.Reader, DefaultCurve, nil)
	if err != nil {
		t.Fatal(err)
	}
	skLen := MaxSharedKeyLength(&prv1.PublicKey) / 2

	prv2, err := GenerateKey(rand.Reader, DefaultCurve, nil)
	if err != nil {
		t.Fatal(err)
	}

	sk1, err := prv1.GenerateShared(&prv2.PublicKey, skLen, skLen)
	if err != nil {
		t.Fatal(err)
	}

	sk2, err := prv2.GenerateShared(&prv1.PublicKey, skLen, skLen)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(sk1, sk2) {
		t.Fatal(ErrBadSharedKeys)
	}
}

func TestSharedKeyPadding(t *testing.T) {
	// sanity checks
	prv0 := hexKey("1adf5c18167d96a1f9a0b1ef63be8aa27eaf6032c233b2b38f7850cf5b859fd9")
	prv1 := hexKey("0097a076fc7fcd9208240668e31c9abee952cbb6e375d1b8febc7499d6e16f1a")
	//x0, _ := new(big.Int).SetString("1a8ed022ff7aec59dc1b440446bdda5ff6bcb3509a8b109077282b361efffbd8", 16)
	//x1, _ := new(big.Int).SetString("6ab3ac374251f638d0abb3ef596d1dc67955b507c104e5f2009724812dc027b8", 16)
	//y0, _ := new(big.Int).SetString("e040bd480b1deccc3bc40bd5b1fdcb7bfd352500b477cb9471366dbd4493f923", 16)
	//y1, _ := new(big.Int).SetString("8ad915f2b503a8be6facab6588731fefeb584fd2dfa9a77a5e0bba1ec439e4fa", 16)
	x0, _ := new(big.Int).SetString("894f0b45e976ff1d368ecb31aa5fdd47e3edb1b980b7d3bf7a7b543a5b2964a0", 16)
	x1, _ := new(big.Int).SetString("99a279d52118fffbaa3f2ac2d60f3bacf10e6cf86f46ee7f3b39b29ec78a94f2", 16)
	y0, _ := new(big.Int).SetString("c942f48766dc44c2a6e808691091de40d84f9b9df5394f6df99454a209d7843e", 16)
	y1, _ := new(big.Int).SetString("7ca3ebd2ea8ac4913c8c0c8cac4571316abab06b2a076caa9369a1ae7fd2d8ce", 16)

	if prv0.PublicKey.X.Cmp(x0) != 0 {
		t.Errorf("mismatched prv0.X:\nhave: %x\nwant: %x\n", prv0.PublicKey.X.Bytes(), x0.Bytes())
	}
	if prv0.PublicKey.Y.Cmp(y0) != 0 {
		t.Errorf("mismatched prv0.Y:\nhave: %x\nwant: %x\n", prv0.PublicKey.Y.Bytes(), y0.Bytes())
	}
	if prv1.PublicKey.X.Cmp(x1) != 0 {
		t.Errorf("mismatched prv1.X:\nhave: %x\nwant: %x\n", prv1.PublicKey.X.Bytes(), x1.Bytes())
	}
	if prv1.PublicKey.Y.Cmp(y1) != 0 {
		t.Errorf("mismatched prv1.Y:\nhave: %x\nwant: %x\n", prv1.PublicKey.Y.Bytes(), y1.Bytes())
	}

	// test shared secret generation
	sk1, err := prv0.GenerateShared(&prv1.PublicKey, 16, 16)
	if err != nil {
		t.Fatal(err)
	}

	sk2, err := prv1.GenerateShared(&prv0.PublicKey, 16, 16)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(sk1, sk2) {
		t.Fatal(ErrBadSharedKeys.Error())
	}
}

// Verify that the key generation code fails when too much key data is
// requested.
func TestTooBigSharedKey(t *testing.T) {
	prv1, err := GenerateKey(rand.Reader, DefaultCurve, nil)
	if err != nil {
		t.Fatal(err)
	}

	prv2, err := GenerateKey(rand.Reader, DefaultCurve, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = prv1.GenerateShared(&prv2.PublicKey, 32, 32)
	if err != util.ErrEccSharedKeyTooBig {
		t.Fatal("ecdh: shared key should be too large for curve")
	}

	_, err = prv2.GenerateShared(&prv1.PublicKey, 32, 32)
	if err != util.ErrEccSharedKeyTooBig {
		t.Fatal("ecdh: shared key should be too large for curve")
	}
}

// Benchmark the generation of P256 keys.
func BenchmarkGenerateKeyP256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := GenerateKey(rand.Reader, elliptic.P256(), nil); err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark the generation of P256 shared keys.
func BenchmarkGenSharedKeyP256(b *testing.B) {
	prv, err := GenerateKey(rand.Reader, elliptic.P256(), nil)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := prv.GenerateShared(&prv.PublicKey, 16, 16)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark the generation of S256 shared keys.
func BenchmarkGenSharedKeyS256(b *testing.B) {
	prv, err := GenerateKey(rand.Reader, DefaultCurve, nil)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := prv.GenerateShared(&prv.PublicKey, 16, 16)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Verify that an encrypted message can be successfully decrypted.
func TestEncryptDecrypt(t *testing.T) {
	prv1, err := GenerateKey(rand.Reader, DefaultCurve, nil)
	if err != nil {
		t.Fatal(err)
	}

	prv2, err := GenerateKey(rand.Reader, DefaultCurve, nil)
	if err != nil {
		t.Fatal(err)
	}

	pk2Str := prv2.String()
	pub2Str := prv2.PublicKey.String()

	message := []byte("Hello, world.")
	ct, err := Encrypt(&prv2.PublicKey, message)
	if err != nil {
		t.Fatal(err)
	}
	pub2, err := PublicFromString(pub2Str)
	if err != nil {
		t.Fatal(err)
	}
	pk2, err := PrivateFromString(pk2Str)
	if err != nil {
		t.Fatal(err)
	}
	ct2, err := Encrypt(pub2, message)
	if err != nil {
		t.Fatal(err)
	}
	pt2, err := pk2.Decrypt(ct2)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(pt2, message) {
		t.Fatal("ecies: plaintext doesn't match message")
	}

	pt, err := prv2.Decrypt(ct)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(pt, message) {
		t.Fatal("ecies: plaintext doesn't match message")
	}
	fmt.Println(pk2Str)
	fmt.Println(pub2Str)
	_, err = prv1.Decrypt(ct)
	if err == nil {
		t.Fatal("ecies: encryption should not have succeeded")
	}
}

type testCase struct {
	Curve    elliptic.Curve
	Name     string
	Expected *Params
}

var testCases = []testCase{
	{
		Curve:    elliptic.P256(),
		Name:     "P256",
		Expected: Aes128Sha256,
	},
	{
		Curve:    elliptic.P384(),
		Name:     "P384",
		Expected: Aes192Sha384,
	},
	{
		Curve:    elliptic.P521(),
		Name:     "P521",
		Expected: Aes256Sha512,
	},
}

// Test parameter selection for each curve, and that P224 fails automatic
// parameter selection (see README for a discussion of P224). Ensures that
// selecting a set of parameters automatically for the given curve works.
func TestParamSelection(t *testing.T) {
	for _, c := range testCases {
		testParamSelection(t, c)
	}
}

func testParamSelection(t *testing.T, c testCase) {
	params := ParamsFromCurve(c.Curve)
	if params == nil {
		t.Fatal("ParamsFromCurve returned nil")
	} else if params != nil && !cmpParams(params, c.Expected) {
		t.Fatalf("ecies: parameters should be invalid (%s)\n", c.Name)
	}

	prv1, err := GenerateKey(rand.Reader, DefaultCurve, nil)
	if err != nil {
		t.Fatal(err)
	}

	prv2, err := GenerateKey(rand.Reader, DefaultCurve, nil)
	if err != nil {
		t.Fatal(err)
	}

	message := []byte("Hello, world.")
	ct, err := Encrypt(&prv2.PublicKey, message)
	if err != nil {
		t.Fatal(err)
	}

	pt, err := prv2.Decrypt(ct)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(pt, message) {
		t.Fatalf("ecies: plaintext doesn't match message (%s)\n", c.Name)
	}

	_, err = prv1.Decrypt(ct)
	if err == nil {
		t.Fatalf("ecies: encryption should not have succeeded (%s)\n", c.Name)
	}
}

// Ensure that the basic public key validation in the decryption operation
// works.
func TestBasicKeyValidation(t *testing.T) {
	badBytes := []byte{0, 1, 5, 6, 7, 8, 9}

	prv, err := GenerateKey(rand.Reader, DefaultCurve, nil)
	if err != nil {
		t.Fatal(err)
	}

	message := []byte("Hello, world.")
	ct, err := Encrypt(&prv.PublicKey, message)
	if err != nil {
		t.Fatal(err)
	}

	for _, b := range badBytes {
		ct[0] = b
		_, err := prv.Decrypt(ct)
		if err != util.ErrEccInvalidPublicKey {
			t.Fatal("ecies: validated an invalid key")
		}
	}
}

func TestBox(t *testing.T) {
	prv1 := hexKey("4b50fa71f5c3eeb8fdc452224b2395af2fcc3d125e06c32c82e048c0559db03f")
	prv2 := hexKey("d0b043b4c5d657670778242d82d68a29d25d7d711127d17b8e299f156dad361a")
	pub2 := &prv2.PublicKey

	message := []byte("Hello, world.")
	ct, err := Encrypt(pub2, message)
	if err != nil {
		t.Fatal(err)
	}

	pt, err := prv2.Decrypt(ct)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(pt, message) {
		t.Fatal("ecies: plaintext doesn't match message")
	}
	if _, err = prv1.Decrypt(ct); err == nil {
		t.Fatal("ecies: encryption should not have succeeded")
	}
}

// Verify GenerateShared against static values - useful when
// debugging changes in underlying libs
func TestSharedKeyStatic(t *testing.T) {
	prv1 := hexKey("7ebbc6a8358bc76dd73ebc557056702c8cfc34e5cfcd90eb83af0347575fd2ad")
	prv2 := hexKey("6a3d6396903245bba5837752b9e0348874e72db0c4e11e9c485a81b4ea4353b9")

	skLen := MaxSharedKeyLength(&prv1.PublicKey) / 2

	sk1, err := prv1.GenerateShared(&prv2.PublicKey, skLen, skLen)
	if err != nil {
		t.Fatal(err)
	}

	sk2, err := prv2.GenerateShared(&prv1.PublicKey, skLen, skLen)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(sk1, sk2) {
		t.Fatal(ErrBadSharedKeys)
	}

	//	sk := decode("167ccc13ac5e8a26b131c3446030c60fbfac6aa8e31149d0869f93626a4cdf62")
	sk := decode("2b71c59bb0495360d20642360998981d3d00c74f6e72ec4d94f1391662f00d10")
	if !bytes.Equal(sk1, sk) {
		t.Fatalf("shared secret mismatch: want: %x have: %x", sk, sk1)
	}
}

func hexKey(prv string) *PrivateKey {
	key, err := HexToECDSA(prv)
	if err != nil {
		panic(err)
	}
	return ImportECDSA(key)
}

func decode(s string) []byte {
	byt, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return byt
}

func TestEciesCodec(t *testing.T) {
	KycPrivateKey := "377d84887abf772ec0992b39326e74f06fb4aff4b1a6bc455fb6e67fe7ffe9c6"
	KycPubKey := "026a38b8aa47f2af2d2163253ff5385cd059f90476990c7e3ee84bca0ead322241"

	pubKey, err := PublicFromString(KycPubKey)
	if err != nil {
		t.Fatal(err)
	}
	pk, err := PrivateFromString(KycPrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	nowStr := "Fri, 30 Dec 2022 03:04:38 GMT"
	ct, err := Encrypt(pubKey, []byte(nowStr))
	if err != nil {
		t.Fatal(err)
	}
	hex1 := hex.EncodeToString(ct)
	fmt.Println(hex1)

	hexStr := "0499e006ccc008ed7c88886262023a801d2547645736b3823bd06aa3aad729cfdc66d3564cd57fe73d4068bddf18dbb852c557c136922fb170f603d1658985f05647b3b5b02f9a1461669a24011925c66e9789f3058536048d547c13c565b724e663ab09067e6f2950b4767ff9e8c0a72be5308383e48af1d58cf477076518a3aa1126c623f33a40039c2b5b4be943f98b4914b67aab66e84e8a"
	ctb, er := hex.DecodeString(hexStr)
	if er != nil {
		t.Fatal(er)
	}
	now, err := pk.Decrypt(ctb)
	if err != nil {
		t.Fatal(err)
	}
	if nowStr != string(now) {
		t.Fatal("not eq")
	}
	fmt.Println(string(now))
}

func TestJavaKey(t *testing.T) {
	pkStr := "95d3c5e483e9b1d4f5fc8e79b2deaf51362980de62dbb082a9a4257eef653d7d"
	pk, err := PrivateFromString(pkStr)
	if err != nil {
		t.Error(err)
	}
	pubStr := pk.PublicKey.String()
	fmt.Println(pubStr)
}
