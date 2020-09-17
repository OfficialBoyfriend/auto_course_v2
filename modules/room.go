package modules

import (
	"encoding/json"
	"errors"
	"github.com/rs/xid"
	bolt "go.etcd.io/bbolt"
)

const roomBoxName = "room"

type Room struct {
	Id        string
	Name      string        // 名称
	TimeTable [5][10]string // 时间表
}

func NewRoom() *Room {
	return new(Room)
}

func (r *Room) Put(fn ...func(func(txn *bolt.Tx) error) error) (err error) {
	run := NewDBBatch()
	if len(fn) > 0 {
		run = fn[0]
	}
	return run(func(tx *bolt.Tx) error {
		// 生成标识
		if r.Id == "" {
			guid := xid.New()
			r.Id = guid.String()
		}
		// 编码数据
		dataJson, err := json.Marshal(r)
		if err != nil {
			return err
		}
		// 打开存储桶
		box := tx.Bucket([]byte(roomBoxName))
		// 保存数据
		return box.Put([]byte(r.Id), dataJson)
	})
}

func (r *Room) Delete(fn ...func(func(txn *bolt.Tx) error) error) (err error) {
	if r.Id == "" {
		return errors.New("ID不能为空")
	}
	run := NewDBBatch()
	if len(fn) > 0 {
		run = fn[0]
	}
	return run(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(roomBoxName))
		// 删除数据
		return box.Delete([]byte(r.Id))
	})
}

func (r *Room) First() (err error) {
	return db.View(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(roomBoxName))
		// 读取数据
		item := box.Get([]byte(r.Id))
		// 解析数据
		return json.Unmarshal(item, r)
	})
}

func (r *Room) List() ([]*Room, error) {

	rooms := make([]*Room, 0)

	err := db.View(func(tx *bolt.Tx) error {
		// 打开存储桶
		box := tx.Bucket([]byte(roomBoxName))
		// 遍历键值
		c := box.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			room := new(Room)
			if err := json.Unmarshal(v, room); err != nil {
				return err
			}
			rooms = append(rooms, room)
		}
		return nil
	})

	return rooms, err
}
