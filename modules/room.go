package modules

import (
	"errors"
	"fmt"
)

type Room struct {
	Model
	Name string `gorm:"not null;unique"` // 名称
	Days []Day  `gorm:"not null"`        // 时间表
}

func (r *Room) Create() (err error) {

	if !db.NewRecord(r) {
		return errors.New("创建教室失败: 主键已存在")
	}

	defer func() {
		if err != nil || db.NewRecord(r) {
			err = fmt.Errorf("创建教室失败: %v", err)
		}
	}()

	return db.Create(r).Error
}

func (r *Room) First() (err error) {

	defer func() {
		if err != nil || db.NewRecord(r) {
			err = fmt.Errorf("读取教室失败: %v", err)
		}
	}()

	return db.Preload("Days.Times").First(r).Error
}

/*
func (r *Room) Create() (err error) {

	if !db.NewRecord(r) {
		return errors.New("创建教室失败: 主键已存在")
	}

	defer func() {
		if err != nil || db.NewRecord(r) {
			err = fmt.Errorf("创建教室失败: %v", err)
		}
	}()

	****** 编码时间表 ******

	var (
		timerJson []byte
		timerIds  [][]uint
	)

	if r.Timer == nil {
		goto CREATE
	}

	for x := range r.Timer {
		yResult := make([]uint, len(r.Timer[x]))
		for y := range r.Timer[x] {
			yResult[y] = r.Timer[x][y].ID
		}
		timerIds = append(timerIds, yResult)
	}

	timerJson, err = json.Marshal(timerIds)
	if err != nil {
		return err
	}
	r.TimerJson = string(timerJson)

CREATE:
	return db.Create(r).Error
}

func (r *Room) First() (err error) {

	defer func() {
		if err != nil || db.NewRecord(r) {
			err = fmt.Errorf("读取教室失败: %v", err)
		}
	}()

	if err = db.First(r).Error; err != nil {
		return err
	}

	****** 解码时间表 ******

	if r.TimerJson == "" {
		return nil
	}

	// 解析JSON数据
	var timerIndex [][]uint
	err = json.Unmarshal([]byte(r.TimerJson), &timerIndex)
	if err != nil {
		return err
	}

	for x := range timerIndex {
		yResult := make([]Class, len(timerIndex[x]))
		for y := range timerIndex[x] {
			// 排除空数据
			if timerIndex[x][y] == 0 {
				continue
			}
			// 查询班级信息
			class := Class{
				Model: Model{ID: timerIndex[x][y]},
			}
			if err = class.First(); err != nil {
				return err
			}
			yResult[y] = class
		}
		r.Timer = append(r.Timer, yResult)
	}

	return nil
}
*/
