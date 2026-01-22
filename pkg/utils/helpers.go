package utils

import "strings"

func MaskSensitiveInfo(s string, start int, maskNumber int, maskChars ...string) string {
	maskChar := "*"
	if maskChars != nil {
		maskChar = maskChars[0]
	}
	
	if start < 0 {
		start = 0
	}

	end := min(start + maskNumber, len(s))
	return s[:start] + strings.Repeat(maskChar, end-start) + s[end:]
}
