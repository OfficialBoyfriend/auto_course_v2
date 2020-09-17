package modules

// 天数
// 时间表中使用
type Day struct {
	Model
	TeacherId uint      `gorm:"not null" json:"-"` // 教师标识
	RoomId    uint      `gorm:"not null" json:"-"` // 教室标识
	Number    uint      `gorm:"not null"`          // 记录天数
	Times     []DayTime `gorm:"not null;"`         // 时间与班级绑定
}

type DayTime struct {
	Model
	DayId   uint `gorm:"not null" json:"-"` // 天数标识
	ClassId uint `json:"-"`                 // 班级标识
	Number  uint `gorm:"not null"`          // 记录天数
}
