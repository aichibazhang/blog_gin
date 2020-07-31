package models

import (
	"strconv"
	"testing"
)

func TestGetTags(t *testing.T) {
	tags := []string{"chine&dong", "dong&han"}
	tagMap := GetTags(tags)
	t.Log(tagMap)
	i, err := strconv.ParseInt("123", 0, 32)
	if err != nil {
		panic(err)
	}
	t.Log(i)
}
