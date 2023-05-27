package user

import (
	"time"
	"unicode/utf8"

	shared "github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type Career struct {
	careerID  CareerID
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

type CareerData struct {
	CareerID  string
	Detail    string
	StartYear int
	EndYear   *int
	CreatedAt time.Time
}

func newCareer(params *CareerParams, careerID CareerID, createdAt shared.CreatedAt) (*Career, error) {
	if err := checkDetailLength(params.Detail); err != nil {
		return nil, err
	}

	if err := validateStartYear(params.StartYear); err != nil {
		return nil, err
	}

	if err := validateEndYear(params.EndYear, params.StartYear); err != nil {
		return nil, err
	}

	return &Career{
		careerID:  careerID,
		detail:    params.Detail,
		startYear: params.StartYear,
		endYear:   params.EndYear,
		createdAt: createdAt,
	}, nil
}

func ReconstructCareersFromData(careersData []*CareerData) []*Career {
	careers := make([]*Career, len(careersData))

	for i, careerData := range careersData {
		var endYear uint16
		if careerData.EndYear != nil {
			endYear = uint16(*careerData.EndYear)
		}

		careers[i] = &Career{
			careerID:  CareerID(careerData.CareerID),
			detail:    careerData.Detail,
			startYear: uint16(careerData.StartYear),
			endYear:   endYear,
			createdAt: shared.CreatedAt(careerData.CreatedAt),
		}
	}

	return careers
}

func ConvertCareersToCareerData(careers []*Career) []*CareerData {
	careersData := make([]*CareerData, len(careers))

	for i, career := range careers {
		endYear := int(career.endYear)
		if endYear == 0 {
			endYear = -1
		}

		careersData[i] = &CareerData{
			CareerID:  career.careerID.String(),
			Detail:    career.detail,
			StartYear: int(career.startYear),
			EndYear:   &endYear,
			CreatedAt: time.Time(career.createdAt),
		}

		if endYear == -1 {
			careersData[i].EndYear = nil
		}
	}

	return careersData
}

func checkDetailLength(detail string) error {
	if utf8.RuneCountInString(detail) > 255 {
		return apperr.BadRequestf("detailは255文字以下である必要があります。(現在%d文字)", utf8.RuneCountInString(detail))
	}
	return nil
}

func validateStartYear(startYear uint16) error {
	if startYear < 1970 {
		return apperr.BadRequest("startYearは1970年以上である必要があります")
	}
	return nil
}

func validateEndYear(endYear uint16, startYear uint16) error {
	if endYear < 1970 || endYear <= startYear {
		return apperr.BadRequest("endYearは1970年以上であり、startYearより後の数値である必要があります")
	}
	return nil
}
