package main

var (
	// 将数据库中已有的所有班级取出来
	classList = []string{"1801", "1802", "1803", "1804"}
	// 分别为每个有的班级分别添加上每一周内的20个时间片段
	classSlot = map[string]interface{}{
		"1801": [5][5]interface{}{
			{}, {}, {}, {}, {},
		},
		"1802": [5][5]interface{}{
			{}, {}, {}, {}, {},
		},
		"1803": [5][5]interface{}{
			{}, {}, {}, {}, {},
		},
		"1804": [5][5]interface{}{
			{}, {}, {}, {}, {},
		},
	}
	// 为每个班级随机加入其要上的课程
	classCourse = map[string]interface{}{
		"1801": []map[string]interface{}{
			{"name": "C语言", "len": 8, "max": 4, "room": []int{1, 2, 3}},
			{"name": "ABB编程", "len": 4, "max": 4, "room": []int{1, 2, 3}},
			{"name": "汽车拆装", "len": 4, "max": 4, "room": []int{1, 2, 3}},
		},
	}
)

func main() {

}
