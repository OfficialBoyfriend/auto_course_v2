package modules

import (
	"auto_course/utils"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"testing"
)

func TestClass_Create(t *testing.T) {
	classNames := []string{"工业机器人1801", "工业机器人1802"}
	courseId := []uint{1, 2, 3, 4, 5, 6, 7}
	teacherIds := []uint{1, 2, 3, 4, 5, 6, 7}

	for i := range classNames {

		courses := make([]ClassCourse, 0)
		for x := range courseId {
			courses = append(courses, ClassCourse{
				CourseId:  courseId[x],
				TeacherId: teacherIds[x],
			})
		}

		defaultDayTime := make([]ClassDayTime, 10)
		for x := range defaultDayTime {
			defaultDayTime[x].Number = uint(x) + 1
		}

		days := make([]ClassDay, 5)
		for y := range days {
			days[y].Number = uint(y) + 1
			days[y].Times = make([]ClassDayTime, 10)
			copy(days[y].Times, defaultDayTime)
		}

		class := Class{
			Name:    classNames[i],
			Courses: courses,
			Days:    days,
		}
		// 创建记录
		if err := class.Create(); err != nil {
			t.Fatal(err)
		}
	}
}

func TestClass_First(t *testing.T) {
	class := Class{
		Model: Model{ID: 1},
	}
	if err := class.First(); err != nil {
		t.Fatal(err)
	}
	resultJson, _ := json.Marshal(class)
	t.Log(string(resultJson))

	var resultTable [][]string
	for i := 0; i < 10; i++ {
		resultTable = append(resultTable, make([]string, 5))
	}
	for x := 0; x < 10; x++ {
		for i := 0; i < 5; i++ {
			text := "空 闲"
			if class.Days[i].Times[x].ClassCourseId != 0 {
				text = fmt.Sprintf("%d", class.Days[i].Times[x].ClassCourseId)
			}
			resultTable[x][i] = text
		}
	}

	t.Log(class.Name)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"周 一", "周 二", "周 三", "周 四", "周 五"})
	for _, v := range resultTable {
		table.Append(v)
	}
	table.Render() // Send output
}

func TestAutoCourse(t *testing.T) {

	// 获取一个未排课班级
	class, err := getFirstNotCourseClass()
	if err != nil {
		t.Fatal(err)
	}
	resultJson, _ := json.Marshal(class)
	t.Logf("[%s] 待安排课程: %v", class.Name, string(resultJson))

	// 获取该班课程详细信息
	courses, err := getClassCourseDetails(class.Courses)
	if err != nil {
		t.Fatal(err)
	}
	resultJson, _ = json.Marshal(courses)
	//t.Log(string(resultJson))

	// 教师、班级空闲时间
	classNilTime, teacherNilTimeMap, err := getNilTime(class, courses)
	if err != nil {
		t.Fatal(err)
	}

	// 教师与班级共同空闲时间
	commonNilTimeMap := make(map[uint][][2]uint)
	for _, v := range courses {
		commonNilTime, err := getCommonNilTime(classNilTime, teacherNilTimeMap[v.ID])
		if err != nil {
			t.Log(err)
		}
		// 打乱时间排序
		utils.Shuffle(commonNilTime)
		commonNilTimeMap[v.ID] = commonNilTime
	}



	// 开始排课
	for i, v := range courses {

		if len(commonNilTimeMap[v.ID]) < 1 {
			t.Logf("[%s] 无共同空余时间", courses[i].Name)
			// 跳过该课程
			continue
		}

		// 选择空闲教室
		roomId := uint(0)
		var roomTime [2]uint
	Exit:
		for r := range courses[i].Rooms {
			for rr := range courses[i].Rooms[r].Days {
				for rrr := range courses[i].Rooms[r].Days[rr].Times {
					if courses[i].Rooms[r].Days[rr].Times[rrr].ClassId != 0 {
						continue
					}
					roomId = courses[i].Rooms[r].ID
					roomTime = [2]uint{courses[i].Rooms[r].Days[rr].Number, courses[i].Rooms[r].Days[rr].Times[rrr].Number}
					break Exit
				}
			}
		}

		if roomId == 0 {
			t.Logf("[%s] [%s] 安排教室失败", courses[i].Name, class.Name)
			break
		}

		// 更新班级时间安排
		class.Days[x].Times[y].ClassCourseId = courses[i].ID
		class.Courses = append(class.Courses, ClassCourse{
			Model:     Model{},
			ClassId:   class.ID,
			CourseId:  courses[i].ID,
			TeacherId: class.Courses[i].TeacherId,
			RoomId:    roomId,
			RoomTime:  roomTime,
			Priority:  0,
		})

		room := Room{
			Model: Model{ID: roomId},
		}
		if err := room.First(); err != nil {
			t.Logf("[%s] [%s] 查询教室失败", courses[i].Name, class.Name)
			break
		}
		room.Days[x].Times[y].ClassId = class.ID

		// 更新教师时间安排
		teacher := Teacher{
			Model: Model{ID: class.Courses[i].TeacherId},
		}
		if err := teacher.First(); err != nil {
			t.Fatal(err)
		}
		teacher.Days[x].Times[y].ClassId = class.ID

	}

}

