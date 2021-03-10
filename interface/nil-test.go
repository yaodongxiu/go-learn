package main

import (
	"fmt"
)

type BaseQuestion struct {
	QuestionId      int
	QuestionContent string
}

type ChoiceQuestion struct {
	BaseQuestion
	Options []string
}

type BlankQuestion struct {
	BaseQuestion
	Blank string
}

func fetchFromChoiceTable(id int) (interface{}, bool) {
	return &ChoiceQuestion{}, true
}

func fetchFromBlankTable(id int) (interface{}, bool) {
	return &BlankQuestion{}, false
}

func fetchQuestion(id int) (interface{}, bool) {
	data1, ok1 := fetchFromChoiceTable(id) // 根据ID到选择题表中找题目，返回(ChoiceQuestion)
	data2, ok2 := fetchFromBlankTable(id)  // 根据ID到填空题表中找题目，返回(BlankQuestion)

	if ok1 {
		return data1, ok1
	}

	if ok2 {
		return data2, ok2
	}

	return nil, false
}

func printQuestion() {
	if data, ok := fetchQuestion(1001); ok {
		switch v := data.(type) {
		case *ChoiceQuestion:
			v.QuestionContent = "选择题" + v.QuestionContent
			fmt.Println(v)
		case *BlankQuestion:
			v.QuestionContent = "wendati" + v.QuestionContent
			fmt.Println(v)
		case nil:
			fmt.Println(v)
		}
		fmt.Println(data)
	}
}

func fetchFromChoiceTableNil(id int) (data *ChoiceQuestion) {
	if id == 1001 {
		return &ChoiceQuestion{
			BaseQuestion: BaseQuestion{
				QuestionId:      1001,
				QuestionContent: "HELLO",
			},
			Options: []string{"A", "B"},
		}
	}
	return nil
}

func fetchQuestionNil(id int) interface{} {
	data1 := fetchFromChoiceTableNil(id) // 根据ID到选择题表中找题目
	//if data1 == nil {
	//	return nil
	//}
	return data1
}

func sendData(data interface{}) {
	fmt.Println("发送数据 ...", data)
}

func main() {
	data := fetchQuestionNil(1002)

	if data != nil {
		sendData(data)
	}
}
