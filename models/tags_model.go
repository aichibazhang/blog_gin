package models

import "strings"

func GetTags(tags []string) (tagMap map[string]int) {
	tagMap = make(map[string]int)
	for _, tag := range tags {
		tagList := strings.Split(tag, "&")
		for _, value := range tagList {
			(tagMap)[value]++
		}
	}
	return tagMap
}
