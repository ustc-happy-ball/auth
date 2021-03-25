package model

import (
	"time"
)

type Account struct {
	ID          int64
	MobilePhone string
	Name        string
	LoginPassword    string
	AccountAvatar      string
	Level       int
	Skin        string
	Deleted     bool
	Region      string
	QQ          string
	WeChat      string
	CreateAt    time.Time
	UpdateAt    time.Time
}

//func init () {
//	var once sync.Once
//	once.Do(func() {
//		createDB(config.DB,config.COLLECTION)
//	})
//}


