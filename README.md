# Enkodo (Marshal, unmarshal) [![GoDoc](https://godoc.org/github.com/mojura/enkodo?status.svg)](https://godoc.org/github.com/mojura/enkodo) ![Status](https://img.shields.io/badge/status-beta-yellow.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/mojura/enkodo)](https://goreportcard.com/report/github.com/mojura/enkodo)

Enkodo is a no-frills encoder/decoder. It is focused around speed and simplicity. 

*Note: Enkodo does not come with training wheels*

## Benchmarks
```bash
# Enkodo
BenchmarkEnkodoEncoding-4          10000000               141 ns/op               0 B/op          0 allocs/op
BenchmarkEnkodoDecoding-4          10000000               227 ns/op              16 B/op          1 allocs/op

# Binny (github.com/missionMeteora/binny.v2)
BenchmarkBinnyEncoding-4         5000000               343 ns/op              32 B/op          6 allocs/op
BenchmarkBinnyDecoding-4         5000000               390 ns/op              32 B/op          2 allocs/op

# JSON (standard library)
BenchmarkJSONEncoding-4         10000000               164 ns/op               0 B/op          0 allocs/op
BenchmarkJSONDecoding-4          5000000               273 ns/op               0 B/op          0 allocs/op
```

## Usage
```go
package main

import (
	"bytes"
	"log"

	"github.com/mojura/enkodo"
)

func main() {
	var (
		// Original user struct
		u User
		// New user struct (will be used to copy values to)
		nu  User
		err error
	)

	u.Email = "johndoe@gmail.com"
	u.Age = 46
	u.Twitter = "@johndoe"

	// Create a buffer to write to
	buf := bytes.NewBuffer(nil)
	// Create encoder
	enc := enkodo.NewEncoder(buf)
	// Encode user
	if err = enc.Encode(&u); err != nil {
		log.Fatalf("Error encoding: %v", err)
	}

	// Create decoder
	dec := enkodo.NewDecoder(buf)
	// Decode new user
	if err = dec.Decode(&nu); err != nil {
		log.Fatalf("Error decoding: %v", err)
	}

	log.Printf("New user: %v", nu)
}

// User holds the basic information for a user
type User struct {
	Email   string
	Age     uint8
	Twitter string
}

// MarshalEnkodo will marshal a User
func (u *User) MarshalEnkodo(enc *enkodo.Encoder) (err error) {
	if err = enc.String(u.Email); err != nil {
		return
	}

	if err = enc.Uint8(u.Age); err != nil {
		return
	}

	if err = enc.String(u.Twitter); err != nil {
		return
	}

	return
}

// UnmarshalEnkodo will unmarshal a User
func (u *User) UnmarshalEnkodo(dec *enkodo.Decoder) (err error) {
	if u.Email, err = dec.String(); err != nil {
		return
	}

	if u.Age, err = dec.Uint8(); err != nil {
		return
	}

	if u.Twitter, err = dec.String(); err != nil {
		return
	}

	return
}

```

## Features
- [x] Support for basic primitives
- [ ] Support for maps (currently needs helper func)
- [ ] Support for complex ints
- [x] Encoding via helper funcs
- [x] Decoding via helper funcs
- [ ] Encoding helper funcs created by reflection
- [ ] Decoding helper funcs created by reflection