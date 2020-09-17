package modules

import (
	"errors"
	"fmt"
)

// 课程结构体
type Course struct {
	Model
	Name     string    `gorm:"not null;unique"`                     // 名称
	Teachers []Teacher `gorm:"not null;many2many:course_teachers;"` // 可选教师
	Rooms    []Room    `gorm:"not null;many2many:course_rooms;"`    // 可选教室
}

func (c *Course) Create() (err error) {
	if !db.NewRecord(c) {
		return errors.New("创建课程失败: 主键已存在")
	}
	defer func() {
		if err != nil || db.NewRecord(c) {
			err = fmt.Errorf("创建课程失败: %v", err)
		}
	}()
	return db.Create(c).Error
}

func (c *Course) First() (err error) {

	defer func() {
		if err != nil || db.NewRecord(c) {
			err = fmt.Errorf("读取课程失败: %v", err)
		}
	}()

	/*
		err = db.First(c).Error
		if err != nil {
			return err
		}
	*/

	// 加载关联数据
	//c.Teachers = teachers
	//c.Rooms = rooms
	//db.Model(c).Related(&c.Teachers, "teachers")
	//db.Model(c).Related(&c.Rooms, "rooms")
	//db.Model(c).Related(&c.Teachers, "teachers")
	//db.Model(c).Related(&c.Rooms)

	// 查询数据
	return db.Preload("Teachers.Days.Times").Preload("Rooms.Days.Times").First(c).Error
	//return nil
}
