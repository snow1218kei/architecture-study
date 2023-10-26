package requirementusecase

import (
	"context"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/requirement"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
	"github.com/yuuki-tsujimura/architecture-study/src/usecase/requirementusecase/requirementinput"
)

type CreateRequirementUsecase struct {
	reqRepo  requirement.Repository
	tagRepo  tag.Repository
	userRepo user.UserRepository
}

func NewCreateRequirementUsecase(reqRepo requirement.Repository, tagRepo tag.Repository, userRepo user.UserRepository) *CreateRequirementUsecase {
	return &CreateRequirementUsecase{
		reqRepo,
		tagRepo,
		userRepo,
	}
}

func (u *CreateRequirementUsecase) Exec(ctx context.Context, input *requirementinput.CreateRequirementInput) error {
	tagExistsService := tag.NewTagIDExistsService(u.tagRepo)
	for tagID := range input.TagIDs {
		isExist, err := tagExistsService.Exec(ctx, tag.TagID(tagID))
		if err != nil {
			return err
		}
		if isExist {
			return apperr.BadRequestWrapf(err, "存在しているので他の名前でお願いします: %s", tagID)
		}
	}

	userExistsService := user.NewIsExistByIDService(u.userRepo)
	isExist, err := userExistsService.Run(ctx, user.UserID(input.UserID))
	if err != nil {
		return err
	}
	if isExist {
		return apperr.BadRequestWrapf(err, "存在しているので他の名前でお願いします: %s", input.UserID)
	}

	mentorRequirementParams := &requirement.MentorRequirementParams{
		Title:              input.Title,
		Category:           input.Category,
		ContractType:       input.ContractType,
		ConsultationMethod: input.ConsultationMethod,
		Description:        input.Description,
		Budget:             requirement.BudgetParams(input.Budget),
		ApplicationPeriod:  input.ApplicationPeriod,
		Status:             input.Status,
		TagIDs:             input.TagIDs,
		UserID:             input.UserID,
	}

	mentorRequirement, err := requirement.NewMentorRequirement(mentorRequirementParams)
	if err != nil {
		return err
	}

	if err := u.reqRepo.Store(ctx, mentorRequirement); err != nil {
		return err
	}

	return nil
}
