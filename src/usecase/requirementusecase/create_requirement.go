package requirementusecase

import (
	"context"
	"strings"

	"github.com/yuuki-tsujimura/architecture-study/src/domain/requirement"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/tag"
	"github.com/yuuki-tsujimura/architecture-study/src/domain/user"
	"github.com/yuuki-tsujimura/architecture-study/src/support/apperr"
)

type CreateRequirementUsecase struct {
	reqRepo  requirement.ReqRepository
	tagRepo  tag.TagRepository
	userRepo user.UserRepository
}

func NewCreateRequirementUsecase(reqRepo requirement.ReqRepository, tagRepo tag.TagRepository, userRepo user.UserRepository) *CreateRequirementUsecase {
	return &CreateRequirementUsecase{
		reqRepo,
		tagRepo,
		userRepo,
	}
}

func (u *CreateRequirementUsecase) Exec(ctx context.Context, input *CreateRequirementInput) error {
	tagExistsService := tag.NewTagIDExistsService(u.tagRepo)
	tagIDs := make([]tag.TagID, len(input.TagIDs))
	for _, tagID := range input.TagIDs {
		newTagID, err := tag.NewTagIDByVal(tagID)
		if err != nil {
			return err
		}
		tagIDs = append(tagIDs, newTagID)
	}
	isTagExist, existTagIDs, err := tagExistsService.Exec(ctx, tagIDs)
	if err != nil {
		return err
	}
	if !isTagExist {
		// existTagIDsをマップに保存し、検索を高速化
		found := make(map[tag.TagID]bool)
		for _, existID := range existTagIDs {
			found[existID] = true
		}

		// 存在しないタグを見つける
		var notExistTags []string
		for _, tagID := range tagIDs {
			if !found[tagID] {
				// TagIDをstringに変換するメソッドを使用して、ここを調整してください。
				notExistTags = append(notExistTags, tagID.String())
			}
		}

		// 存在しないタグのIDを結合し、エラーメッセージとして返す
		notExistTagIDsStr := strings.Join(notExistTags, ", ")
		return apperr.BadRequestWrapf(err, "存在していないtagIDが含まれています: %s", notExistTagIDsStr)
	}

	userExistsService := user.NewIsExistByIDService(u.userRepo)
	userID, err := user.NewUserIDByVal(input.UserID)
	if err != nil {
		return err
	}
	isUserExist, err := userExistsService.Run(ctx, userID)
	if err != nil {
		return err
	}
	if isUserExist {
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
		TagIDs:             tagIDs,
		UserID:             userID,
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
