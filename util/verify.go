/**
 * @Author: Hardews
 * @Date: 2023/10/9 11:45
 * @Description:
**/

package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"os"
)

var secret = os.Getenv("WEBHOOK_SECRET")

func VerifySignature(signature string, content []byte) bool {
	return signature == generateSignature(content)
}

func generateSignature(content []byte) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write(content)
	return "sha256=" + fmt.Sprintf("%x", h.Sum(nil))
}
