package utils

import "strings"

func IsEmptyOrWhitespace(s string) bool {
	// 全角空白を取り除く
	s = strings.Replace(s, "\u3000", "", -1)
	// 半角空白を取り除く
	s = strings.Replace(s, " ", "", -1)

	// 空文字列かどうかチェック
	return s == ""
}
