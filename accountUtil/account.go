// Package accountUtil 提供游戏账号相关的生成工具。
package accountUtil

import (
	"fmt"
	"strings"

	cryptoutil "github.com/lily0749labs/goutils/crypto"
	randutil "github.com/lily0749labs/goutils/rand"
	timeutil "github.com/lily0749labs/goutils/time"
)

const nicknamePrefix = "PLAY__"

const nicknameChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// GenerateNickname 生成 PLAY__ 前缀加 9 个随机英文字母的玩家昵称。
func GenerateNickname() string {
	r := randutil.Rand.GetRand()
	var nickname strings.Builder
	nickname.Grow(len(nicknamePrefix) + 9)
	nickname.WriteString(nicknamePrefix)
	for range 9 {
		nickname.WriteByte(nicknameChars[r.Intn(len(nicknameChars))])
	}
	return nickname.String()
}

// GenerateSMSCode 生成 [100000, 999998] 范围内的 6 位短信验证码。
// 上界与迁移前 RandSmsCode 的算法保持一致。
func GenerateSMSCode() int64 {
	r := randutil.Rand.GetRand()
	return int64(r.Intn(899999) + 100000)
}

// GeneratePassword 使用用户 ID 和当前 Unix 秒构造初始明文，并返回 bcrypt 哈希。
// 保留旧业务行为：bcrypt 失败时返回空字符串。
func GeneratePassword(userID uint64) string {
	return GeneratePasswordAt(userID, timeutil.Time.NowUnix())
}

// GeneratePasswordAt 使用指定 Unix 秒生成初始密码，便于可重复测试或迁移历史逻辑。
func GeneratePasswordAt(userID uint64, unixSeconds int64) string {
	plainText := fmt.Sprintf("%d_generate_%d", userID, unixSeconds)
	return cryptoutil.Crypto.BcryptHash(plainText)
}
