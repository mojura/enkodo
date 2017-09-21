package main

import (
	"bytes"
	"log"

	"github.com/itsmontoya/mum"
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
	enc := mum.NewEncoder(buf)
	// Encode user
	if err = enc.Encode(&u); err != nil {
		log.Fatalf("Error encoding: %v", err)
	}

	// Create decoder
	dec := mum.NewDecoder(buf)
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

// MarshalMum will marshal a User
func (u *User) MarshalMum(enc *mum.Encoder) (err error) {
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

// UnmarshalMum will unmarshal a User
func (u *User) UnmarshalMum(dec *mum.Decoder) (err error) {
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
