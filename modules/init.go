package modules

import (
	"auto_course/utils"
	bolt "go.etcd.io/bbolt"
	"log"
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

var db *bolt.DB

func NewDBBatch() func(func(txn *bolt.Tx) error) error {
	return db.Batch
}

func tmpInit() {

	var err error

	// 打开数据库
	db, err = bolt.Open("data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	// 创建存储桶
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte("class"))
		utils.CheckError(err)
		_, err = tx.CreateBucketIfNotExists([]byte("class_course"))
		utils.CheckError(err)
		_, err = tx.CreateBucketIfNotExists([]byte("course"))
		utils.CheckError(err)
		_, err = tx.CreateBucketIfNotExists([]byte("room"))
		utils.CheckError(err)
		_, err = tx.CreateBucketIfNotExists([]byte("teacher"))
		utils.CheckError(err)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func tmpClose() {
	db.Close()
}
