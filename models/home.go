package models

import (
	"strconv"
	"strings"
)

type HomeBlockParam struct {
	Article *Article
	Tags    []*TagLink
	//查看文章的地址
	Link string

	//修改文章的地址
	UpdateLink string
	DeleteLink string

	//记录是否登录
	IsLogin bool
}

type TagLink struct {
	TagName string
	TagUrl  string
}

func GenHomeBlocks(articleList []*Article, isLogin bool) (ret []HomeBlockParam) {
	//为了节省资源,一开始就申请好内存
	ret = make([]HomeBlockParam, 0, len(articleList))
	for _, art := range articleList {
		homeParam := HomeBlockParam{
			Article: art,
			IsLogin: isLogin,
		}
		homeParam.Tags = createTagsLinks(art.Tags)
		homeParam.Link = "/show/" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		ret = append(ret, homeParam)
	}
	return ret
}
func createTagsLinks(tagStr string) []*TagLink {
	var tagLinks = make([]*TagLink, 0, strings.Count(tagStr, "&"))
	tagList := strings.Split(tagStr, "&")
	for _, tag := range tagList {
		tagLinks = append(tagLinks, &TagLink{tag, "/?tag=" + tag})
	}
	return tagLinks
}
