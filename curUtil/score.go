package curUtil

import moneyutil "github.com/lily0749labs/goutils/money"

// 分数转货币字符串
func ScoreToStrCur(score int64) string {
	return moneyutil.FormatCents(score)
}
