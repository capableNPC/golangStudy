package main

import "fmt"

type student struct {
	id, name   string
	age, score int
}

var selectedStuPtr *student
var studentsList []*student

func newStudent(id, name string, age, score int) *student {
	return &student{
		id:    id,
		name:  name,
		age:   age,
		score: score,
	}
}
func (s *student) updateInfo(id, name string, age, score int) {
	s.id = id
	s.name = name
	s.age = age
	s.score = score
	fmt.Println("学生的信息更新完毕")
}

func initList() {
	count := 10
	strID := "s2020"
	var letter rune = 'a'
	age := 10
	score := 60
	for i := 0; i < 10; i++ {
		studentsList = append(studentsList, &student{
			id:    strID + fmt.Sprintf("%d", count),
			name:  string(letter),
			age:   age,
			score: score * (11 - i) / 10,
		})
		letter++
		age++
		count++
	}
}

func showStudentsList() {
	fmt.Printf("   id    姓名    年龄    分数\n")
	for _, v := range studentsList {
		fmt.Printf("%s    %s      %d      %d\n", v.id, v.name, v.age, v.score)
	}
}

func addStudent() {
	fmt.Println("请输入学生的信息")
	var id, name string
	var age, score int
	fmt.Print("ID:")
	fmt.Scanln(&id)
	fmt.Print("姓名:")
	fmt.Scanln(&name)
	fmt.Print("年龄:")
	fmt.Scanln(&age)
	fmt.Print("分数:")
	fmt.Scanln(&score)
	studentsList = append(studentsList, &student{
		id:    id,
		name:  name,
		age:   age,
		score: score,
	})
}

func editStudentInfo() {
	fmt.Println("请输入学生的信息")
	var id, name string
	var age, score int
	fmt.Print("ID:")
	fmt.Scanln(&id)
	fmt.Print("姓名:")
	fmt.Scanln(&name)
	fmt.Print("年龄:")
	fmt.Scanln(&age)
	fmt.Print("分数:")
	fmt.Scanln(&score)
	selectedStuPtr.updateInfo(id, name, age, score)
}

func searchStudent() {
	var ID string
	fmt.Println("请输入学生的id")
	fmt.Print("ID:")
	fmt.Scanln(&ID)

	selectedStuPtr = nil
	for _, v := range studentsList {
		if ID == v.id {
			selectedStuPtr = v
		}
	}
	if selectedStuPtr == nil {
		fmt.Println("id：", ID, "不存在。")
	} else {
		fmt.Printf("   id    姓名    年龄    分数\n")
		fmt.Printf("%s    %s      %d      %d\n", selectedStuPtr.id,
			selectedStuPtr.name, selectedStuPtr.age, selectedStuPtr.score)
		fmt.Println("学生 ", selectedStuPtr.name, " 已选中")
	}
}

func isSelect() bool {
	if selectedStuPtr == nil {
		return false
	}
	return true
}
func deleteStudent() {
	cursor := -1
	for i, v := range studentsList {
		if selectedStuPtr.id == v.id {
			cursor = i
			break
		}
	}
	name := selectedStuPtr.name
	fmt.Println("学生 ", name, " 正在删除")
	if cursor == (len(studentsList) - 1) {
		studentsList = append(studentsList[:cursor])
	} else {
		studentsList = append(studentsList[:cursor], studentsList[cursor+1:]...)
	}
	fmt.Println("学生 ", name, " 已删除")
}
func menu() bool {
	fmt.Println("1.添加学生")
	fmt.Println("2.编辑学生")
	fmt.Println("3.删除学生")
	fmt.Println("4.选中学生")
	fmt.Println("5.显示学生名单")
	fmt.Println("6.退出")
	code := -1
	for code == -1 {
		fmt.Print("选择功能:")
		fmt.Scanln(&code)
		if code != 1 && code != 2 && code != 3 && code != 4 && code != 5 && code != 6 {
			fmt.Println("无效的指令:", code)
			code = -1
		}
	}
	switch code {
	case 1:
		addStudent()
	case 2:
		if isSelect() {
			editStudentInfo()
		} else {
			fmt.Println("未选择学生请先选择一名学生")
		}
	case 3:
		if isSelect() {
			deleteStudent()
		} else {
			fmt.Println("未选择学生请先选择一名学生")
		}
	case 4:
		searchStudent()
	case 5:
		showStudentsList()
	case 6:
		return false
	default:
		fmt.Println("爬")
	}
	return true
}

func main() {
	initList()
	fmt.Println("-------------------------------------------------")
	keepLoop := true
	for keepLoop {
		keepLoop = menu()
	}
}
