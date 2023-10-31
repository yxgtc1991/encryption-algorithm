package token_sha256

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestTokenSha256(t *testing.T) {
	appName := "app-test"
	version := "8.0.30"
	salt := "james" // 盐值：降低由于用户数据被盗而带来的密码泄露风险
	input := appName + version + salt
	sha256Obj := sha256.New()
	sha256Obj.Write([]byte(input))
	result := sha256Obj.Sum(nil)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	token := hex.EncodeToString(result) + timestamp // 有效期校验（3s）
	fmt.Println("token值：", token)
}
