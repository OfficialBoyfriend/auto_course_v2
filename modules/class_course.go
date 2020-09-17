package modules

import (
	"encoding/json"
	"errors"
	"github.com/rs/xid"
	bolt "go.etcd.io/bbolt"
)

const classCourseBoxName = "class_course"

type ClassCourse struct {
	Id        string
	Number    uint   // 课时（课程数量）
	CourseId  string // 课程标识
	TeacherId string // 教师标识
	RoomId    string // 教室标识（系统自动编排）
	RoomTime  [2]int // 教室使用时间（系统自动编排）
}

func NewClassCourse() *ClassCourse {
	return new(ClassCourse)
}

func (c *ClassCourse) Put(fn ...func(func(txn *bolt.Tx) error) error) (err error) {
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
		box := tx.Bucket([]byte(classCourseBoxName))
		// 保存数据
		return box.Put([]byte(c.Id), dataJson)
	})
}

func (c *ClassCourse) Delete(fn ...func(func(txn *bolt.Tx) error) error) (err error) {
	if c.Id == "" {
		return errors.New("ID不能为空")
	}
	run := NewDBBatch()
	if len(fn) > 0 {
		run = fn[0]
	}
	return run(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(classCourseBoxName))
		// 删除数据
		return box.Delete([]byte(c.Id))
	})
}

func (c *ClassCourse) First() (err error) {
	return db.View(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(classCourseBoxName))
		// 读取数据
		item := box.Get([]byte(c.Id))
		// 解析数据
		return json.Unmarshal(item, c)
	})
}

func (c *ClassCourse) List() ([]*ClassCourse, error) {

	classCourses := make([]*ClassCourse, 0)

	err := db.View(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(classCourseBoxName))
		// 遍历键值
		c := box.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			classCourse := new(ClassCourse)
			if err := json.Unmarshal(v, classCourse); err != nil {
				return err
			}
			classCourses = append(classCourses, classCourse)
		}
		return nil
	})

	return classCourses, err
}
