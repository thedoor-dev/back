package db

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/thedoor-dev/back/models"
	"github.com/thedoor-dev/back/utils"
)

var pIDGen *utils.SnowFlake
var tIDGen *utils.SnowFlake

func init() {
	pIDGen = &utils.SnowFlake{}
	tIDGen = &utils.SnowFlake{}
}

func PostList(pLs *[]models.PostList, start, len int) (err error) {
	err = db.Select(
		pLs,
		"select pid, title, abstract, ctime\n"+
			"from posts\n"+
			"WHERE `public` = 1\n"+
			"ORDER BY `pid` DESC\n"+
			"LIMIT ?\n"+
			"OFFSET ?\n",
		len,
		start,
	)
	if err != nil {
		return err
	}
	return
}

func PostListTags(tags *models.TagArr, pid int64) (err error) {
	return db.Select(
		tags,
		"SELECT * FROM tags WHERE pid = ?",
		pid,
	)
}

func PostNew(p *models.Post) error {
	if p.Article == "" {
		return errors.New("？？？？博客内容呢")
	}
	p.ID = pIDGen.GetVal()
	p.CreateTime = time.Now()
	if p.Abstract == "" {
		if len(p.Abstract) > 10 {
			p.Abstract = p.Article[:10]
		} else {
			p.Abstract = p.Article
		}
	}
	_, err := db.Exec(
		"INSERT INTO `posts`(`pid`, `abstract`, `article`, `public`, `ctime`) VALUES"+
			"(?, ?, ?, ?, ?)",
		p.ID, p.Abstract, p.Article, p.Public, p.CreateTime,
	)
	return err
}

func PostNewWithTag(title, abstract, article string, public bool, tags []string) (err error) {
	if title == "" || abstract == "" || article == "" || len(tags) == 0 {
		return errors.New("博客呢？？？？")
	}
	var tx *sqlx.Tx
	tx, err = db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		p := recover()
		if p != nil {
			tx.Rollback()
			err = errors.New("recover error")
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	pid := pIDGen.GetVal()
	ctime := time.Now()
	_, err = tx.Exec("INSERT INTO `posts`(`pid`, `title`, `abstract`, `article`, `public`, `ctime`) VALUES"+
		"(?, ?, ?, ?, ?, ?)",
		pid, title, abstract, article, public, ctime,
	)
	if err != nil {
		return err
	}
	for _, v := range tags {
		_, err = tx.Exec("INSERT INTO `tags`(`tid`, `pid`, `name`) VALUES(?, ?, ?)",
			tIDGen.GetVal(), pid, v,
		)
		if err != nil {
			return err
		}
	}
	return
}

func PostOne(p *models.Post, pid int64, public bool) (err error) {
	if public {
		err = db.Get(
			p,
			"SELECT `pid`, `article`, `ctime`\n"+
				"FROM `posts`\n"+
				"WHERE `pid` = ? AND public = ?",
			pid, public,
		)
	} else {
		err = db.Get(
			p,
			"SELECT `pid`, `article`, `ctime`\n"+
				"FROM `posts`\n"+
				"WHERE `pid` = ?",
			pid,
		)
	}
	return err
}
