package authToken

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type AuthToken struct {
	key string
	ttl int64
}

// Create instance of AuthToken
func NewAuthToken(key string, ttl int64) *AuthToken {
	return &AuthToken{
		key: key,
		ttl: ttl,
	}
}

func (structure *AuthToken) CheckToken(token string, content string, ts time.Time) bool {
	if "" == strings.TrimSpace(token) {
		return false
	}
	tsn := time.Now()
	if tsn.Unix() < ts.Unix() || structure.ttl < tsn.Unix()-ts.Unix() {
		return false
	}

	return token == structure.NewToken(content, ts)
}

func (structure *AuthToken) NewToken(content string, ts time.Time) string {
	return structure.createToken(ts, structure.clearSpace(content))
}

func (structure *AuthToken) getContentHashBase64(content string) string {
	hasher := sha256.New()
	hasher.Write([]byte(content))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

func (structure *AuthToken) createToken(ts time.Time, content string) string {
	hc := hmac.New(sha256.New, []byte(structure.key))
	contentHash := structure.getContentHashBase64(content)
	stringToSign := fmt.Sprintf("%s;%d", contentHash, ts.Unix())
	hc.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(hc.Sum(nil))
}

func (structure *AuthToken) clearSpace(str string) string {
	r := regexp.MustCompile("\\s+")

	return r.ReplaceAllString(str, " ")
}
