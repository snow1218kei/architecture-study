package requirementusecase

import "context"

type GetAllRequirementsUsecase struct {
	queryService RequirementQueryService
}

func NewGetAllRequirementsUsecase(queryService RequirementQueryService) *GetAllRequirementsUsecase {
	return &GetAllRequirementsUsecase{
		queryService,
	}
}

func (u *GetAllRequirementsUsecase) Exec(ctx context.Context) ([]*GetMentorRequirementDTO, error) {
	mentorRequirents, err := u.queryService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return mentorRequirents, nil
}