// getFirstNotCourseClass 获取一个未排课班级
func getFirstNotCourseClass() (*Class, error) {

	class := &Class{IsOk: false}
	if err := class.First(); err != nil {
		return nil, fmt.Errorf("获取未排课班级失败: %v", err)
	}

	return class, nil
}

// getClassCourseDetails 根据索引获取班级课程详情
func getClassCourseDetails(c []ClassCourse) ([]*Course, error) {

	courses := make([]*Course, 0)
	for i := range c {
		id := c[i].CourseId
		course := &Course{Model: Model{ID: id}}
		if err := course.First(); err != nil {
			return nil, fmt.Errorf("获取课程详情失败: %v", err)
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func getNilTime(class *Class, courses []*Course) ([][2]uint, map[uint][][2]uint, error) {

	// 班级空闲时间
	classNilTime, err := getClassNilTime(class.ID)
	if err != nil {
		return nil, nil, err
	}

	// 教师空闲时间
	teacherNilTimeMap := make(map[uint][][2]uint)
	for i := range courses {
		teacherNilTime, err := getTeacherNilTime(courses[i].ID)
		if err != nil {
			return nil, nil, err
		}
		teacherNilTimeMap[courses[i].ID] = teacherNilTime
	}

	return classNilTime, teacherNilTimeMap, nil
}

// getClassNilTime 获取班级空闲时间
// 返回空闲时间的编号组合 [ day, time ]
func getClassNilTime(id uint) ([][2]uint, error) {

	class := Class{
		Model: Model{ID: id},
	}
	if err := class.First(); err != nil {
		return nil, fmt.Errorf("获取班级空闲时间失败，查询班级信息是发生错误: %v", err)
	}

	var classNilTime [][2]uint
	for x := range class.Days {
		for y := range class.Days[x].Times {
			if class.Days[x].Times[y].ClassCourseId != 0 {
				continue
			}
			classNilTime = append(classNilTime, [2]uint{class.Days[x].Number, class.Days[x].Times[y].Number})
		}
	}

	return classNilTime, nil
}

// getTeacherNilTime 获取教室空闲时间
// 返回空闲时间的编号组合 [ day, time ]
func getTeacherNilTime(id uint) ([][2]uint, error) {
	teacherNilTime := make([][2]uint, 0)
	// 查询课程任课教师信息
	teacher := Teacher{
		Model: Model{ID: id},
	}
	if err := teacher.First(); err != nil {
		return nil, fmt.Errorf("查询教室信息失败: %v", err)
	}
	// 汇总教师空闲时间
	for x := range teacher.Days {
		for y := range teacher.Days[x].Times {
			if teacher.Days[x].Times[y].ClassId != 0 {
				continue
			}
			teacherNilTime = append(teacherNilTime, [2]uint{teacher.Days[x].Number, teacher.Days[x].Times[y].Number})
		}
	}
	return teacherNilTime, nil
}

// getCommonNilTime 获取指定班级与指定教师的空闲时间中筛选共同空闲时间
// 返回空闲时间的编号组合 [ day, time ]
func getCommonNilTime(classNilTime [][2]uint, teacherNilTime [][2]uint) ([][2]uint, error) {
	commonNilTime := make([][2]uint, 0)
	for _, c := range classNilTime {
		for _, t := range teacherNilTime {
			if c[0] != t[0] || c[1] != t[1] {
				continue
			}
			commonNilTime = append(commonNilTime, [2]uint{c[0], c[1]})
			break
		}
	}
	return commonNilTime, nil
}

// 选择空闲教室
func getNilRoom(rooms []*Room, nilTime [][2]uint) (uint, [2]uint, error) {

	roomId := uint(0)
	roomTime := [2]uint{}

Exit:
		// 可选教室
		for _, r := range rooms {
			// 星期
			for _, d := range r.Days {
				// 具体时间
				for _, t := range d.Times {
					// 排除已占用教室
					if t.ClassId != 0 {
						continue
					}

					if d.Number != nilTime[0] || c[1] != t[1] {
						continue
					}

					roomId = t.ID
					roomTime = [2]uint{d.Number, t.Number}
					break Exit
				}
			}
		}

	if roomId == 0 {
		return 0, roomTime, fmt.Errorf("安排教室失败")
	}

	return roomId, roomTime, nil
}
