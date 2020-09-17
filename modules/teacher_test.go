package modules

import (
	"testing"
)

func TestTeacher_ResetTimeTable(test *testing.T) {

	//tmpInit()
	//defer tmpClose()

	teachers, err := NewTeacher().List()
	if err != nil {
		test.Fatal(err)
	}

	for _, t := range teachers {
		for x := 0; x < 5; x++ {
			for y := 0; y < 10; y++ {
				t.TimeTable[x][y] = ""
			}
		}
		if err := t.Put(); err != nil {
			test.Fatal(err)
		}
	}
}

func TestTeacher_Put(t *testing.T) {

	teacherName := []string{"王鹏", "刘旭", "王小刚", "吴杰", "罗丽梅", "周琼莉", "王小红"}

	tmpInit()
	defer tmpClose()

	for _, name := range teacherName {
		teacher := Teacher{
			Name: name,
			TimeTable: [5][10]string{
				{"", "", "", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", "", "", ""},
			},
		}

		if err := teacher.Put(); err != nil {
			t.Fatal(err)
		}
		t.Log(teacher.Id)
	}
}

func TestTeacher_Delete(t *testing.T) {
	teacher := Teacher{
		Id: "bthp3ueeivh4tj6slo1g",
	}

	tmpInit()
	defer tmpClose()

	t.Log(teacher.Delete())
}

func TestTeacher_First(t *testing.T) {

	teacher := Teacher{
		Id: "bthp3ueeivh4tj6slo1g",
	}

	tmpInit()
	defer tmpClose()

	if err := teacher.First(); err != nil {
		t.Fatal(err)
	}
	t.Log(teacher)
}

func TestTeacher_List(t *testing.T) {

	tmpInit()
	defer tmpClose()

	teachers, err := NewTeacher().List()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("教师数量:", len(teachers))
	for _, v := range teachers {
		t.Logf("%v %v %v", v.Id, v.Name, v.TimeTable)
	}

	tmp := make([]string, 0)
	for _, r := range teachers {
		tmp = append(tmp, r.Id)
	}
	t.Logf("%#v", tmp)
}
