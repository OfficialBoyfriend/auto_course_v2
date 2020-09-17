package modules

import (
	"fmt"
	"testing"
)

func TestRoom_ResetTimeTable(t *testing.T) {

	//tmpInit()
	//defer tmpClose()

	rooms, err := NewRoom().List()
	if err != nil {
		t.Fatal(err)
	}

	for _, r := range rooms {
		for x := 0; x < 5; x++ {
			for y := 0; y < 10; y++ {
				r.TimeTable[x][y] = ""
			}
		}
		if err := r.Put(); err != nil {
			t.Fatal(err)
		}
	}
}

func TestRoom_Put(t *testing.T) {

	tmpInit()
	defer tmpClose()

	for i := 0; i < 56; i++ {
		room := Room{
			Name: fmt.Sprintf("[%d号]教室", i+1),
			TimeTable: [5][10]string{
				{"", "", "", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", "", "", ""},
				{"", "", "", "", "", "", "", "", "", ""},
			},
		}

		if err := room.Put(); err != nil {
			t.Fatal(err)
		}
		t.Log(room.Id)
	}
}

func TestRoom_Delete(t *testing.T) {
	room := Room{
		Id: "bthp67ueivh4555iv380",
	}

	tmpInit()
	defer tmpClose()

	if err := room.Delete(); err != nil {
		t.Fatal(err)
	}
}

func TestRoom_First(t *testing.T) {

	room := Room{
		Id: "bthp67ueivh4555iv380",
	}

	tmpInit()
	defer tmpClose()

	if err := room.First(); err != nil {
		t.Fatal(err)
	}
	t.Log(room)
}

func TestRoom_List(t *testing.T) {

	tmpInit()
	defer tmpClose()

	rooms, err := NewRoom().List()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("教室数量:", len(rooms))
	for _, r := range rooms {
		t.Logf("%v %v %v", r.Id, r.Name, r.TimeTable)
	}

	tmp := make([]string, 0)
	for _, r := range rooms {
		tmp = append(tmp, r.Id)
	}
	t.Logf("%#v", tmp)
}
