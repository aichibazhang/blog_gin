package models

import (
	"blogweb_gin/config"
	"blogweb_gin/dao"
	logger "blogweb_gin/gb"
	"testing"
)

func TestQueryArticleWithPage(t *testing.T) {
	if err := config.Init("../conf/conf.json"); err != nil {
		return
	}
	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		return
	}
	if err := dao.InitMySQL(config.Conf.MySQLConfig); err != nil {
		t.Logf("init redis failed, err:%v\n", err)
		return
	}
	//var artList []Article
	//sql := "select id,title,tags,short,content,author,create_time from article limit 0,4"
	//_ = dao.QueryRows(&artList, sql)
	//num := GetArticleRowsNum()
	info, _ := GetArticleById(1)
	//t.Log("分页查询出文章有:", artList)
	//t.Log("分页查询出文章个数有:", num)
	t.Log("文章详情", info)
	ids := []int64{1, 2, 3}
	articleList, _ := QueryArticleByIds(ids)
	t.Log(articleList[0])
}
