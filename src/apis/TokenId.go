package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"time"
)

func SHA1() string {

	h := sha1.New()
	io.WriteString(h, time.Now().String())
	return fmt.Sprintf("%x", h.Sum(nil))
}
