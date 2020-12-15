package utils

import "encoding/base64"

/**
 * 对msg字符串进行base64编码
 */
func Base64Str(msg string)string  {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}