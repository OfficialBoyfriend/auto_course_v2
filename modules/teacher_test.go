package modules

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"testing"
)

func TestTeacher_Create(t *testing.T) {

	teacherName := []string{"王鹏", "刘旭", "王小刚", "吴杰", "罗丽梅", "周琼莉", "王小红"}

	for i := uint(0); i < 7; i++ {

		defaultDayTime := make([]DayTime, 10)
		for x := range defaultDayTime {
			defaultDayTime[x].Number = uint(x) + 1
		}

		days := make([]Day, 5)
		for y := range days {
			days[y].Number = uint(y) + 1
			days[y].Times = make([]DayTime, 10)
			copy(days[y].Times, defaultDayTime)
		}

		room := Teacher{
			Name: teacherName[i],
			Days: days,
		}
		if err := room.Create(); err != nil {
			t.Fatal(err)
		}
	}
}

func TestTeacher_First(t *testing.T) {
	teacher := Teacher{
		Model: Model{
			ID: 1,
		},
	}
	if err := teacher.First(); err != nil {
		t.Fatal(err)
	}
	resultJson, _ := json.Marshal(teacher)
	t.Log(string(resultJson))

	var resultTable [][]string
	for i := 0; i < 10; i++ {
		resultTable = append(resultTable, make([]string, 5))
	}
	for x := 0; x < 10; x++ {
		for i := 0; i < 5; i++ {
			text := "空 闲"
			if teacher.Days[i].Times[x].ClassId != 0 {
				text = fmt.Sprintf("%d", teacher.Days[i].Times[x].ClassId)
			}
			resultTable[x][i] = text
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"周 一", "周 二", "周 三", "周 四", "周 五"})
	for _, v := range resultTable {
		table.Append(v)
	}
	table.Render() // Send output
}
