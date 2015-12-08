package models

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Uname    string
	Upw      string
	Showname string
}

func (this *User) Login() (r bool, err error) {
	r = false
	var spw string
	db, err := sql.Open("sqlite3", "./neo.s3db")
	if err != nil {
		fmt.Println("打开数据库出错" + err.Error())
		return false, err
	}
	fmt.Println("连接数据库成功")
	rows, err := db.Query("select upw from t_user where uname=?", this.Uname)
	if err != nil {
		fmt.Println("执行query出错" + err.Error())
		return false, err
	}
	if rows.Next() {
		rows.Scan(&spw)
		h := md5.New()
		h.Write([]byte(this.Upw))
		if spw == hex.EncodeToString(h.Sum(nil)) {
			r = true
		}
	}
	rows.Close()
	db.Close()
	return r, err
}
