package modules

import "testing"

func TestClassCourse_DeleteAll(t *testing.T) {

	//tmpInit()
	//defer tmpClose()

	classCourse := new(ClassCourse)
	classCourses, err := classCourse.List()
	if err != nil {
		t.Fatal(err)
	}
	for _, cc := range classCourses {
		t.Log(cc.Delete())
	}
}

func TestClassCourse_First(t *testing.T) {

	tmpInit()
	defer tmpClose()

	classCourse := &ClassCourse{Id: "bthnsh6eivh1qt1p44kg"}
	if err := classCourse.First(); err != nil {
		t.Fatal(err)
	}
	t.Log(classCourse)
}

func TestClassCourse_List(t *testing.T) {
	tmpInit()
	defer tmpClose()

	classCourse := new(ClassCourse)
	classCourses, err := classCourse.List()
	if err != nil {
		t.Fatal(err)
	}

	for _, cc := range classCourses {
		t.Log(cc)
	}
}
