package models

import (
	"fmt"
	"strconv"
)

type HomeFooterPage struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}

func ConfigHomeFooterPageCode(page int) HomeFooterPage {
	pageCode := HomeFooterPage{}
	records := GetArticleRowsNum()
	total := (records-1)/PageSize + 1
	pageCode.ShowPage = fmt.Sprintf("当前页:%d,共:%d", page, total)
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}
	if page >= total {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode
}
