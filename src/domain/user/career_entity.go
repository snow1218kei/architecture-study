package user

import (
	"fmt"

	"github.com/yuuki-tsujimura/architecture-study/src/domain/id"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase"
)

type Career struct {
	CareerId  id.CareerId
	Detail    string
	StartYear int
	EndYear   int
}

func NewCareer(input usecase.Career) *Career {
	careerId := id.NewCareerId()
	detail := input.Detail
	startYear := input.StartYear
	endYear := input.EndYear

	career := &Career{
		CareerId:  careerId,
		Detail:    detail,
		StartYear: startYear,
		EndYear:   endYear,
	}

	career.CheckDetailLength()
	career.ValidateStartYear()
	career.ValidateEndYear()

	return career
}

func (career Career) CheckDetailLength() error {
	if len(career.Detail) > 255 {
		return fmt.Errorf("名前は255文字以下である必要があります。(現在%d文字)", len(career.Detail))
	}
	return nil
}

func (career Career) ValidateStartYear() error {
	if career.StartYear < 1970 {
		return fmt.Errorf("開始年は1970年以上である必要があります")
	}
	return nil
}

func (career Career) ValidateEndYear() error {
	if career.EndYear < 1970 || career.EndYear <= career.StartYear {
		return fmt.Errorf("終了年は1970年以上であり、開始年より後の数値である必要があります")
	}
	return nil
}
