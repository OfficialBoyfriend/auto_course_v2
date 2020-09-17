package modules

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

var db *gorm.DB

func init() {
	var err error

	db, err = gorm.Open("sqlite3", "main.db")
	if err != nil {
		panic("连接数据库失败")
	}
	// TODO: debug
	//defer db.Close()

	// 自动创建表(如果不存在)
	autoCreateTable()

	// 班级相关
	db.AutoMigrate(new(Class))
	db.AutoMigrate(new(ClassCourse))
	db.AutoMigrate(new(ClassDay))
	db.AutoMigrate(new(ClassDayTime))

	// 基础信息
	db.AutoMigrate(new(Course))
	db.AutoMigrate(new(Room))
	db.AutoMigrate(new(Teacher))

	// 时间表
	db.AutoMigrate(new(Day))
	db.AutoMigrate(new(DayTime))
}

func autoCreateTable() {
	// 班级相关
	if !db.HasTable(new(Class)) {
		db.CreateTable(new(Class))
	}
	if !db.HasTable(new(ClassCourse)) {
		db.CreateTable(new(ClassCourse))
	}
	if !db.HasTable(new(ClassDay)) {
		db.CreateTable(new(ClassDay))
	}
	if !db.HasTable(new(ClassDayTime)) {
		db.CreateTable(new(ClassDayTime))
	}

	if !db.HasTable(new(Course)) {
		db.CreateTable(new(Course))
	}
	if !db.HasTable(new(Room)) {
		db.CreateTable(new(Room))
	}
	if !db.HasTable(new(Teacher)) {
		db.CreateTable(new(Teacher))
	}

	if !db.HasTable(new(Day)) {
		db.CreateTable(new(Day))
	}
	if !db.HasTable(new(DayTime)) {
		db.CreateTable(new(DayTime))
	}
}
