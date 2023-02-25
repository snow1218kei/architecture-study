package user

import (
	"fmt"

	"github.com/yuuki-tsujimura/architecture-study/src/domain/shared"
)

type Career struct {
	careerId  CareerID
	detail    string
	startYear uint16
	endYear   uint16
	createdAt shared.CreatedAt
}

type CareerParams struct {
	Detail    string
	StartYear uint16
	EndYear   uint16
}

func NewCareer(params CareerParams, careerID CareerID, createdAt shared.CreatedAt) (*Career, error) {
	if err := checkDetailLength(params.Detail); err != nil {
		return nil, err
	}

	if err := validateStartYear(params.StartYear); err != nil {
		return nil, err
	}

	if err := validateEndYear(params.EndYear, params.StartYear); err != nil {
		return nil, err
	}

	career := &Career{
		careerId:  careerID,
		detail:    params.Detail,
		startYear: params.StartYear,
		endYear:   params.EndYear,
		createdAt: createdAt,
	}
	return career, nil
}

func checkDetailLength(detail string) error {
	if len(detail) > 255 {
		return fmt.Errorf("名前は255文字以下である必要があります。(現在%d文字)", len(detail))
	}
	return nil
}

func validateStartYear(startYear uint16) error {
	if startYear < 1970 {
		return fmt.Errorf("開始年は1970年以上である必要があります")
	}
	return nil
}

func validateEndYear(endYear uint16, startYear uint16) error {
	if endYear < 1970 || endYear <= startYear {
		return fmt.Errorf("終了年は1970年以上であり、開始年より後の数値である必要があります")
	}
	return nil
}
