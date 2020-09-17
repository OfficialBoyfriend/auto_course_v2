package modules

import (
	"encoding/json"
	"errors"
	"github.com/rs/xid"
	bolt "go.etcd.io/bbolt"
)

const teacherBoxName = "teacher"

// 教师结构体
type Teacher struct {
	Id        string
	Name      string        // 名称
	TimeTable [5][10]string // 时间表
}

func NewTeacher() *Teacher {
	return new(Teacher)
}

func (t *Teacher) Put(fn ...func(func(txn *bolt.Tx) error) error) (err error) {
	run := NewDBBatch()
	if len(fn) > 0 {
		run = fn[0]
	}
	return run(func(tx *bolt.Tx) error {
		// 生成标识
		if t.Id == "" {
			guid := xid.New()
			t.Id = guid.String()
		}
		// 编码数据
		dataJson, err := json.Marshal(t)
		if err != nil {
			return err
		}
		// 打开存储桶
		box := tx.Bucket([]byte(teacherBoxName))
		// 保存数据
		return box.Put([]byte(t.Id), dataJson)
	})
}

func (t *Teacher) Delete(fn ...func(func(txn *bolt.Tx) error) error) (err error) {
	if t.Id == "" {
		return errors.New("ID不能为空")
	}
	run := NewDBBatch()
	if len(fn) > 0 {
		run = fn[0]
	}
	return run(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(teacherBoxName))
		// 删除数据
		return box.Delete([]byte(t.Id))
	})
}

func (t *Teacher) First() (err error) {
	return db.View(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(teacherBoxName))
		// 读取数据
		item := box.Get([]byte(t.Id))
		// 解析数据
		return json.Unmarshal(item, t)
	})
}

func (t *Teacher) List() ([]*Teacher, error) {

	teachers := make([]*Teacher, 0)

	err := db.View(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(teacherBoxName))
		// 遍历键值
		c := box.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			teacher := new(Teacher)
			if err := json.Unmarshal(v, teacher); err != nil {
				return err
			}
			teachers = append(teachers, teacher)
		}
		return nil
	})

	return teachers, err
}
