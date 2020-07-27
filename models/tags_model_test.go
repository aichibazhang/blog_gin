package models

import "testing"

func TestGetTags(t *testing.T) {
	tags := []string{"chine&dong", "dong&han"}
	tagMap := GetTags(tags)
	t.Log(tagMap)
}
