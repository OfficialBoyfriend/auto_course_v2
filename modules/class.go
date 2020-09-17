package modules

import (
	"encoding/json"
	"errors"
	"github.com/rs/xid"
	bolt "go.etcd.io/bbolt"
)

const classBoxName = "class"

// 班级结构体
type Class struct {
	Id        string
	Name      string         // 名称
	IsOk      bool           // 是否已排课成功
	Courses   []*ClassCourse // 需要上的课程
	TimeTable [5][10]string  // 时间表 (课程安排) (星期\时间)
}

func NewClass() *Class {
	return new(Class)
}

func (c *Class) Put(fn ...func(func(txn *bolt.Tx) error) error) (err error) {
	run := NewDBBatch()
	if len(fn) > 0 {
		run = fn[0]
	}
	return run(func(tx *bolt.Tx) error {
		// 生成标识
		if c.Id == "" {
			guid := xid.New()
			c.Id = guid.String()
		}
		// 编码数据
		dataJson, err := json.Marshal(c)
		if err != nil {
			return err
		}
		// 打开存储桶
		box := tx.Bucket([]byte(classBoxName))
		// 保存数据
		return box.Put([]byte(c.Id), dataJson)
	})
}

func (c *Class) Delete(fn ...func(func(txn *bolt.Tx) error) error) (err error) {
	if c.Id == "" {
		return errors.New("ID不能为空")
	}
	run := NewDBBatch()
	if len(fn) > 0 {
		run = fn[0]
	}
	return run(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(classBoxName))
		// 删除数据
		return box.Delete([]byte(c.Id))
	})
}

func (c *Class) First() (err error) {
	return db.View(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(classBoxName))
		// 读取数据
		item := box.Get([]byte(c.Id))
		// 解析数据
		return json.Unmarshal(item, c)
	})
}

func (c *Class) List() ([]*Class, error) {

	classes := make([]*Class, 0)

	err := db.View(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(classBoxName))
		// 遍历键值
		c := box.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			class := new(Class)
			if err := json.Unmarshal(v, class); err != nil {
				return err
			}
			classes = append(classes, class)
		}
		return nil
	})

	return classes, err
}
