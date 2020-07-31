package models

import (
	"blogweb_gin/dao"
	logger "blogweb_gin/gb"
	"fmt"
	sql "github.com/jmoiron/sqlx"
	"go.uber.org/zap"
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

var articleNum = 0

//-----------数据库操作---------------
func AddArticle(article *Article) (int64, error) {
	return dao.ModifyDB("insert into article(title,tags,short,content,author,create_time,status) values(?,?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.CreateTime, article.Status)
}

func UpdateArticle(article *Article) (int64, error) {
	sqlStr := "update article set title=?,tags=?,short=?,content=?,author=? where id=?"
	return dao.ModifyDB(sqlStr, article.Title, article.Tags, article.Short, article.Content, article.Author, article.Id)
}

func DeleteArticle(id string) (int64, error) {
	sqlStr := "delete from article where id=?"
	return dao.ModifyDB(sqlStr, id)
}
func FindArticleWithPage(page int) ([]*Article, error) {
	page--
	return QueryArticleWithPage(page, PageSize)
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
func GetArticleById(id int) (article *Article, err error) {
	article = new(Article)
	sqlStr := "select id,title,tags,short,content,author,create_time from article where id=?"
	err = dao.QueryRowDB(article, sqlStr, id)
	return
}
func QueryArticleByParam(param string) (tags []string, err error) {
	sqlStr := fmt.Sprintf("select %s from article", param)
	err = dao.QueryRows(&tags, sqlStr)
	return
}

// 根据查询条件查询指定页数有的文章
func queryArticleWithCon(pageNum int, sqlStr string, args ...interface{}) (articleList []*Article, err error) {
	pageNum--
	args = append(args, pageNum*PageSize, PageSize)
	logger.Debug("queryArticleWithCon", zap.Any("pageNum", pageNum), zap.Any("args", args))
	sqlStr += "limit ?,?"
	fmt.Println(sqlStr)
	err = dao.QueryRows(&articleList, sqlStr, args...)
	logger.Debug("dao.QueryRows result", zap.Any("articleList", articleList))
	return
}
func QueryArticlesWithTag(pageNum int, tag string) ([]*Article, error) {
	sql := " where tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"
	sql = fmt.Sprintf("select * from article %s", sql)
	return queryArticleWithCon(pageNum, sql)
}
func QueryArticleByIds(ids []int64) (articles []*Article, err error) {
	query, args, err := sql.In("select id, title from article where id in (?)", ids)
	if err != nil {
		logger.Error("QueryArticlesByIds", zap.Any("error", err))
		return nil, err
	}
	err = dao.QueryRows(&articles, query, args...)
	return
}
