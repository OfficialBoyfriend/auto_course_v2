package modules

import (
	"encoding/json"
	"errors"
	"github.com/rs/xid"
	bolt "go.etcd.io/bbolt"
)

const courseBoxName = "course"

// 课程结构体
type Course struct {
	Id         string
	Name       string   // 名称
	TeacherIds []string // 可选教师
	RoomIds    []string // 可选教室
}

func NewCourse() *Course {
	return new(Course)
}

func (c *Course) Put(fn ...func(func(txn *bolt.Tx) error) error) (err error) {
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
		box := tx.Bucket([]byte(courseBoxName))
		// 保存数据
		return box.Put([]byte(c.Id), dataJson)
	})
}

func (c *Course) Delete(fn ...func(func(txn *bolt.Tx) error) error) (err error) {
	if c.Id == "" {
		return errors.New("ID不能为空")
	}
	run := NewDBBatch()
	if len(fn) > 0 {
		run = fn[0]
	}
	return run(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(courseBoxName))
		// 删除数据
		return box.Delete([]byte(c.Id))
	})
}

func (c *Course) First() (err error) {
	return db.View(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(courseBoxName))
		// 读取数据
		item := box.Get([]byte(c.Id))
		// 解析数据
		return json.Unmarshal(item, c)
	})
}

func (c *Course) List() ([]*Course, error) {

	courses := make([]*Course, 0)

	err := db.View(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(courseBoxName))
		// 遍历键值
		c := box.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			course := new(Course)
			if err := json.Unmarshal(v, course); err != nil {
				return err
			}
			courses = append(courses, course)
		}
		return nil
	})

	return courses, err
}
