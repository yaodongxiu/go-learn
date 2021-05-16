package do

import "fmt"

const (
	_ = iota
	ChoiceQuestionType
	BlankQuestionType
)

type IQuestion interface {
	GetQuestionType() int
	GetQuestionContent() string
	AddQuestionContentPrefix(prefix string)
}

type BaseQuestion struct {
	QuestionId      int
	QuestionContent string
	QuestionType    int
}

func (bq *BaseQuestion) GetQuestionType() int {
	return bq.QuestionType
}

func (bq *BaseQuestion) GetQuestionContent() string {
	return bq.QuestionContent
}

func (bq *BaseQuestion) AddQuestionContentPrefix(prefix string) {
	bq.QuestionContent = prefix + bq.QuestionContent
}

type IChoiceQuestion interface {
	IQuestion
	GetOptionsLen() int
}

type ChoiceQuestion struct {
	BaseQuestion
	Options []string
}

func (cq *ChoiceQuestion) GetOptionsLen() int {
	return len(cq.Options)
}

type BlankQuestion struct {
	BaseQuestion
	Blank []string
}

func fetchFromBlankTable(id int) (BlankQuestion, bool) {
	if id == 1002 {
		return BlankQuestion{
			BaseQuestion: BaseQuestion{
				QuestionId:      id,
				QuestionContent: "BLANK",
				QuestionType:    BlankQuestionType,
			},
			Blank: []string{"ET", "AI"},
		}, true
	}
	return BlankQuestion{}, false
}

func fetchFromChoiceTable(id int) (ChoiceQuestion, bool) {
	if id == 1001 {
		return ChoiceQuestion{
			BaseQuestion: BaseQuestion{
				QuestionId:      id,
				QuestionContent: "CHOICE",
				QuestionType:    ChoiceQuestionType,
			},
			Options: []string{"A", "B"},
		}, true
	}
	return ChoiceQuestion{}, false
}

func fetchQuestion(id int) (IQuestion, bool) {
	data1, ok1 := fetchFromChoiceTable(id) // 根据ID到选择题表中找题目
	data2, ok2 := fetchFromBlankTable(id)  // 根据ID到选择题表中找题目

	if ok1 {
		return &data1, ok1
	}

	if ok2 {
		return &data2, ok2
	}

	return nil, false
}

func showOptionsLen(data IQuestion) {
	if choice, ok := data.(IChoiceQuestion); ok {
		fmt.Println("Choice has :", choice.GetOptionsLen())
	}
}

func printQuestion() {
	if data, ok := fetchQuestion(1001); ok {
		var questionPrefix string

		switch data.GetQuestionType() {
		case ChoiceQuestionType:
			questionPrefix = "选择题"
		case BlankQuestionType:
			questionPrefix = "填空题"
		}

		showOptionsLen(data)
		data.AddQuestionContentPrefix(questionPrefix)
		fmt.Println("data - ", data)
	}
}
