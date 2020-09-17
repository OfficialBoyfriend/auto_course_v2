package modules

import (
	"testing"
)

func TestCourse_Put(t *testing.T) {

	roomIds := []string{"bthpe66eivh0se14cqv0", "bthpe66eivh0se14cqvg", "bthpe66eivh0se14cr00"}
	roomIdsDefault := []string{"bthpe66eivh0se14cr0g", "bthpe66eivh0se14cr10", "bthpe66eivh0se14cr1g", "bthpe66eivh0se14cr20", "bthpe66eivh0se14cr2g", "bthpe66eivh0se14cr30", "bthpe66eivh0se14cr3g", "bthpe66eivh0se14cr40", "bthpe66eivh0se14cr4g", "bthpe66eivh0se14cr50", "bthpe66eivh0se14cr5g", "bthpe66eivh0se14cr60", "bthpe66eivh0se14cr6g", "bthpe66eivh0se14cr70", "bthpe66eivh0se14cr7g", "bthpe66eivh0se14cr80", "bthpe66eivh0se14cr8g", "bthpe66eivh0se14cr90", "bthpe66eivh0se14cr9g", "bthpe66eivh0se14cra0", "bthpe66eivh0se14crag", "bthpe66eivh0se14crb0", "bthpe66eivh0se14crbg", "bthpe66eivh0se14crc0", "bthpe66eivh0se14crcg", "bthpe66eivh0se14crd0", "bthpe66eivh0se14crdg", "bthpe66eivh0se14cre0", "bthpe66eivh0se14creg", "bthpe66eivh0se14crf0", "bthpe66eivh0se14crfg", "bthpe66eivh0se14crg0", "bthpe66eivh0se14crgg", "bthpe66eivh0se14crh0", "bthpe66eivh0se14crhg", "bthpe66eivh0se14cri0", "bthpe66eivh0se14crig", "bthpe66eivh0se14crj0", "bthpe66eivh0se14crjg", "bthpe66eivh0se14crk0", "bthpe66eivh0se14crkg", "bthpe66eivh0se14crl0", "bthpe66eivh0se14crlg", "bthpe66eivh0se14crm0", "bthpe66eivh0se14crmg", "bthpe66eivh0se14crn0", "bthpe66eivh0se14crng", "bthpe66eivh0se14cro0", "bthpe66eivh0se14crog", "bthpe66eivh0se14crp0", "bthpe66eivh0se14crpg", "bthpe66eivh0se14crq0", "bthpe66eivh0se14crqg"}
	courseNames := []string{"机械制图与CAD", "先进生产技术", "工业机器人虚拟仿真", "党史国史", "计算机编程技术", "电机与电气控制技术", "公差配合与测量技术"}
	teacherIds := []string{"bthpe46eivh3ns7gaa7g", "bthpe46eivh3ns7gaa80", "bthpe46eivh3ns7gaa8g", "bthpe46eivh3ns7gaa90", "bthpe46eivh3ns7gaa9g", "bthpe46eivh3ns7gaaa0", "bthpe46eivh3ns7gaaag"}

	tmpInit()
	defer tmpClose()

	for i := range courseNames {

		course := Course{
			Name:       courseNames[i],
			TeacherIds: []string{teacherIds[i]},
		}

		// 设置可选教室
		switch courseNames[i] {
		case "工业机器人虚拟仿真":
			fallthrough
		case "计算机编程技术":
			course.RoomIds = roomIds
		default:
			course.RoomIds = roomIdsDefault
		}

		// 创建记录
		if err := course.Put(); err != nil {
			t.Fatal(err)
		}

		t.Log(course.Id)
	}
}

func TestCourse_Delete(t *testing.T) {
	course := Course{
		Id: "bthpfmmeivh1si4g42vg",
	}

	tmpInit()
	defer tmpClose()

	if err := course.Delete(); err != nil {
		t.Fatal(err)
	}
}

func TestCourse_First(t *testing.T) {
	course := Course{
		Id: "bthpfmmeivh1si4g42vg",
	}

	tmpInit()
	defer tmpClose()

	if err := course.First(); err != nil {
		t.Fatal(err)
	}
	t.Log(course)
}

func TestCourse_List(t *testing.T) {

	tmpInit()
	defer tmpClose()

	courses, err := NewCourse().List()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("课程数量:", len(courses))
	for _, v := range courses {
		t.Logf("%s %s", v.Id, v.Name)
	}

	tmp := make([]string, 0)
	for _, r := range courses {
		tmp = append(tmp, r.Id)
	}
	t.Logf("%#v", tmp)
}
