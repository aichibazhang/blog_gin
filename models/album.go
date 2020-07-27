package models

import "blogweb_gin/dao"

type Album struct {
	Id         int
	FilePath   string
	FileName   string
	Status     int
	CreateTime int64
}

//-----------数据库操作---------------
func AddAlbum(album *Album) (int64, error) {
	return dao.ModifyDB("insert into album(filepath,filename,status,createtime)values(?,?,?,?)",
		album.FilePath, album.FileName, album.Status, album.CreateTime)
}
func GetAlbums() (albums []*Album, err error) {
	sqlStr := "select id,filepath,filename,status,createtime from album"
	err = dao.QueryRows(&albums, sqlStr)
	return
}
