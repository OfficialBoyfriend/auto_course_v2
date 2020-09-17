package modules

import (
	"auto_course/utils"
	"fmt"
	"github.com/modood/table"
	"testing"
)

func TestClass_DeleteAll(t *testing.T) {

	isDeleteAll := true

	tmpInit()
	defer tmpClose()

	class := new(Class)
	classes, err := class.List()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range classes {
		if c.Name != "" && !isDeleteAll {
			continue
		}
		t.Log(c.Delete())
	}

	if isDeleteAll {
		TestClassCourse_DeleteAll(t)
		TestRoom_ResetTimeTable(t)
		TestTeacher_ResetTimeTable(t)
	}
}

func TestClass_Data(t *testing.T) {
	courseIds := []string{"bthpfmmeivh1si4g42vg", "bthpfmmeivh1si4g4300", "bthpfmmeivh1si4g430g", "bthpfmmeivh1si4g4310", "bthpfmmeivh1si4g431g", "bthpfmmeivh1si4g4320", "bthpfmmeivh1si4g432g"}
	teacherIds := []string{"bthpe46eivh3ns7gaa7g", "bthpe46eivh3ns7gaa80", "bthpe46eivh3ns7gaa8g", "bthpe46eivh3ns7gaa90", "bthpe46eivh3ns7gaa9g", "bthpe46eivh3ns7gaaa0", "bthpe46eivh3ns7gaaag"}

	tmpInit()
	defer tmpClose()

	for _, c := range courseIds {
		course := &Course{Id: c}
		if err := course.First(); err != nil {
			t.Fatal(err)
		}
		t.Log(course)
	}

	for _, teacher := range teacherIds {
		teacher := &Teacher{Id: teacher}
		if err := teacher.First(); err != nil {
			t.Fatal(err)
		}
		t.Log(teacher)
	}
}

