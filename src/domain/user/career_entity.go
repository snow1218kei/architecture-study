package user

import (
	"time"
	"unicode/utf8"

	shared "github.com/yuuki-tsujimura/architecture-study/src/domain/shared/vo"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

const (
	maxDetailLength = 255
	minStartYear    = 1970
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

type CareerUpdateParams struct {
	ID        CareerID
	Detail    *string
	StartYear *uint16
	EndYear   *uint16
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
	if l := utf8.RuneCountInString(detail); l > maxDetailLength {
		return apperr.BadRequestf("detailは%d文字以下である必要があります。(現在%d文字)", maxDetailLength, l)
	}

	return nil
}

func validateStartYear(startYear uint16) error {
	if y := startYear; y < minStartYear {
		return apperr.BadRequestf("startYearは%d年以上である必要があります: %d", minStartYear, y)
	}

	return nil
}

func validateEndYear(endYear uint16, startYear uint16) error {
	if endYear < minStartYear || endYear <= startYear {
		return apperr.BadRequestf("endYearは%d年以上であり、startYearより後の数値である必要があります: %d", minStartYear, endYear)
	}
	return nil
}

func (c *Career) update(params *CareerUpdateParams) error {
	if params.Detail != nil {
		if err := checkDetailLength(*params.Detail); err != nil {
			return err
		}

		c.detail = *params.Detail
	}

	if params.StartYear != nil {
		if err := validateStartYear(*params.StartYear); err != nil {
			return err
		}

		c.startYear = *params.StartYear
	}

	if params.EndYear != nil && params.StartYear != nil {
		if err := validateEndYear(*params.EndYear, *params.StartYear); err != nil {
			return err
		}

		c.endYear = *params.EndYear
	} else if params.EndYear != nil {
		if err := validateEndYear(*params.EndYear, c.startYear); err != nil {
			return err
		}

		c.endYear = *params.EndYear
	}

	return nil
}
