package models

import (
	"blogweb_gin/dao"
	"testing"
)

func TestQueryArticleWithPage(t *testing.T) {
	dao.InitMySQL()
	var artList []Article
	sql := "select id,title,tags,short,content,author,create_time from article limit 0,4"
	_ = dao.QueryRows(&artList, sql)
	num := GetArticleRowsNum()
	t.Log("分页查询出文章有:", artList)
	t.Log("分页查询出文章个数有:", num)
	t.Log(2/3)
}
