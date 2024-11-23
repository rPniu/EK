package id

import "regexp"

func ExtractIDFromURL(URL string) string {

	re := regexp.MustCompile(`id=([^&]+)`)
	// 提取匹配结果
	match := re.FindStringSubmatch(URL)
	if match == nil {
		return ""
	}
	return match[1]
}
