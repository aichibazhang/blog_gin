package models

import (
	"blogweb_gin/dao"
	"fmt"
)

const (
	PageSize = 4
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	CreateTime int64 `db:"create_time"`
	Status     int   // Status=0为正常，1为删除，2为冻结
}
var articleNum=0
//-----------数据库操作---------------
func AddArticle(article *Article) (int64, error) {
	return dao.ModifyDB("insert into article(title,tags,short,content,author,create_time,status) values(?,?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.CreateTime, article.Status)
}

func UpdateArticle(article *Article) (int64, error) {
	sqlStr := "update article set title=?,tags=?,short=?,content=? where id=?"
	return dao.ModifyDB(sqlStr, article.Title, article.Tags, article.Short, article.Content, article.Id)
}

func DeleteArticle(id string) (int64, error) {
	sqlStr := "delete from article where id=?"
	return dao.ModifyDB(sqlStr, id)
}
func FindArticleWithPage(page int) ([]*Article, error) {
	page--
	return QueryArticleWithPage(page,PageSize)
}
func GetArticleRowsNum() int {
	articleNum = 0
	dao.QueryRowDB(&articleNum, "select count(id) from article")
	return articleNum
}

func QueryArticleWithPage(page int, num int) (artList []*Article, err error) {
	sql := fmt.Sprintf("limit %d,%d", page*num, num)
	sql = "select id,title,tags,short,content,author,create_time from article " + sql
	fmt.Println("分页查询sql:", sql)
	err = dao.QueryRows(&artList, sql)
	fmt.Println("分页查询出文章有:", len(artList))
	return
}
