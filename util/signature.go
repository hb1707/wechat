package util

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
)

// Signature sha1签名
func Signature(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		_, _ = io.WriteString(h, s)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func CalSignature(params ...string) string {
	sort.Strings(params)
	var buffer bytes.Buffer
	for _, value := range params {
		buffer.WriteString(value)
	}

	sha := sha1.New()
	sha.Write(buffer.Bytes())
	signature := fmt.Sprintf("%x", sha.Sum(nil))
	return string(signature)
}
