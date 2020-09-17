package modules

import (
	"errors"
	"fmt"
)

// 班级结构体
type Class struct {
	Model
	Name    string        `gorm:"not null;unique"`       // 名称
	IsOk    bool          `json:"-"`                     // 是否已排课成功
	Courses []ClassCourse `gorm:"not null;"`             // 需要上的课程
	Days    []ClassDay    `gorm:"not null;" json:"days"` // 时间表 (课程安排) (星期\时间)
}

type ClassCourse struct {
	Model
	ClassId   uint    `gorm:"not null;" json:"-"` // 班级标识
	CourseId  uint    `gorm:"not null;" json:"-"` // 课程标识
	TeacherId uint    `gorm:"not null;" json:"-"` // 教师标识
	RoomId    uint    `gorm:"not null" json:"-"`  // 教室标识（系统自动编排）
	RoomTime  [2]uint `gorm:"not null" json:"-"`  // 教室使用时间
	Priority  uint    `json:"-"`                  // 优先级
}

// 天数结构体
// 时间表中使用
// 班级联名定制款
type ClassDay struct {
	Model
	ClassId uint           `gorm:"not null;" json:"-"`     // 班级标识
	Number  uint           `gorm:"not null;"`              // 记录天数
	Times   []ClassDayTime `gorm:"not null;" json:"times"` // 时间与班级课程绑定
}

// 班级课程结构体
type ClassDayTime struct {
	Model
	ClassDayId    uint `gorm:"not null;" json:"-"` // 所属天数标识
	Number        uint `gorm:"not null;"`          // 当天的时间编号
	ClassCourseId uint `json:"-"`                  // 该时间的班级课程
}

func (c *Class) Create() (err error) {

	if !db.NewRecord(c) {
		return errors.New("创建班级失败: 主键已存在")
	}

	defer func() {
		if err != nil || db.NewRecord(c) {
			err = fmt.Errorf("创建班级失败: %v", err)
		}
	}()

	return db.Create(c).Error
}

func (c *Class) First() (err error) {

	defer func() {
		if err != nil || db.NewRecord(c) {
			err = fmt.Errorf("读取班级失败: %v", err)
		}
	}()

	err = db.Preload("Days.Times").Preload("Courses").First(c).Error
	if err != nil {
		return err
	}

	return nil
}
