package modules

import (
	"encoding/json"
	"testing"
)

func TestCourse_Create(t *testing.T) {
	courseNames := []string{"机械制图与CAD", "先进生产技术", "工业机器人虚拟仿真", "党史国史", "计算机编程技术", "电机与电气控制技术", "公差配合与测量技术"}
	teacherIds := []uint{1, 2, 3, 4, 5, 6, 7}

	// 默认可选教室
	rooms := make([]Room, 53)
	for i := uint(0); i < 53; i++ {
		room := Room{Model: Model{ID: i + 1}}
		room.First()
		rooms[i] = room
	}

	for i := range courseNames {

		teacher := Teacher{Model: Model{ID: teacherIds[i]}}
		teacher.First()

		course := Course{
			Name: courseNames[i],
			Teachers: []Teacher{
				teacher,
			},
		}
		// 设置可选教室
		switch courseNames[i] {
		case "工业机器人虚拟仿真":
			fallthrough
		case "计算机编程技术":
			for i := uint(53); i < 56; i++ {
				room := Room{Model: Model{ID: i + 1}}
				room.First()
				course.Rooms = append(course.Rooms, room)
			}
		default:
			course.Rooms = rooms
		}
		// 创建记录
		if err := course.Create(); err != nil {
			t.Fatal(err)
		}
	}
}

func TestCourse_First(t *testing.T) {
	course := Course{
		Model: Model{ID: 1},
	}
	if err := course.First(); err != nil {
		t.Fatal(err)
	}
	resultJson, _ := json.Marshal(course)
	t.Log(string(resultJson))
}
