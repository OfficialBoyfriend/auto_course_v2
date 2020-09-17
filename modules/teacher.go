package modules

import (
	"errors"
	"fmt"
)

// 教师结构体
type Teacher struct {
	Model
	Name string `gorm:"not null;unique"` // 名称
	Days []Day  `gorm:"not null"`
}

func (r *Teacher) Create() (err error) {

	if !db.NewRecord(r) {
		return errors.New("创建教师失败: 主键已存在")
	}

	defer func() {
		if err != nil || db.NewRecord(r) {
			err = fmt.Errorf("创建教师失败: %v", err)
		}
	}()

	return db.Create(r).Error
}

func (r *Teacher) First() (err error) {

	defer func() {
		if err != nil || db.NewRecord(r) {
			err = fmt.Errorf("读取教师失败: %v", err)
		}
	}()

	return db.Preload("Days.Times").First(r).Error
}