func TestClass_Set(t *testing.T) {
	classNames := []string{"工业机器人1801", "工业机器人1802"}
	courseIds := []string{"bthpfmmeivh1si4g42vg", "bthpfmmeivh1si4g4300", "bthpfmmeivh1si4g430g", "bthpfmmeivh1si4g4310", "bthpfmmeivh1si4g431g", "bthpfmmeivh1si4g4320", "bthpfmmeivh1si4g432g"}
	teacherIds := []string{"bthpe46eivh3ns7gaa7g", "bthpe46eivh3ns7gaa80", "bthpe46eivh3ns7gaa8g", "bthpe46eivh3ns7gaa90", "bthpe46eivh3ns7gaa9g", "bthpe46eivh3ns7gaaa0", "bthpe46eivh3ns7gaaag"}
	courseTime := []uint{4, 2, 2, 2, 4, 2, 2}

	tmpInit()
	defer tmpClose()

	for i := range classNames {

		courses := make([]*ClassCourse, 0)
		for x := range courseIds {
			courses = append(courses, &ClassCourse{
				CourseId:  courseIds[x],
				TeacherId: teacherIds[x],
				Number:    courseTime[x] * 2, // TODO
			})
		}

		class := Class{
			//Id: "bthmg1ueivh3kn1qa6b0",
			Name:    classNames[i],
			Courses: courses,
			IsOk:    false,
			TimeTable: [5][10]string{
				{"", "", "", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", "", "", ""},
			},
		}
		// 创建记录
		if err := class.Put(); err != nil {
			t.Fatal(err)
		}

		t.Log(class.Id)
		//break
	}
}

func TestClass_Delete(t *testing.T) {

	tmpInit()
	defer tmpClose()

	class := Class{Id: "bthmg1ueivh3kn1qa6b0"}
	t.Log(class.Delete())
}

func TestClass_First(t *testing.T) {
	class := Class{
		Id: "bthmg1ueivh3kn1qa6b0",
	}

	tmpInit()
	defer tmpClose()

	if err := class.First(); err != nil {
		t.Fatal(err)
	}
	t.Log(class)
}

func TestClass_List(t *testing.T) {

	tmpInit()
	defer tmpClose()

	classes, err := NewClass().List()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("班级数: %d", len(classes))

	for _, v := range classes {

		if v.Name == "" {
			t.Log(v.Delete())
			continue
		}
		t.Logf("[%s]", v.Name)

		/****** 打印表格 start *******/

		type TableData struct {
			Todo string
			A    string
			B    string
			C    string
			D    string
			E    string
		}
		var (
			resultTable []TableData
		)
		for i := 0; i < 10; i++ {
			tmp := make([]string, 6)
			tmp[0] = fmt.Sprintf("第%d节", i+1)
			for ii := 0; ii < 5; ii++ {
				// 检查空课程
				if v.TimeTable[ii][i] == "" {
					tmp[ii+1] = "暂无课程"
					continue
				}
				classCourse := &ClassCourse{Id: v.TimeTable[ii][i]}
				if err := classCourse.First(); err != nil {
					t.Fatal(err)
				}
				teacher := &Teacher{Id: classCourse.TeacherId}
				if err := teacher.First(); err != nil {
					t.Fatal(err)
				}
				room := &Room{Id: classCourse.RoomId}
				if err := room.First(); err != nil {
					t.Fatal(err)
				}
				course := &Course{Id: classCourse.CourseId}
				if err := course.First(); err != nil {
					t.Fatal(err)
				}
				tmp[ii+1] = fmt.Sprintf("%v-%v-%v", course.Name, room.Name, teacher.Name)
			}
			resultTable = append(resultTable, TableData{
				Todo: tmp[0],
				A:    tmp[1],
				B:    tmp[2],
				C:    tmp[3],
				D:    tmp[4],
				E:    tmp[5],
			})
		}

		table.Output(resultTable)
		/****** 打印表格 end *******/
	}
}

func TestAutoCourse(t *testing.T) {

	tmpInit()
	defer tmpClose()

	// 获取未排课班级
	classes, err := getNotArrangementClass()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(classes)
	if len(classes) < 1 {
		t.Log("无待排课班级")
		return
	}

	// 获取班级课程详情(包含教师与课程详细数据)
	// TODO: Key: classId	Value: Slice[ Slice[ 0: teacher, 1: course 2: number(课时) ] ... ]
	classCourseMap := make(map[string][][3]interface{})
	for _, c := range classes {
		courseDetails := make([][3]interface{}, 0)
		for _, cc := range c.Courses {
			teacher, course, err := getClassCourseDetails(cc)
			if err != nil {
				t.Fatal(err)
			}
			courseDetails = append(courseDetails, [3]interface{}{teacher, course, cc.Number})
		}
		classCourseMap[c.Id] = courseDetails
	}

	// TODO 班级
	for _, c := range classes {
		// TODO 课程
		for _, cc := range classCourseMap[c.Id] {
			for i := uint(0); i < cc[2].(uint); i++ {

				// 获取教师与班级共同空闲时间
				nilTime, err := getNilTime(c, cc[0].(*Teacher))
				if err != nil {
					t.Fatal(err)
				}
				// 查询课程可用教室
				rooms := make([]*Room, 0)
				for _, id := range cc[1].(*Course).RoomIds {
					room := &Room{Id: id}
					if err := room.First(); err != nil {
						t.Fatal(err)
					}
					rooms = append(rooms, room)
				}
				// 具有共同空闲时间教室
				okRooms, err := getNilRoom(nilTime, rooms)
				if err != nil {
					t.Fatal(err)
				}
				// 打乱可用教室顺序
				utils.Shuffle(okRooms)

				t.Logf("[%s] [%s] [%s] [%d] %v", c.Name, cc[0].(*Teacher).Name, cc[1].(*Course).Name, len(okRooms), nilTime)

				/****** 更新数据 start *******/

				// 新建读写事务
				db := NewDBBatch()

				// 添加课程表
				classCourse := ClassCourse{
					CourseId:  cc[1].(*Course).Id,
					TeacherId: cc[0].(*Teacher).Id,
					RoomId:    okRooms[0][0].(string),
					RoomTime:  [2]int{okRooms[0][1].(int), okRooms[0][2].(int)},
				}
				if err := classCourse.Put(db); err != nil {
					t.Fatal(err)
				}

				// 更新班级时间表
				c.TimeTable[okRooms[0][1].(int)][okRooms[0][2].(int)] = classCourse.Id
				c.IsOk = true
				if err := c.Put(db); err != nil {
					t.Fatal(err)
				}

				// 更新教师数据
				// 更新时间表
				cc[0].(*Teacher).TimeTable[okRooms[0][1].(int)][okRooms[0][2].(int)] = c.Id
				if err := cc[0].(*Teacher).Put(db); err != nil {
					t.Fatal(err)
				}

				// 更新教室数据
				// 更新时间表
				room := Room{Id: okRooms[0][0].(string)}
				if err := room.First(); err != nil {
					t.Fatal(err)
				}
				room.TimeTable[okRooms[0][1].(int)][okRooms[0][2].(int)] = c.Id
				if err := room.Put(db); err != nil {
					t.Fatal(err)
				}

				/****** 更新数据 end ******/
			}
		}
	}
}

// getNotCourseClass 获取未安排课程的班级
func getNotArrangementClass() ([]*Class, error) {

	classes, err := NewClass().List()
	if err != nil {
		return nil, err
	}

	var result []*Class
	for _, v := range classes {
		if v.IsOk {
			continue
		}
		result = append(result, v)
	}

	return result, nil
}

// getClassCourseDetails 根据索引获取班级课程与教师详情
func getClassCourseDetails(c *ClassCourse) (*Teacher, *Course, error) {

	// 查询课程信息
	course := &Course{Id: c.CourseId}
	if err := course.First(); err != nil {
		return nil, nil, err
	}

	// 查询任课教师信息
	teacher := &Teacher{Id: c.TeacherId}
	if err := teacher.First(); err != nil {
		return nil, nil, err
	}

	return teacher, course, nil
}

// getNilTime 获取班级与教师的共同空闲时间
func getNilTime(class *Class, teacher *Teacher) ([][2]int, error) {
	// 班级空闲时间
	var classNilTime [][2]int
	for i, d := range class.TimeTable {
		for ii, t := range d {
			if t != "" {
				continue
			}
			classNilTime = append(classNilTime, [2]int{i, ii})
		}
	}
	// 教师空闲时间
	var teacherNilTime [][2]int
	for i, d := range teacher.TimeTable {
		for ii, t := range d {
			if t != "" {
				continue
			}
			teacherNilTime = append(teacherNilTime, [2]int{i, ii})
		}
	}
	// 共同空闲时间
	var commonNilTime [][2]int
	for _, c := range classNilTime {
		for _, t := range teacherNilTime {
			if c != t {
				continue
			}
			commonNilTime = append(commonNilTime, c)
		}
	}
	return commonNilTime, nil
}

// getNilRoom 选择空闲教室
// 从多个可选教室中获取指定时间内空闲的教室
// TODO: return value Slice[ Array(3)[ id x y ] ... ]
func getNilRoom(nilTime [][2]int, rooms []*Room) ([][3]interface{}, error) {
	var nilRooms [][3]interface{}
	for _, n := range nilTime {
		for _, r := range rooms {
			if r.TimeTable[n[0]][n[1]] != "" {
				continue
			}
			nilRooms = append(nilRooms, [3]interface{}{r.Id, n[0], n[1]})
		}
	}
	return nilRooms, nil
}
