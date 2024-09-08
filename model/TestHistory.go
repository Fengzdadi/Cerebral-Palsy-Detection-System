// Test x

package model

import (
	"Cerebral-Palsy-Detection-System/Pkg/e"
	"Cerebral-Palsy-Detection-System/Serializer"
	"fmt"
	logging "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type TestHistory struct {
	//gorm.Model
	BelongToChildID uint      `gorm:"BelongToChildID" json:"belongToChildID" form:"belongToChildID"`
	TestTime        time.Time `gorm:"TestDate" json:"testTime" form:"testTime"`
	RawPath         string    `gorm:"RawPath" json:"rawPath" form:"rawPath"`
	ResPath         string    `gorm:"ResPath" json:"resPath" form:"resPath"`
	ResProbability  float64   `gorm:"ResProbability" json:"resProbability" form:"resProbability"`
}

func GetTestHistory(belongToChildID uint) ([]TestHistory, Serializer.Response) {
	var tHistory []TestHistory
	err := DB.Table("TestHistory").Where("belong_to_child_id = ?", belongToChildID).Find(&tHistory).Error
	if err != nil {
		logging.Info(err)
		code := e.ERROR
		return tHistory, Serializer.Response{
			Code:  code,
			Msg:   e.GetMsg(code),
			Error: "查询失败",
		}
	} else {
		code := e.SUCCESS
		return tHistory, Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}
}

// AddTestHistory function need to fix
func (t TestHistory) AddTestHistory() Serializer.Response {

	err := DB.Table("test_history").Create(&t).Error
	if err != nil {
		code := e.ERROR
		return Serializer.Response{
			Code:  code, // code fix
			Msg:   e.GetMsg(code),
			Error: "创建失败",
		}
	} else {
		code := e.SUCCESS
		return Serializer.Response{
			Code: code,
			Msg:  e.GetMsg(code),
		}
	}

}

type TestHisYear struct {
	Year          int         `json:"year"`
	Month         []int       `json:"month"`
	Days          [][]string  `json:"day"`
	Probabilities [][]float64 `json:"probability"`
}

func GetTestHisYear(belongToChildID uint) ([]TestHisYear, Serializer.Response) {
	var ths []TestHistory
	var thys []TestHisYear
	m := make(map[int]*TestHisYear)
	var tmp = struct {
		year  int
		month int
		day   int
	}{}
	DB.Model(&TestHistory{}).Where("belong_to_child_id = ?", belongToChildID).Find(&ths)
	if len(ths) == 0 {
		return nil, Serializer.Response{
			Code:  500,
			Msg:   e.GetMsg(500),
			Error: "空数据",
		}
	}
	var dayList []string
	var probaList []float64
	var year int
	for _, th := range ths {
		t := th.TestTime
		year = t.Year()
		month := int(t.Month())
		day := strconv.Itoa(t.Day())
		p := th.ResProbability

		if year != tmp.year {
			m[year] = new(TestHisYear)
			m[year].Year = year
			tmp.year = year
		}
		if month != tmp.month {
			tmp.month = month
			m[year].Month = append(m[year].Month, month)
			if len(dayList) != 0 {
				m[year].Days = append(m[year].Days, dayList)
			}
			if len(probaList) != 0 {
				m[year].Probabilities = append(m[year].Probabilities, probaList)
			}
			dayList = make([]string, 0)
			probaList = make([]float64, 0)
		}
		dayList = append(dayList, fmt.Sprintf("%d-%s", month, day))
		probaList = append(probaList, p)
	}
	m[year].Days = append(m[year].Days, dayList)
	m[year].Probabilities = append(m[year].Probabilities, probaList)
	for y, v := range m {
		var thy TestHisYear
		thy.Year = y
		thy.Month = v.Month
		thy.Days = v.Days
		thy.Probabilities = v.Probabilities
		thys = append(thys, thy)
	}
	if len(thys) == 0 {
		return nil, Serializer.Response{
			Code:  500,
			Msg:   e.GetMsg(500),
			Error: "空数据",
		}
	}
	return thys, Serializer.Response{
		Code: 200,
		Msg:  e.GetMsg(200),
		Data: "",
	}
}
